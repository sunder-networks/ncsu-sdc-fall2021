package main

import (
	"context"
	"flag"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	topoYamlPath := flag.String("topo_yaml", "", "filesystem path to the simplified yaml topology")
	topo := flag.String("topo", "", "one of: clos|linear|single|tree topology type must include count flag!!")
	count := flag.Int("count", 0, "count of nodes to deploy, when not specified through topo_yaml")

	// parse topology / retrieve
	flag.Parse()
	if *topoYamlPath != "" {
		// parse yaml into internal topology representation
	}
	if *topo != "" {
		// ensure count var not empty
		if *count == 0 {
			// this is an error
		}
	}
	// process veth creation

	// start switch processes

	// configure chassis config

	// set pipeline

	// devtest space down here
	fmt.Println("hello world")
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	go func() {
		time.Sleep(30 * time.Second)
		cancel()
	}()
	cmd := exec.CommandContext(ctx, "tail", "-f", "/dev/null")
	err := cmd.Run()
	fmt.Println(err.Error())

}

type stratumExec struct {
	deviceID          uint32
	chassisConfigFile string
	cpuPort           uint32
	logLevel          string
}

func (st *stratumExec) stratumStart() (string, []string) {
	binName := "stratum_bmv2"
	args := []string{
		fmt.Sprintf("-device_id=%d", st.deviceID),
		fmt.Sprintf("-chassis_config_file=%s", st.chassisConfigFile),
		fmt.Sprintf("-cpu_port=%d", st.cpuPort),
		fmt.Sprintf("-bmv2_log_level=%s", st.logLevel),
	}
	/*
					STRATUM_BMV2,
		            '-device_id=%d' % self.nodeId,
		            '-chassis_config_file=%s' % self.chassisConfigFile,
		            '-forwarding_pipeline_configs_file=%s/pipe.txt' % self.tmpDir,
		            '-persistent_config_dir=%s' % self.tmpDir,
		            '-initial_pipeline=%s' % self.json,
		            '-cpu_port=%s' % self.cpuPort,
		            '-external_stratum_urls=0.0.0.0:%d' % self.grpcPort,
		            '-local_stratum_url=localhost:%d' % pickUnusedPort(),
		            '-max_num_controllers_per_node=%d' % MAX_CONTROLLERS_PER_NODE,
		            '-write_req_log_file=%s/write-reqs.txt' % self.tmpDir,
		            '-bmv2_log_level=%s' % self.loglevel,
	*/
	return binName, args
}
