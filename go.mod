module sunder.io/ncsu-sdc-fall2021

go 1.16

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20210826220005-b48c857c3a0e
	github.com/eth0xFEED/p4runtime-go-client v0.0.0-20210720013434-aa21cd92ed2c // indirect
	github.com/golang/protobuf v1.5.2
	github.com/openconfig/gnmi v0.0.0-20210903142221-87b435c38f6a
	github.com/openconfig/ygot v0.12.0
	github.com/p4lang/p4runtime v1.3.0 // indirect
	github.com/stratum/testvectors v0.0.0-20210708195704-f3d4d0750ada
	github.com/vishvananda/netlink v1.1.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
	sunder.io/api v0.0.0-00010101000000-000000000000
)

replace sunder.io/api => ./sunder
