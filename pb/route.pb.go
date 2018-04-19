// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: route.proto

/*
Package valhalla is a generated protocol buffer package.

It is generated from these files:
	route.proto

It has these top-level messages:
	Route
*/
package valhalla

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Route struct {
	Trip             *Route_Trip `protobuf:"bytes,1,opt,name=trip" json:"trip,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0} }

func (m *Route) GetTrip() *Route_Trip {
	if m != nil {
		return m.Trip
	}
	return nil
}

type Route_Location struct {
	Lat              *float32 `protobuf:"fixed32,1,opt,name=lat" json:"lat,omitempty"`
	Lon              *float32 `protobuf:"fixed32,2,opt,name=lon" json:"lon,omitempty"`
	Type             *string  `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Heading          *uint32  `protobuf:"varint,4,opt,name=heading" json:"heading,omitempty"`
	Name             *string  `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Street           *string  `protobuf:"bytes,6,opt,name=street" json:"street,omitempty"`
	City             *string  `protobuf:"bytes,7,opt,name=city" json:"city,omitempty"`
	State            *string  `protobuf:"bytes,8,opt,name=state" json:"state,omitempty"`
	PostalCode       *string  `protobuf:"bytes,9,opt,name=postal_code,json=postalCode" json:"postal_code,omitempty"`
	Country          *string  `protobuf:"bytes,10,opt,name=country" json:"country,omitempty"`
	DateTime         *string  `protobuf:"bytes,11,opt,name=date_time,json=dateTime" json:"date_time,omitempty"`
	SideOfStreet     *string  `protobuf:"bytes,12,opt,name=side_of_street,json=sideOfStreet" json:"side_of_street,omitempty"`
	OriginalIndex    *uint32  `protobuf:"varint,13,opt,name=original_index,json=originalIndex" json:"original_index,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Route_Location) Reset()                    { *m = Route_Location{} }
func (m *Route_Location) String() string            { return proto.CompactTextString(m) }
func (*Route_Location) ProtoMessage()               {}
func (*Route_Location) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0, 0} }

func (m *Route_Location) GetLat() float32 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *Route_Location) GetLon() float32 {
	if m != nil && m.Lon != nil {
		return *m.Lon
	}
	return 0
}

func (m *Route_Location) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Route_Location) GetHeading() uint32 {
	if m != nil && m.Heading != nil {
		return *m.Heading
	}
	return 0
}

func (m *Route_Location) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Route_Location) GetStreet() string {
	if m != nil && m.Street != nil {
		return *m.Street
	}
	return ""
}

func (m *Route_Location) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *Route_Location) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *Route_Location) GetPostalCode() string {
	if m != nil && m.PostalCode != nil {
		return *m.PostalCode
	}
	return ""
}

func (m *Route_Location) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *Route_Location) GetDateTime() string {
	if m != nil && m.DateTime != nil {
		return *m.DateTime
	}
	return ""
}

func (m *Route_Location) GetSideOfStreet() string {
	if m != nil && m.SideOfStreet != nil {
		return *m.SideOfStreet
	}
	return ""
}

func (m *Route_Location) GetOriginalIndex() uint32 {
	if m != nil && m.OriginalIndex != nil {
		return *m.OriginalIndex
	}
	return 0
}

type Route_Summary struct {
	Length           *float32 `protobuf:"fixed32,1,opt,name=length" json:"length,omitempty"`
	Time             *uint32  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	MinLat           *float32 `protobuf:"fixed32,3,opt,name=min_lat,json=minLat" json:"min_lat,omitempty"`
	MinLon           *float32 `protobuf:"fixed32,4,opt,name=min_lon,json=minLon" json:"min_lon,omitempty"`
	MaxLat           *float32 `protobuf:"fixed32,5,opt,name=max_lat,json=maxLat" json:"max_lat,omitempty"`
	MaxLon           *float32 `protobuf:"fixed32,6,opt,name=max_lon,json=maxLon" json:"max_lon,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Route_Summary) Reset()                    { *m = Route_Summary{} }
func (m *Route_Summary) String() string            { return proto.CompactTextString(m) }
func (*Route_Summary) ProtoMessage()               {}
func (*Route_Summary) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0, 1} }

func (m *Route_Summary) GetLength() float32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *Route_Summary) GetTime() uint32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Route_Summary) GetMinLat() float32 {
	if m != nil && m.MinLat != nil {
		return *m.MinLat
	}
	return 0
}

func (m *Route_Summary) GetMinLon() float32 {
	if m != nil && m.MinLon != nil {
		return *m.MinLon
	}
	return 0
}

func (m *Route_Summary) GetMaxLat() float32 {
	if m != nil && m.MaxLat != nil {
		return *m.MaxLat
	}
	return 0
}

func (m *Route_Summary) GetMaxLon() float32 {
	if m != nil && m.MaxLon != nil {
		return *m.MaxLon
	}
	return 0
}

type Route_TransitStop struct {
	Type              *string  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	OnestopId         *string  `protobuf:"bytes,2,opt,name=onestop_id,json=onestopId" json:"onestop_id,omitempty"`
	Name              *string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	ArrivalDateTime   *string  `protobuf:"bytes,4,opt,name=arrival_date_time,json=arrivalDateTime" json:"arrival_date_time,omitempty"`
	DepartureDateTime *string  `protobuf:"bytes,5,opt,name=departure_date_time,json=departureDateTime" json:"departure_date_time,omitempty"`
	IsParentStop      *bool    `protobuf:"varint,6,opt,name=is_parent_stop,json=isParentStop" json:"is_parent_stop,omitempty"`
	AssumedSchedule   *bool    `protobuf:"varint,7,opt,name=assumed_schedule,json=assumedSchedule" json:"assumed_schedule,omitempty"`
	Lat               *float32 `protobuf:"fixed32,8,opt,name=lat" json:"lat,omitempty"`
	Lon               *float32 `protobuf:"fixed32,9,opt,name=lon" json:"lon,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *Route_TransitStop) Reset()                    { *m = Route_TransitStop{} }
func (m *Route_TransitStop) String() string            { return proto.CompactTextString(m) }
func (*Route_TransitStop) ProtoMessage()               {}
func (*Route_TransitStop) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0, 2} }

