package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/proto/ywrapper"
	"github.com/vishvananda/netlink"
	"gopkg.in/yaml.v2"

	"log"
	"os"
	"os/signal"
	"path/filepath"

	"sunder.io/api/gen/stratumoc"
	stratumEnums "sunder.io/api/gen/stratumoc/enums"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/protobuf/proto"

	"github.com/eth0xFEED/p4runtime-go-client/pkg/client"
	p4v1 "github.com/p4lang/p4runtime/go/p4/v1"
)

type simpleTopology struct {
	Nodes []topoNode `yaml:"nodes"`
	Links []topoLink `yaml:"links"`
}
type topoNode struct {
	Name  string `yaml:"name"`
	DevID uint   `yaml:"id"`
}
type topoLink struct {
	SourceDeviceID   uint   `yaml:"src_dev_id,omitempty"`
	DestDeviceID     uint   `yaml:"dst_dev_id,omitempty"`
	SourceDeviceName string `yaml:"src_dev_name,omitempty"`
	DestDeviceName   string `yaml:"dst_dev_name,omitempty"`
}

func up(veth netlink.Link) error {
	err := netlink.LinkSetUp(veth)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	topoYamlPath := flag.String("topo_yaml", "", "filesystem path to the simplified yaml topology")
	grpcPortStart := flag.Uint("grpc_port", 50001, "base port to start from, incremented per node in topology default 50001, with 10 nodes range 50001-50011")

	// do we want to support non topology file driven spin ups ?
	// topo := flag.String("topo", "", "one of: clos|linear|single|tree topology type must include count flag!!")
	// count := flag.Int("count", 0, "count of nodes to deploy, when not specified through topo_yaml")

	flag.Parse()

	// parse topology / retrieve
	topology := simpleTopology{Nodes: []topoNode{}, Links: []topoLink{}}
	if *topoYamlPath != "" {
		log.Printf("parsing topology file: %s", *topoYamlPath)
		// parse yaml into internal topology representation
		sb, err := ioutil.ReadFile(*topoYamlPath)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(sb, &topology)
		if err != nil {
			panic(err)
		}
	}

	// find out how many interfaces per device, store interface name
	deviceInterfaces := make(map[string][]string)
	// map device_id to name
	deviceIDtoName := make(map[uint]string)
	// map name to device_id
	nameTodeviceID := make(map[string]uint)
	// list of vethPairs to create
	veths := make([]*netlink.Veth, 0)
	for _, dev := range topology.Nodes {
		if dev.DevID == 0 {
			// this is an error
		}
		if dev.Name == "" {
			dev.Name = fmt.Sprintf("device_%d", dev.DevID)
		}
		deviceIDtoName[dev.DevID] = dev.Name
		nameTodeviceID[dev.Name] = dev.DevID

		// make the first interface on every device a diag port
		// for packets in/out on a per device basis
		// leaf1_io <-> leaf1_pkts
		switchSideDiag := fmt.Sprintf("%s_io", dev.Name)
		hostSideDiag := fmt.Sprintf("%s_pkts", dev.Name)
		deviceInterfaces[dev.Name] = []string{
			switchSideDiag,
		}

		la := netlink.NewLinkAttrs()
		la.Name = switchSideDiag
		veths = append(veths, &netlink.Veth{LinkAttrs: la, PeerName: hostSideDiag})
	}

	// increment the link count to enable the case where multiple links exist between devices
	linkIdxCount := 0
	for _, link := range topology.Links {
		var srcDev, dstDev string

		if link.SourceDeviceName != "" && link.DestDeviceName != "" {
			srcDev = link.SourceDeviceName
			dstDev = link.DestDeviceName
		}
		if link.SourceDeviceID != 0 && link.DestDeviceID != 0 {
			srcDev = deviceIDtoName[link.SourceDeviceID]
			dstDev = deviceIDtoName[link.DestDeviceID]
		}

		// leaf1_spine1
		srcIntName := fmt.Sprintf("%s_%s_%d", srcDev, dstDev, linkIdxCount)
		deviceInterfaces[srcDev] = append(deviceInterfaces[srcDev], srcIntName)
		dstIntName := fmt.Sprintf("%s_%s_%d", dstDev, srcDev, linkIdxCount)
		deviceInterfaces[dstDev] = append(deviceInterfaces[dstDev], dstIntName)

		la := netlink.NewLinkAttrs()
		la.Name = srcIntName
		la.OperState = netlink.OperUp
		veths = append(veths, &netlink.Veth{LinkAttrs: la, PeerName: dstIntName})

		linkIdxCount++
	}

	// process veth creation
	log.Printf("creating %d veth interface pairs ", len(veths))
	for _, veth := range veths {
		err := netlink.LinkAdd(veth)
		if err != nil {
			panic(err)
		}
		lowerLayer, err := netlink.LinkByName(veth.PeerName)
		if err != nil {
			panic(err)
		}

		up(lowerLayer)
		up(veth)
	}

	// start switch processes
	errChan := make(chan error)
	// map device_name to dial Address
	deviceNametoAddress := make(map[string]string)

	ctx, cancel := context.WithCancel(context.Background())
	for _, dev := range topology.Nodes {
		st := &stratumExec{
			deviceID:   uint32(dev.DevID),
			deviceName: dev.Name,
			cpuPort:    255,
			logLevel:   "debug",
			// logLevel: "trace",
			grpcPort: *grpcPortStart,
		}
		log.Printf("starting switch %s serving on 127.0.0.1:%d ", st.deviceName, st.grpcPort)
		bin, args := st.stratumCmd()
		go func(st *stratumExec, bin string, args []string) {
			cmd := exec.CommandContext(ctx, bin, args...)
			logFile, err := os.Create(filepath.Join(st.baseFilePath, "bmv2.log"))
			if err != nil {
				panic(err)
			}
			defer logFile.Close()
			cmd.Stdout = logFile
			cmd.Stderr = logFile
			errChan <- cmd.Run()
		}(st, bin, args)
		deviceNametoAddress[dev.Name] = fmt.Sprintf("127.0.0.1:%d", *grpcPortStart)
		*grpcPortStart++
	}

	// configure chassis config
	for name, addr := range deviceNametoAddress {
		// /components/component/port
		makePort := func(dev *stratumoc.Device, name string, deviceID, port uint64) {
			dev.Component = append(dev.Component, &stratumoc.Device_ComponentKey{
				Name: name,
				Component: &stratumoc.Component{
					IntegratedCircuit: &stratumoc.Component_IntegratedCircuit{
						NodeId: &ywrapper.UintValue{Value: deviceID},
					},
					Linecard: &stratumoc.Component_Linecard{
						SlotId: &ywrapper.StringValue{Value: "1"},
					},
					Port: &stratumoc.Component_Port{
						PortId: &ywrapper.UintValue{Value: port},
					},
				},
			})
		}
		// /interfaces/interface
		makeInterface := func(dev *stratumoc.Device, name string, sdnPort uint64, enabled bool) {
			dev.Interface = append(dev.Interface, &stratumoc.Device_InterfaceKey{
				Name: name,
				Interface: &stratumoc.Interface{
					Id:      &ywrapper.UintValue{Value: sdnPort},
					Enabled: &ywrapper.BoolValue{Value: enabled},
					Ethernet: &stratumoc.Interface_Ethernet{
						AutoNegotiate: &ywrapper.BoolValue{Value: true},
						PortSpeed:     stratumEnums.OpenconfigIfEthernetETHERNETSPEED_OPENCONFIGIFETHERNETETHERNETSPEED_SPEED_100GB,
					},
				},
			})
		}

		device := new(stratumoc.Device)

		chassis := new(stratumoc.Device_ComponentKey)
		chassis.Name = name
		chassis.Component = &stratumoc.Component{Chassis: &stratumoc.Component_Chassis{Platform: stratumEnums.OpenconfigHerculesPlatformPLATFORMTYPE_OPENCONFIGHERCULESPLATFORMPLATFORMTYPE_P4_SOFT_SWITCH}}
		device.Component = append(device.Component, chassis)

		lineCard := new(stratumoc.Device_ComponentKey)
		lineCard.Name = ":lc-1"
		lineCard.Component = &stratumoc.Component{
			Id: &ywrapper.StringValue{Value: strconv.Itoa(int(nameTodeviceID[name]))},
			Linecard: &stratumoc.Component_Linecard{
				SlotId: &ywrapper.StringValue{Value: "1"},
			}}
		device.Component = append(device.Component, lineCard)

		for idx, iface := range deviceInterfaces[name] {
			makePort(device, iface, uint64(nameTodeviceID[name]), uint64(idx+1))
			makeInterface(device, iface, uint64(idx+1), true)
		}

		newDeviceConfig, err := proto.Marshal(device)
		if err != nil {
			log.Printf("error while marshal openconfig device %+v", err.Error())
		}

		cc, err := grpc.DialContext(ctx, addr, stratumGRPCopts...)
		if err != nil {
			panic(err)
		}
		defer cc.Close()
		gnmiClient := gnmi.NewGNMIClient(cc)
		resp, err := gnmiClient.Set(ctx, &gnmi.SetRequest{
			Replace: []*gnmi.Update{
				{
					Path: &gnmi.Path{},
					Val: &gnmi.TypedValue{
						// Value: &gnmi.TypedValue_AnyVal{
						// 	AnyVal: &anypb.Any{
						// 		TypeUrl: "type.googleapis.com/openconfig.Device",
						// 		Value:   newDeviceConfig,
						// 	},
						// },
						Value: &gnmi.TypedValue_BytesVal{
							BytesVal: newDeviceConfig,
						},
					},
				},
			},
		})
		if err != nil {
			log.Printf("error while setting device config %+v", err.Error())
			panic(err)
		}
		fmt.Printf("%T | %+v\n", resp, resp)
	}

	// set pipeline

	for name, addr := range deviceNametoAddress {
		log.Printf("setting %s pipeline on %s", name, addr)
		cc, err := grpc.DialContext(ctx, addr, stratumGRPCopts...)
		if err != nil {
			panic(err)
		}
		// defer cc.Close()

		highSideElection := time.Now().UnixNano()

		p4Client := client.NewClient(
			p4v1.NewP4RuntimeClient(cc),
			uint64(nameTodeviceID[name]),
			p4v1.Uint128{High: uint64(highSideElection), Low: 0},
		)
		stopCh := make(chan struct{})
		arpCh := make(chan bool)
		messageCh := make(chan *p4v1.StreamMessageResponse, 100)
		go p4Client.Run(stopCh, arpCh, messageCh)

		_, err = p4Client.SetFwdPipe("/p4/out/bmv2.json", "/p4/out/p4info.pb.txt", 0)
		if err != nil {
			panic(err)
		}
		log.Printf("success %s pipeline set ", name)
		// set the  pipeline and shutdown grpc connection
		// signal to stop stream channel since we wont continuously write through this p4runtime session
		stopCh <- struct{}{}
		cc.Close()

	}

	// wait for signal or other event before shutdown
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case err := <-errChan:
				log.Println(err)
			case sig := <-sigs:
				log.Println(sig)
				cancel()
				done <- true
			}
		}
	}()
	<-done
	log.Println("shutting down")

}

