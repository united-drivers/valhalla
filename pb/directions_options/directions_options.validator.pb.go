// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: directions_options.proto

/*
Package directions_options is a generated protocol buffer package.

It is generated from these files:
	directions_options.proto

It has these top-level messages:
	DirectionsOptions
*/
package directions_options

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/united-drivers/valhalla/pb/tripcommon"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DirectionsOptions) Validate() error {
	for _, item := range this.Locations {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Locations", err)
		}
	}
	for _, item := range this.AvoidLocations {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("AvoidLocations", err)
		}
	}
	for _, item := range this.Sources {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Sources", err)
		}
	}
	for _, item := range this.Targets {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Targets", err)
		}
	}
	for _, item := range this.Shape {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Shape", err)
		}
	}
	return nil
}