func (m *Route_TransitStop) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Route_TransitStop) GetOnestopId() string {
	if m != nil && m.OnestopId != nil {
		return *m.OnestopId
	}
	return ""
}

func (m *Route_TransitStop) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Route_TransitStop) GetArrivalDateTime() string {
	if m != nil && m.ArrivalDateTime != nil {
		return *m.ArrivalDateTime
	}
	return ""
}

func (m *Route_TransitStop) GetDepartureDateTime() string {
	if m != nil && m.DepartureDateTime != nil {
		return *m.DepartureDateTime
	}
	return ""
}

func (m *Route_TransitStop) GetIsParentStop() bool {
	if m != nil && m.IsParentStop != nil {
		return *m.IsParentStop
	}
	return false
}

func (m *Route_TransitStop) GetAssumedSchedule() bool {
	if m != nil && m.AssumedSchedule != nil {
		return *m.AssumedSchedule
	}
	return false
}

func (m *Route_TransitStop) GetLat() float32 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *Route_TransitStop) GetLon() float32 {
	if m != nil && m.Lon != nil {
		return *m.Lon
	}
	return 0
}

type Route_TransitInfo struct {
	OnestopId         *string              `protobuf:"bytes,1,opt,name=onestop_id,json=onestopId" json:"onestop_id,omitempty"`
	ShortName         *string              `protobuf:"bytes,2,opt,name=short_name,json=shortName" json:"short_name,omitempty"`
	LongName          *string              `protobuf:"bytes,3,opt,name=long_name,json=longName" json:"long_name,omitempty"`
	Headsign          *string              `protobuf:"bytes,4,opt,name=headsign" json:"headsign,omitempty"`
	Color             *uint32              `protobuf:"varint,5,opt,name=color" json:"color,omitempty"`
	TextColor         *uint32              `protobuf:"varint,6,opt,name=text_color,json=textColor" json:"text_color,omitempty"`
	Description       *string              `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	OperatorOnestopId *string              `protobuf:"bytes,8,opt,name=operator_onestop_id,json=operatorOnestopId" json:"operator_onestop_id,omitempty"`
	OperatorName      *string              `protobuf:"bytes,9,opt,name=operator_name,json=operatorName" json:"operator_name,omitempty"`
	OperatorUrl       *string              `protobuf:"bytes,10,opt,name=operator_url,json=operatorUrl" json:"operator_url,omitempty"`
	TransitStops      []*Route_TransitStop `protobuf:"bytes,11,rep,name=transit_stops,json=transitStops" json:"transit_stops,omitempty"`
	XXX_unrecognized  []byte               `json:"-"`
}

func (m *Route_TransitInfo) Reset()                    { *m = Route_TransitInfo{} }
func (m *Route_TransitInfo) String() string            { return proto.CompactTextString(m) }
func (*Route_TransitInfo) ProtoMessage()               {}
func (*Route_TransitInfo) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0, 3} }

func (m *Route_TransitInfo) GetOnestopId() string {
	if m != nil && m.OnestopId != nil {
		return *m.OnestopId
	}
	return ""
}

func (m *Route_TransitInfo) GetShortName() string {
	if m != nil && m.ShortName != nil {
		return *m.ShortName
	}
	return ""
}

func (m *Route_TransitInfo) GetLongName() string {
	if m != nil && m.LongName != nil {
		return *m.LongName
	}
	return ""
}

func (m *Route_TransitInfo) GetHeadsign() string {
	if m != nil && m.Headsign != nil {
		return *m.Headsign
	}
	return ""
}

func (m *Route_TransitInfo) GetColor() uint32 {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return 0
}

func (m *Route_TransitInfo) GetTextColor() uint32 {
	if m != nil && m.TextColor != nil {
		return *m.TextColor
	}
	return 0
}

func (m *Route_TransitInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Route_TransitInfo) GetOperatorOnestopId() string {
	if m != nil && m.OperatorOnestopId != nil {
		return *m.OperatorOnestopId
	}
	return ""
}

func (m *Route_TransitInfo) GetOperatorName() string {
	if m != nil && m.OperatorName != nil {
		return *m.OperatorName
	}
	return ""
}

func (m *Route_TransitInfo) GetOperatorUrl() string {
	if m != nil && m.OperatorUrl != nil {
		return *m.OperatorUrl
	}
	return ""
}

func (m *Route_TransitInfo) GetTransitStops() []*Route_TransitStop {
	if m != nil {
		return m.TransitStops
	}
	return nil
}

type Route_Maneuver struct {
	Type                             *uint32              `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Instruction                      *string              `protobuf:"bytes,2,opt,name=instruction" json:"instruction,omitempty"`
	StreetNames                      []string             `protobuf:"bytes,3,rep,name=street_names,json=streetNames" json:"street_names,omitempty"`
	Length                           *float32             `protobuf:"fixed32,4,opt,name=length" json:"length,omitempty"`
	Time                             *uint32              `protobuf:"varint,5,opt,name=time" json:"time,omitempty"`
	BeginCardinalDirection           *string              `protobuf:"bytes,6,opt,name=begin_cardinal_direction,json=beginCardinalDirection" json:"begin_cardinal_direction,omitempty"`
	BeginHeading                     *uint32              `protobuf:"varint,7,opt,name=begin_heading,json=beginHeading" json:"begin_heading,omitempty"`
	BeginShapeIndex                  *uint32              `protobuf:"varint,8,opt,name=begin_shape_index,json=beginShapeIndex" json:"begin_shape_index,omitempty"`
	EndShapeIndex                    *uint32              `protobuf:"varint,9,opt,name=end_shape_index,json=endShapeIndex" json:"end_shape_index,omitempty"`
	Toll                             *bool                `protobuf:"varint,10,opt,name=toll" json:"toll,omitempty"`
	Rough                            *bool                `protobuf:"varint,11,opt,name=rough" json:"rough,omitempty"`
	VerbalTransitionAlertInstruction *string              `protobuf:"bytes,12,opt,name=verbal_transition_alert_instruction,json=verbalTransitionAlertInstruction" json:"verbal_transition_alert_instruction,omitempty"`
	VerbalPreTransitionInstruction   *string              `protobuf:"bytes,13,opt,name=verbal_pre_transition_instruction,json=verbalPreTransitionInstruction" json:"verbal_pre_transition_instruction,omitempty"`
	VerbalPostTransitionInstruction  *string              `protobuf:"bytes,14,opt,name=verbal_post_transition_instruction,json=verbalPostTransitionInstruction" json:"verbal_post_transition_instruction,omitempty"`
	BeginStreetNames                 []string             `protobuf:"bytes,15,rep,name=begin_street_names,json=beginStreetNames" json:"begin_street_names,omitempty"`
	Sign                             *Route_Maneuver_Sign `protobuf:"bytes,16,opt,name=sign" json:"sign,omitempty"`
	RoundaboutExitCount              *uint32              `protobuf:"varint,17,opt,name=roundabout_exit_count,json=roundaboutExitCount" json:"roundabout_exit_count,omitempty"`
	DepartInstruction                *string              `protobuf:"bytes,18,opt,name=depart_instruction,json=departInstruction" json:"depart_instruction,omitempty"`
	VerbalDepartInstruction          *string              `protobuf:"bytes,19,opt,name=verbal_depart_instruction,json=verbalDepartInstruction" json:"verbal_depart_instruction,omitempty"`
	ArriveInstruction                *string              `protobuf:"bytes,20,opt,name=arrive_instruction,json=arriveInstruction" json:"arrive_instruction,omitempty"`
	VerbalArriveInstruction          *string              `protobuf:"bytes,21,opt,name=verbal_arrive_instruction,json=verbalArriveInstruction" json:"verbal_arrive_instruction,omitempty"`
	TransitInfo                      *Route_TransitInfo   `protobuf:"bytes,22,opt,name=transit_info,json=transitInfo" json:"transit_info,omitempty"`
	VerbalMultiCue                   *bool                `protobuf:"varint,23,opt,name=verbal_multi_cue,json=verbalMultiCue" json:"verbal_multi_cue,omitempty"`
	TravelMode                       *string              `protobuf:"bytes,24,opt,name=travel_mode,json=travelMode" json:"travel_mode,omitempty"`
	TravelType                       *string              `protobuf:"bytes,25,opt,name=travel_type,json=travelType" json:"travel_type,omitempty"`
	XXX_unrecognized                 []byte               `json:"-"`
}