type stratumExec struct {
	deviceID uint32
	// chassisConfigFile string
	cpuPort      uint32
	logLevel     string
	deviceName   string
	grpcPort     uint
	baseFilePath string
}

func (st *stratumExec) stratumCmd() (string, []string) {
	binName := "/usr/bin/stratum_bmv2"

	st.baseFilePath = filepath.Join("/tmp", st.deviceName)
	os.MkdirAll(st.baseFilePath, os.ModePerm)

	args := []string{
		fmt.Sprintf("-device_id=%d", st.deviceID),
		// fmt.Sprintf("-chassis_config_file=%s", st.chassisConfigFile),
		fmt.Sprintf("-cpu_port=%d", st.cpuPort),
		// "-cpu_port=255",
		fmt.Sprintf("-bmv2_log_level=%s", st.logLevel),
		fmt.Sprintf("-persistent_config_dir=%s", st.baseFilePath),
		fmt.Sprintf("-chassis_config_file=%s", filepath.Join(st.baseFilePath, "chassis_config.pb.txt")),
		fmt.Sprintf("-external_stratum_urls=0.0.0.0:%d", st.grpcPort),
		fmt.Sprintf("-local_stratum_url=localhost:%d", st.grpcPort-100),
		fmt.Sprintf("-write_req_log_file=%s", filepath.Join(st.baseFilePath, "p4_writes.pb.txt")),
		fmt.Sprintf("-read_req_log_file=%s", filepath.Join(st.baseFilePath, "p4_reads.pb.txt")),
	}
	return binName, args
}

var stratumGRPCopts = []grpc.DialOption{
	grpc.WithBlock(), // until returns from dialer
	grpc.WithInsecure(),
	grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)),
	grpc.WithConnectParams(
		grpc.ConnectParams{
			// MinConnectTimeout: 5 * time.Second,
			Backoff: backoff.DefaultConfig,
		}),
}
