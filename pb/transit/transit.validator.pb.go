// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transit.proto

/*
Package transit is a generated protocol buffer package.

It is generated from these files:
	transit.proto

It has these top-level messages:
	Transit
*/
package transit

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Transit) Validate() error {
	for _, item := range this.Nodes {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Nodes", err)
		}
	}
	for _, item := range this.StopPairs {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("StopPairs", err)
		}
	}
	for _, item := range this.Routes {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Routes", err)
		}
	}
	for _, item := range this.Shapes {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Shapes", err)
		}
	}
	return nil
}
func (this *Transit_Node) Validate() error {
	return nil
}
func (this *Transit_StopPair) Validate() error {
	return nil
}
func (this *Transit_Route) Validate() error {
	return nil
}
func (this *Transit_Shape) Validate() error {
	return nil
}