func (m *Route_Maneuver) Reset()                    { *m = Route_Maneuver{} }
func (m *Route_Maneuver) String() string            { return proto.CompactTextString(m) }
func (*Route_Maneuver) ProtoMessage()               {}
func (*Route_Maneuver) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0, 4} }

func (m *Route_Maneuver) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Route_Maneuver) GetInstruction() string {
	if m != nil && m.Instruction != nil {
		return *m.Instruction
	}
	return ""
}

func (m *Route_Maneuver) GetStreetNames() []string {
	if m != nil {
		return m.StreetNames
	}
	return nil
}

func (m *Route_Maneuver) GetLength() float32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *Route_Maneuver) GetTime() uint32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Route_Maneuver) GetBeginCardinalDirection() string {
	if m != nil && m.BeginCardinalDirection != nil {
		return *m.BeginCardinalDirection
	}
	return ""
}

func (m *Route_Maneuver) GetBeginHeading() uint32 {
	if m != nil && m.BeginHeading != nil {
		return *m.BeginHeading
	}
	return 0
}

func (m *Route_Maneuver) GetBeginShapeIndex() uint32 {
	if m != nil && m.BeginShapeIndex != nil {
		return *m.BeginShapeIndex
	}
	return 0
}

func (m *Route_Maneuver) GetEndShapeIndex() uint32 {
	if m != nil && m.EndShapeIndex != nil {
		return *m.EndShapeIndex
	}
	return 0
}

func (m *Route_Maneuver) GetToll() bool {
	if m != nil && m.Toll != nil {
		return *m.Toll
	}
	return false
}

func (m *Route_Maneuver) GetRough() bool {
	if m != nil && m.Rough != nil {
		return *m.Rough
	}
	return false
}

func (m *Route_Maneuver) GetVerbalTransitionAlertInstruction() string {
	if m != nil && m.VerbalTransitionAlertInstruction != nil {
		return *m.VerbalTransitionAlertInstruction
	}
	return ""
}

