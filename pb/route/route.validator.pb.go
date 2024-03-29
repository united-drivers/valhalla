// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: route.proto

/*
Package route is a generated protocol buffer package.

It is generated from these files:
	route.proto

It has these top-level messages:
	Route
*/
package route

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Route) Validate() error {
	if this.Trip != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Trip))); err != nil {
			return go_proto_validators.FieldError("Trip", err)
		}
	}
	return nil
}
func (this *Route_Location) Validate() error {
	return nil
}
func (this *Route_Summary) Validate() error {
	return nil
}
func (this *Route_TransitStop) Validate() error {
	return nil
}
func (this *Route_TransitInfo) Validate() error {
	for _, item := range this.TransitStops {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("TransitStops", err)
		}
	}
	return nil
}
func (this *Route_Maneuver) Validate() error {
	if this.Sign != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Sign))); err != nil {
			return go_proto_validators.FieldError("Sign", err)
		}
	}
	if this.TransitInfo != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.TransitInfo))); err != nil {
			return go_proto_validators.FieldError("TransitInfo", err)
		}
	}
	return nil
}
func (this *Route_Maneuver_Sign) Validate() error {
	for _, item := range this.ExitNumberElements {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("ExitNumberElements", err)
		}
	}
	for _, item := range this.ExitBranchElements {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("ExitBranchElements", err)
		}
	}
	for _, item := range this.ExitTowardElements {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("ExitTowardElements", err)
		}
	}
	for _, item := range this.ExitNameElements {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("ExitNameElements", err)
		}
	}
	return nil
}
func (this *Route_Maneuver_Sign_Element) Validate() error {
	return nil
}
func (this *Route_Leg) Validate() error {
	if this.Summary != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Summary))); err != nil {
			return go_proto_validators.FieldError("Summary", err)
		}
	}
	for _, item := range this.Maneuvers {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Maneuvers", err)
		}
	}
	return nil
}
func (this *Route_Trip) Validate() error {
	for _, item := range this.Locations {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Locations", err)
		}
	}
	if this.Summary != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Summary))); err != nil {
			return go_proto_validators.FieldError("Summary", err)
		}
	}
	for _, item := range this.Legs {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Legs", err)
		}
	}
	return nil
}
