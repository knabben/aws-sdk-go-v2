// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
)

func ExampleConfigTypeData_outputUsage() {
	var union types.ConfigTypeData
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfigTypeDataMemberAntennaDownlinkConfig:
		_ = v.Value // Value is AntennaDownlinkConfig

	case *types.ConfigTypeDataMemberAntennaDownlinkDemodDecodeConfig:
		_ = v.Value // Value is AntennaDownlinkDemodDecodeConfig

	case *types.ConfigTypeDataMemberAntennaUplinkConfig:
		_ = v.Value // Value is AntennaUplinkConfig

	case *types.ConfigTypeDataMemberDataflowEndpointConfig:
		_ = v.Value // Value is DataflowEndpointConfig

	case *types.ConfigTypeDataMemberTrackingConfig:
		_ = v.Value // Value is TrackingConfig

	case *types.ConfigTypeDataMemberUplinkEchoConfig:
		_ = v.Value // Value is UplinkEchoConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}