func (m *Route_Maneuver) GetVerbalPreTransitionInstruction() string {
	if m != nil && m.VerbalPreTransitionInstruction != nil {
		return *m.VerbalPreTransitionInstruction
	}
	return ""
}

func (m *Route_Maneuver) GetVerbalPostTransitionInstruction() string {
	if m != nil && m.VerbalPostTransitionInstruction != nil {
		return *m.VerbalPostTransitionInstruction
	}
	return ""
}

func (m *Route_Maneuver) GetBeginStreetNames() []string {
	if m != nil {
		return m.BeginStreetNames
	}
	return nil
}

func (m *Route_Maneuver) GetSign() *Route_Maneuver_Sign {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *Route_Maneuver) GetRoundaboutExitCount() uint32 {
	if m != nil && m.RoundaboutExitCount != nil {
		return *m.RoundaboutExitCount
	}
	return 0
}

func (m *Route_Maneuver) GetDepartInstruction() string {
	if m != nil && m.DepartInstruction != nil {
		return *m.DepartInstruction
	}
	return ""
}

func (m *Route_Maneuver) GetVerbalDepartInstruction() string {
	if m != nil && m.VerbalDepartInstruction != nil {
		return *m.VerbalDepartInstruction
	}
	return ""
}

func (m *Route_Maneuver) GetArriveInstruction() string {
	if m != nil && m.ArriveInstruction != nil {
		return *m.ArriveInstruction
	}
	return ""
}

func (m *Route_Maneuver) GetVerbalArriveInstruction() string {
	if m != nil && m.VerbalArriveInstruction != nil {
		return *m.VerbalArriveInstruction
	}
	return ""
}

func (m *Route_Maneuver) GetTransitInfo() *Route_TransitInfo {
	if m != nil {
		return m.TransitInfo
	}
	return nil
}

func (m *Route_Maneuver) GetVerbalMultiCue() bool {
	if m != nil && m.VerbalMultiCue != nil {
		return *m.VerbalMultiCue
	}
	return false
}

func (m *Route_Maneuver) GetTravelMode() string {
	if m != nil && m.TravelMode != nil {
		return *m.TravelMode
	}
	return ""
}

func (m *Route_Maneuver) GetTravelType() string {
	if m != nil && m.TravelType != nil {
		return *m.TravelType
	}
	return ""
}

type Route_Maneuver_Sign struct {
	ExitNumberElements []*Route_Maneuver_Sign_Element `protobuf:"bytes,1,rep,name=exit_number_elements,json=exitNumberElements" json:"exit_number_elements,omitempty"`
	ExitBranchElements []*Route_Maneuver_Sign_Element `protobuf:"bytes,2,rep,name=exit_branch_elements,json=exitBranchElements" json:"exit_branch_elements,omitempty"`
	ExitTowardElements []*Route_Maneuver_Sign_Element `protobuf:"bytes,3,rep,name=exit_toward_elements,json=exitTowardElements" json:"exit_toward_elements,omitempty"`
	ExitNameElements   []*Route_Maneuver_Sign_Element `protobuf:"bytes,4,rep,name=exit_name_elements,json=exitNameElements" json:"exit_name_elements,omitempty"`
	XXX_unrecognized   []byte                         `json:"-"`
}

func (m *Route_Maneuver_Sign) Reset()                    { *m = Route_Maneuver_Sign{} }
func (m *Route_Maneuver_Sign) String() string            { return proto.CompactTextString(m) }
func (*Route_Maneuver_Sign) ProtoMessage()               {}
func (*Route_Maneuver_Sign) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0, 4, 0} }

func (m *Route_Maneuver_Sign) GetExitNumberElements() []*Route_Maneuver_Sign_Element {
	if m != nil {
		return m.ExitNumberElements
	}
	return nil
}

func (m *Route_Maneuver_Sign) GetExitBranchElements() []*Route_Maneuver_Sign_Element {
	if m != nil {
		return m.ExitBranchElements
	}
	return nil
}

func (m *Route_Maneuver_Sign) GetExitTowardElements() []*Route_Maneuver_Sign_Element {
	if m != nil {
		return m.ExitTowardElements
	}
	return nil
}

func (m *Route_Maneuver_Sign) GetExitNameElements() []*Route_Maneuver_Sign_Element {
	if m != nil {
		return m.ExitNameElements
	}
	return nil
}

type Route_Maneuver_Sign_Element struct {
	Text             *string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	ConsecutiveCount *uint32 `protobuf:"varint,2,opt,name=consecutive_count,json=consecutiveCount" json:"consecutive_count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Route_Maneuver_Sign_Element) Reset()         { *m = Route_Maneuver_Sign_Element{} }
func (m *Route_Maneuver_Sign_Element) String() string { return proto.CompactTextString(m) }
func (*Route_Maneuver_Sign_Element) ProtoMessage()    {}
func (*Route_Maneuver_Sign_Element) Descriptor() ([]byte, []int) {
	return fileDescriptorRoute, []int{0, 4, 0, 0}
}

func (m *Route_Maneuver_Sign_Element) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *Route_Maneuver_Sign_Element) GetConsecutiveCount() uint32 {
	if m != nil && m.ConsecutiveCount != nil {
		return *m.ConsecutiveCount
	}
	return 0
}

type Route_Leg struct {
	Summary          *Route_Summary    `protobuf:"bytes,1,opt,name=summary" json:"summary,omitempty"`
	Maneuvers        []*Route_Maneuver `protobuf:"bytes,2,rep,name=maneuvers" json:"maneuvers,omitempty"`
	Shape            *string           `protobuf:"bytes,3,opt,name=shape" json:"shape,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Route_Leg) Reset()                    { *m = Route_Leg{} }
func (m *Route_Leg) String() string            { return proto.CompactTextString(m) }
func (*Route_Leg) ProtoMessage()               {}
func (*Route_Leg) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0, 5} }

