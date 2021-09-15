package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/stratum/testvectors/proto/testvector"
	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	tv := &testvector.TestVector{
		TestCases: []*testvector.TestCase{
			{
				TestCaseId: "demo-testcase",
				ActionGroups: []*testvector.ActionGroup{
					{
						ActionGroup: &testvector.ActionGroup_SequentialActionGroup{
							SequentialActionGroup: &testvector.SequentialActionGroup{
								Actions: []*testvector.Action{
									{Actions: &testvector.Action_ConfigOperation{
										ConfigOperation: &testvector.ConfigOperation{
											GnmiSetRequest: &gnmi.SetRequest{
												Update: []*gnmi.Update{
													{
														Path: &gnmi.Path{
															Elem: []*gnmi.PathElem{
																{Name: "interfaces"},
																{Name: "interface", Key: map[string]string{"name": "32/0"}},
																{Name: "config"},
																{Name: "health-indicator"},
															},
														},
														Val: &gnmi.TypedValue{
															Value: &gnmi.TypedValue_StringVal{
																StringVal: "GOOD",
															},
														},
													},
												},
											},
											GnmiSetResponse: &gnmi.SetResponse{
												Response: []*gnmi.UpdateResult{
													{
														Path: &gnmi.Path{
															Elem: []*gnmi.PathElem{
																{Name: "interfaces"},
																{Name: "interface", Key: map[string]string{"name": "32/0"}},
																{Name: "config"},
																{Name: "health-indicator"},
															},
														},
														Op: gnmi.UpdateResult_REPLACE,
													},
												},
											},
										},
									}},
								},
							},
						},
					},
				},
				Expectations: []*testvector.Expectation{},
			},
		},
	}
	// supported google/protobuf
	fmt.Println(prototext.Format(tv))
	// NOT supported golang/protobuf
	fmt.Println(proto.MarshalTextString(tv))
}
