// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmformat.proto

/*
Package OSMPBF is a generated protocol buffer package.

It is generated from these files:
	osmformat.proto

It has these top-level messages:
	HeaderBlock
	HeaderBBox
	PrimitiveBlock
	PrimitiveGroup
	StringTable
	Info
	DenseInfo
	ChangeSet
	Node
	DenseNodes
	Way
	Relation
*/
package OSMPBF

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *HeaderBlock) Validate() error {
	if this.Bbox != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Bbox))); err != nil {
			return go_proto_validators.FieldError("Bbox", err)
		}
	}
	return nil
}
func (this *HeaderBBox) Validate() error {
	return nil
}
func (this *PrimitiveBlock) Validate() error {
	if this.Stringtable != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Stringtable))); err != nil {
			return go_proto_validators.FieldError("Stringtable", err)
		}
	}
	for _, item := range this.Primitivegroup {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Primitivegroup", err)
		}
	}
	return nil
}
func (this *PrimitiveGroup) Validate() error {
	for _, item := range this.Nodes {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Nodes", err)
		}
	}
	if this.Dense != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Dense))); err != nil {
			return go_proto_validators.FieldError("Dense", err)
		}
	}
	for _, item := range this.Ways {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Ways", err)
		}
	}
	for _, item := range this.Relations {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Relations", err)
		}
	}
	for _, item := range this.Changesets {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Changesets", err)
		}
	}
	return nil
}
func (this *StringTable) Validate() error {
	return nil
}
func (this *Info) Validate() error {
	return nil
}
func (this *DenseInfo) Validate() error {
	return nil
}
func (this *ChangeSet) Validate() error {
	return nil
}
func (this *Node) Validate() error {
	if this.Info != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Info))); err != nil {
			return go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *DenseNodes) Validate() error {
	if this.Denseinfo != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Denseinfo))); err != nil {
			return go_proto_validators.FieldError("Denseinfo", err)
		}
	}
	return nil
}
func (this *Way) Validate() error {
	if this.Info != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Info))); err != nil {
			return go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *Relation) Validate() error {
	if this.Info != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Info))); err != nil {
			return go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