func (m *Route_Leg) GetSummary() *Route_Summary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *Route_Leg) GetManeuvers() []*Route_Maneuver {
	if m != nil {
		return m.Maneuvers
	}
	return nil
}

func (m *Route_Leg) GetShape() string {
	if m != nil && m.Shape != nil {
		return *m.Shape
	}
	return ""
}

type Route_Trip struct {
	Locations        []*Route_Location `protobuf:"bytes,1,rep,name=locations" json:"locations,omitempty"`
	Summary          *Route_Summary    `protobuf:"bytes,2,opt,name=summary" json:"summary,omitempty"`
	Legs             []*Route_Leg      `protobuf:"bytes,3,rep,name=legs" json:"legs,omitempty"`
	StatusMessage    *string           `protobuf:"bytes,4,opt,name=status_message,json=statusMessage" json:"status_message,omitempty"`
	Status           *uint32           `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	Units            *string           `protobuf:"bytes,6,opt,name=units" json:"units,omitempty"`
	Language         *string           `protobuf:"bytes,7,opt,name=language" json:"language,omitempty"`
	Id               *string           `protobuf:"bytes,8,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Route_Trip) Reset()                    { *m = Route_Trip{} }
func (m *Route_Trip) String() string            { return proto.CompactTextString(m) }
func (*Route_Trip) ProtoMessage()               {}
func (*Route_Trip) Descriptor() ([]byte, []int) { return fileDescriptorRoute, []int{0, 6} }

func (m *Route_Trip) GetLocations() []*Route_Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *Route_Trip) GetSummary() *Route_Summary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *Route_Trip) GetLegs() []*Route_Leg {
	if m != nil {
		return m.Legs
	}
	return nil
}

func (m *Route_Trip) GetStatusMessage() string {
	if m != nil && m.StatusMessage != nil {
		return *m.StatusMessage
	}
	return ""
}

func (m *Route_Trip) GetStatus() uint32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *Route_Trip) GetUnits() string {
	if m != nil && m.Units != nil {
		return *m.Units
	}
	return ""
}

func (m *Route_Trip) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *Route_Trip) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Route)(nil), "valhalla.Route")
	proto.RegisterType((*Route_Location)(nil), "valhalla.Route.Location")
	proto.RegisterType((*Route_Summary)(nil), "valhalla.Route.Summary")
	proto.RegisterType((*Route_TransitStop)(nil), "valhalla.Route.TransitStop")
	proto.RegisterType((*Route_TransitInfo)(nil), "valhalla.Route.TransitInfo")
	proto.RegisterType((*Route_Maneuver)(nil), "valhalla.Route.Maneuver")
	proto.RegisterType((*Route_Maneuver_Sign)(nil), "valhalla.Route.Maneuver.Sign")
	proto.RegisterType((*Route_Maneuver_Sign_Element)(nil), "valhalla.Route.Maneuver.Sign.Element")
	proto.RegisterType((*Route_Leg)(nil), "valhalla.Route.Leg")
	proto.RegisterType((*Route_Trip)(nil), "valhalla.Route.Trip")
}

func init() { proto.RegisterFile("route.proto", fileDescriptorRoute) }

var fileDescriptorRoute = []byte{
	// 1356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x6e, 0x1b, 0xb7,
	0x13, 0x87, 0x3e, 0x1c, 0x49, 0x23, 0xcb, 0x96, 0x69, 0x27, 0xde, 0x28, 0xc8, 0x3f, 0x4a, 0xf2,
	0x4f, 0xeb, 0xa6, 0x8d, 0x80, 0xe4, 0x50, 0x14, 0x3d, 0x14, 0x4d, 0xed, 0x00, 0x75, 0x1b, 0x27,
	0xc1, 0xca, 0x45, 0x8f, 0x04, 0xad, 0xa5, 0x25, 0x02, 0xbb, 0xe4, 0x82, 0xe4, 0xba, 0xf2, 0x0b,
	0xf4, 0x0d, 0x7a, 0x28, 0x7a, 0xe8, 0xa9, 0xf7, 0xa2, 0x2f, 0xd3, 0xd7, 0x29, 0x38, 0xe4, 0xae,
	0xd6, 0x96, 0x53, 0x34, 0xb7, 0x9d, 0xdf, 0x0c, 0x7f, 0xe4, 0x7c, 0xee, 0x40, 0x5f, 0xab, 0xc2,
	0xf2, 0x49, 0xae, 0x95, 0x55, 0xa4, 0x7b, 0xc1, 0xd2, 0x05, 0x4b, 0x53, 0xf6, 0xe8, 0xcf, 0xbb,
	0xb0, 0x11, 0x3b, 0x0d, 0x39, 0x80, 0xb6, 0xd5, 0x22, 0x8f, 0x1a, 0xe3, 0xc6, 0x41, 0xff, 0xc5,
	0xde, 0xa4, 0x34, 0x99, 0xa0, 0x7a, 0x72, 0xaa, 0x45, 0x1e, 0xa3, 0xc5, 0xe8, 0xef, 0x26, 0x74,
	0x5f, 0xab, 0x19, 0xb3, 0x42, 0x49, 0x32, 0x84, 0x56, 0xca, 0x2c, 0x9e, 0x6a, 0xc6, 0xee, 0x13,
	0x11, 0x25, 0xa3, 0x66, 0x40, 0x94, 0x24, 0x04, 0xda, 0xf6, 0x32, 0xe7, 0x51, 0x6b, 0xdc, 0x38,
	0xe8, 0xc5, 0xf8, 0x4d, 0x22, 0xe8, 0x2c, 0x38, 0x4b, 0x84, 0x9c, 0x47, 0xed, 0x71, 0xe3, 0x60,
	0x10, 0x97, 0xa2, 0xb3, 0x96, 0x2c, 0xe3, 0xd1, 0x86, 0xb7, 0x76, 0xdf, 0xe4, 0x0e, 0xdc, 0x32,
	0x56, 0x73, 0x6e, 0xa3, 0x5b, 0x88, 0x06, 0xc9, 0xd9, 0xce, 0x84, 0xbd, 0x8c, 0x3a, 0xde, 0xd6,
	0x7d, 0x93, 0x3d, 0xd8, 0x30, 0x96, 0x59, 0x1e, 0x75, 0x11, 0xf4, 0x02, 0x79, 0x00, 0xfd, 0x5c,
	0x19, 0xcb, 0x52, 0x3a, 0x53, 0x09, 0x8f, 0x7a, 0xa8, 0x03, 0x0f, 0x1d, 0xaa, 0x04, 0x1f, 0x34,
	0x53, 0x85, 0xb4, 0xfa, 0x32, 0x02, 0x54, 0x96, 0x22, 0xb9, 0x07, 0xbd, 0x84, 0x59, 0x4e, 0xad,
	0xc8, 0x78, 0xd4, 0x47, 0x5d, 0xd7, 0x01, 0xa7, 0x22, 0xe3, 0xe4, 0xff, 0xb0, 0x65, 0x44, 0xc2,
	0xa9, 0x3a, 0xa7, 0xe1, 0x85, 0x9b, 0x68, 0xb1, 0xe9, 0xd0, 0xb7, 0xe7, 0x53, 0xff, 0xce, 0x27,
	0xb0, 0xa5, 0xb4, 0x98, 0x0b, 0xc9, 0x52, 0x2a, 0x64, 0xc2, 0x97, 0xd1, 0x00, 0x9d, 0x1e, 0x94,
	0xe8, 0xb1, 0x03, 0x47, 0xbf, 0x36, 0xa0, 0x33, 0x2d, 0xb2, 0x8c, 0xe9, 0x4b, 0xe7, 0x72, 0xca,
	0xe5, 0xdc, 0x2e, 0x42, 0x6c, 0x83, 0x84, 0xc1, 0x74, 0x0f, 0x69, 0x22, 0x01, 0x7e, 0x93, 0x7d,
	0xe8, 0x64, 0x42, 0x52, 0x97, 0x88, 0x96, 0x37, 0xce, 0x84, 0x7c, 0xcd, 0x6c, 0xa5, 0x50, 0x12,
	0xa3, 0x1c, 0x14, 0x4a, 0xa2, 0x82, 0x2d, 0xf1, 0xc4, 0x46, 0x50, 0xb0, 0x65, 0x79, 0xc2, 0x29,
	0x94, 0xc4, 0x50, 0x07, 0x85, 0x92, 0xa3, 0x3f, 0x9a, 0xd0, 0x3f, 0xd5, 0x4c, 0x1a, 0x61, 0xa7,
	0x56, 0xe5, 0x55, 0x52, 0x1b, 0xb5, 0xa4, 0xde, 0x07, 0x50, 0x92, 0x1b, 0xab, 0x72, 0x2a, 0x12,
	0x7c, 0x61, 0x2f, 0xee, 0x05, 0xe4, 0x38, 0xa9, 0x32, 0xdb, 0xaa, 0x65, 0xf6, 0x29, 0xec, 0x30,
	0xad, 0xc5, 0x05, 0x4b, 0xe9, 0x2a, 0xc8, 0x6d, 0x34, 0xd8, 0x0e, 0x8a, 0xa3, 0x32, 0xd6, 0x13,
	0xd8, 0x4d, 0x78, 0xce, 0xb4, 0x2d, 0x34, 0xaf, 0x59, 0xfb, 0x42, 0xd9, 0xa9, 0x54, 0x47, 0xb5,
	0xdc, 0x08, 0x43, 0x73, 0xa6, 0xb9, 0xb4, 0xd4, 0xbd, 0x01, 0x5d, 0xea, 0xc6, 0x9b, 0xc2, 0xbc,
	0x43, 0x10, 0x1d, 0xf9, 0x04, 0x86, 0xcc, 0x98, 0x22, 0xe3, 0x09, 0x35, 0xb3, 0x05, 0x4f, 0x8a,
	0x94, 0x63, 0x3d, 0x75, 0xe3, 0xed, 0x80, 0x4f, 0x03, 0x5c, 0x16, 0x7b, 0x77, 0xad, 0xd8, 0x7b,
	0x55, 0xb1, 0x8f, 0x7e, 0x69, 0x55, 0x71, 0x3a, 0x96, 0xe7, 0xea, 0x5a, 0x4c, 0x1a, 0xd7, 0x63,
	0x72, 0x1f, 0xc0, 0x2c, 0x94, 0xb6, 0x14, 0x23, 0x13, 0x42, 0x86, 0xc8, 0x1b, 0x17, 0x9e, 0x7b,
	0xd0, 0x4b, 0x95, 0x9c, 0xd3, 0x5a, 0xdc, 0xba, 0x0e, 0x40, 0xe5, 0x08, 0xba, 0xae, 0x69, 0x8c,
	0x98, 0xcb, 0x10, 0xb2, 0x4a, 0x76, 0x5d, 0x30, 0x53, 0xa9, 0xd2, 0x18, 0x9d, 0x41, 0xec, 0x05,
	0x77, 0x9b, 0xe5, 0x4b, 0x4b, 0xbd, 0xea, 0x16, 0xaa, 0x7a, 0x0e, 0x39, 0x44, 0xf5, 0x18, 0xfa,
	0x09, 0x37, 0x33, 0x2d, 0x72, 0xd7, 0xdb, 0xa1, 0xab, 0xea, 0x90, 0x4b, 0x81, 0xca, 0xb9, 0x66,
	0x56, 0x69, 0x5a, 0x73, 0xcb, 0xb7, 0xda, 0x4e, 0xa9, 0x7a, 0x5b, 0xb9, 0xf7, 0x18, 0x06, 0x95,
	0x3d, 0xfa, 0xe0, 0x1b, 0x6f, 0xb3, 0x04, 0xd1, 0x8f, 0x87, 0x50, 0xc9, 0xb4, 0xd0, 0x69, 0xe8,
	0xbf, 0x7e, 0x89, 0xfd, 0xa0, 0x53, 0xf2, 0x35, 0x0c, 0xac, 0x0f, 0x2a, 0x26, 0xd2, 0x44, 0xfd,
	0x71, 0xeb, 0xa0, 0xff, 0xe2, 0xde, 0xfa, 0x98, 0xaa, 0x2a, 0x34, 0xde, 0xb4, 0x2b, 0xc1, 0x8c,
	0x7e, 0xdf, 0x84, 0xee, 0x09, 0x93, 0xbc, 0xb8, 0xe0, 0xfa, 0x4a, 0xf1, 0x0e, 0x42, 0xf1, 0x8e,
	0xa1, 0x2f, 0xa4, 0xb1, 0xba, 0x98, 0xa1, 0xf3, 0x3e, 0x15, 0x75, 0xc8, 0xbd, 0xd3, 0xf7, 0x38,
	0xba, 0x62, 0xa2, 0xd6, 0xb8, 0xe5, 0x4c, 0x3c, 0xe6, 0x3c, 0x31, 0xb5, 0xae, 0x6d, 0xdf, 0xd8,
	0xb5, 0x1b, 0xb5, 0xae, 0xfd, 0x02, 0xa2, 0x33, 0x3e, 0x17, 0x92, 0xce, 0x98, 0x4e, 0x70, 0x34,
	0x24, 0x42, 0x73, 0x7f, 0xbb, 0x1f, 0x73, 0x77, 0x50, 0x7f, 0x18, 0xd4, 0x47, 0xa5, 0xd6, 0x45,
	0xd5, 0x9f, 0x2c, 0x47, 0x68, 0x07, 0x69, 0x37, 0x11, 0xfc, 0x36, 0xcc, 0xd1, 0xa7, 0xb0, 0xe3,
	0x8d, 0xcc, 0x82, 0xe5, 0x3c, 0x8c, 0x9d, 0x2e, 0x1a, 0x6e, 0xa3, 0x62, 0xea, 0x70, 0x1c, 0x3c,
	0xe4, 0x23, 0xd8, 0xe6, 0x32, 0xb9, 0x62, 0xd9, 0xf3, 0x03, 0x8a, 0xcb, 0xa4, 0x66, 0xe7, 0xdc,
	0x50, 0xa9, 0xcf, 0x50, 0x37, 0xc6, 0x6f, 0x57, 0x69, 0x5a, 0x15, 0xf3, 0x05, 0x8e, 0xc6, 0x6e,
	0xec, 0x05, 0x72, 0x02, 0x8f, 0x2f, 0xb8, 0x3e, 0x63, 0x29, 0x0d, 0x59, 0x10, 0x4a, 0x52, 0x96,
	0x72, 0x6d, 0x69, 0x3d, 0xca, 0x7e, 0x58, 0x8e, 0xbd, 0xe9, 0x69, 0x65, 0xf9, 0xd2, 0x19, 0x1e,
	0xd7, 0x42, 0x7f, 0x0c, 0x0f, 0x03, 0x5d, 0xae, 0x79, 0x9d, 0xb2, 0x4e, 0x36, 0x40, 0xb2, 0xff,
	0x79, 0xc3, 0x77, 0x9a, 0xaf, 0xf8, 0xea, 0x54, 0xdf, 0xc3, 0xa3, 0x92, 0x4a, 0x19, 0xfb, 0x3e,
	0xae, 0x2d, 0xe4, 0x7a, 0x10, 0xb8, 0x94, 0xb1, 0x37, 0x93, 0x7d, 0x06, 0x24, 0x04, 0xb9, 0x5e,
	0x18, 0xdb, 0x58, 0x18, 0x43, 0x1f, 0xe5, 0x5a, 0x75, 0x3c, 0x87, 0x36, 0x36, 0xeb, 0x10, 0xff,
	0xb1, 0xf7, 0xaf, 0x17, 0x6f, 0x59, 0x9e, 0x93, 0xa9, 0x98, 0xcb, 0x18, 0x4d, 0xc9, 0x0b, 0xb8,
	0xad, 0x55, 0x21, 0x13, 0x76, 0xa6, 0x0a, 0x4b, 0xf9, 0x52, 0xb8, 0xe6, 0x2d, 0xa4, 0x8d, 0x76,
	0x30, 0x3f, 0xbb, 0x2b, 0xe5, 0xab, 0xa5, 0xb0, 0x87, 0x4e, 0x45, 0x9e, 0x01, 0xf1, 0xc3, 0xf0,
	0x8a, 0x47, 0xa4, 0x3e, 0x26, 0xeb, 0x3e, 0x7c, 0x09, 0x77, 0x43, 0x40, 0x6e, 0x38, 0xb5, 0x8b,
	0xa7, 0xf6, 0xbd, 0xc1, 0xd1, 0xda, 0xd9, 0x67, 0x40, 0x70, 0x4a, 0xf3, 0x2b, 0x87, 0xf6, 0xfc,
	0x55, 0x5e, 0x73, 0xf3, 0x55, 0x37, 0x9c, 0xba, 0x5d, 0xbf, 0xea, 0xe5, 0xda, 0xd9, 0xaf, 0xa0,
	0x6c, 0x68, 0x2a, 0xe4, 0xb9, 0x8a, 0xee, 0x60, 0x10, 0xdf, 0x37, 0x01, 0xdc, 0xec, 0x8d, 0xfb,
	0xb6, 0x36, 0x88, 0x0f, 0x60, 0x18, 0xee, 0xce, 0x8a, 0xd4, 0x0a, 0x3a, 0x2b, 0x78, 0xb4, 0x8f,
	0x25, 0xbb, 0xe5, 0xf1, 0x13, 0x07, 0x1f, 0x16, 0xb8, 0x2b, 0x58, 0xcd, 0x2e, 0x78, 0x4a, 0x33,
	0xb7, 0x2b, 0x44, 0x7e, 0x57, 0xf0, 0xd0, 0x89, 0xdb, 0x15, 0x56, 0x06, 0x38, 0x45, 0xee, 0xd6,
	0x0d, 0x4e, 0x2f, 0x73, 0x3e, 0xfa, 0xab, 0x05, 0x6d, 0x97, 0x44, 0xf2, 0x23, 0xec, 0x61, 0xce,
	0x64, 0x91, 0x9d, 0x71, 0x4d, 0x79, 0xca, 0x33, 0x2e, 0xad, 0x89, 0x1a, 0x38, 0xbe, 0x9e, 0xfc,
	0x6b, 0x05, 0x4c, 0x5e, 0x79, 0xeb, 0x98, 0x38, 0x8a, 0x37, 0xc8, 0x10, 0x20, 0x53, 0x11, 0x9f,
	0x69, 0x26, 0x67, 0x8b, 0x15, 0x71, 0xf3, 0x83, 0x89, 0xbf, 0x41, 0x86, 0x35, 0x62, 0xab, 0x7e,
	0x62, 0x3a, 0x59, 0x11, 0xb7, 0x3e, 0x98, 0xf8, 0x14, 0x19, 0x2a, 0xe2, 0x29, 0x10, 0x1f, 0x0a,
	0x96, 0xf1, 0x15, 0x6d, 0xfb, 0x43, 0x68, 0x87, 0x18, 0x08, 0x96, 0xf1, 0x92, 0x74, 0xf4, 0x1d,
	0x74, 0xc2, 0x37, 0xce, 0x26, 0xbe, 0xb4, 0xd5, 0x42, 0xc2, 0x97, 0x96, 0x7c, 0x0a, 0x3b, 0x33,
	0x25, 0x0d, 0x9f, 0x15, 0xd6, 0x55, 0x9b, 0xef, 0x1c, 0xbf, 0x39, 0x0d, 0x6b, 0x0a, 0x6c, 0x9b,
	0xd1, 0xcf, 0x0d, 0x68, 0xbd, 0xe6, 0x73, 0xf2, 0x1c, 0x3a, 0xc6, 0x2f, 0x61, 0x61, 0x19, 0xde,
	0xbf, 0xfe, 0xba, 0xb0, 0xa3, 0xc5, 0xa5, 0x1d, 0xf9, 0x1c, 0x7a, 0x59, 0x78, 0x71, 0x99, 0x82,
	0xe8, 0x7d, 0x2e, 0xc5, 0x2b, 0x53, 0xdc, 0x55, 0xdd, 0x74, 0x0d, 0xbf, 0x76, 0x2f, 0x8c, 0x7e,
	0x6b, 0x42, 0xdb, 0xed, 0xdb, 0x8e, 0x36, 0x0d, 0x8b, 0x76, 0x59, 0x32, 0x6b, 0xb4, 0xe5, 0x26,
	0x1e, 0xaf, 0x4c, 0xeb, 0x1e, 0x34, 0xff, 0xa3, 0x07, 0x1f, 0x43, 0x3b, 0xe5, 0xf3, 0x32, 0xcd,
	0xbb, 0x6b, 0xb7, 0xf0, 0x79, 0x8c, 0x06, 0x6e, 0x95, 0x75, 0x1b, 0x75, 0x61, 0x68, 0xc6, 0x8d,
	0x61, 0xf3, 0x72, 0x5b, 0x1b, 0x78, 0xf4, 0xc4, 0x83, 0x7e, 0x63, 0x77, 0x40, 0xf8, 0xe5, 0x05,
	0xc9, 0x79, 0x5c, 0x48, 0x61, 0x4d, 0xf8, 0xc3, 0x79, 0xc1, 0x6d, 0x32, 0x29, 0x93, 0xf3, 0xc2,
	0xd1, 0x75, 0xc2, 0x96, 0x13, 0x64, 0xb2, 0x05, 0xcd, 0x6a, 0xc3, 0x68, 0x8a, 0xe4, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x48, 0x44, 0x36, 0x72, 0xca, 0x0c, 0x00, 0x00,
}
