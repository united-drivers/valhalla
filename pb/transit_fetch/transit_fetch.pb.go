// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transit_fetch.proto

/*
Package transit_fetch is a generated protocol buffer package.

It is generated from these files:
	transit_fetch.proto

It has these top-level messages:
	Transit_Fetch
*/
package transit_fetch

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

type Transit_Fetch_VehicleType int32

const (
	Transit_Fetch_kTram      Transit_Fetch_VehicleType = 0
	Transit_Fetch_kMetro     Transit_Fetch_VehicleType = 1
	Transit_Fetch_kRail      Transit_Fetch_VehicleType = 2
	Transit_Fetch_kBus       Transit_Fetch_VehicleType = 3
	Transit_Fetch_kFerry     Transit_Fetch_VehicleType = 4
	Transit_Fetch_kCableCar  Transit_Fetch_VehicleType = 5
	Transit_Fetch_kGondola   Transit_Fetch_VehicleType = 6
	Transit_Fetch_kFunicular Transit_Fetch_VehicleType = 7
)

var Transit_Fetch_VehicleType_name = map[int32]string{
	0: "kTram",
	1: "kMetro",
	2: "kRail",
	3: "kBus",
	4: "kFerry",
	5: "kCableCar",
	6: "kGondola",
	7: "kFunicular",
}
var Transit_Fetch_VehicleType_value = map[string]int32{
	"kTram":      0,
	"kMetro":     1,
	"kRail":      2,
	"kBus":       3,
	"kFerry":     4,
	"kCableCar":  5,
	"kGondola":   6,
	"kFunicular": 7,
}

func (x Transit_Fetch_VehicleType) Enum() *Transit_Fetch_VehicleType {
	p := new(Transit_Fetch_VehicleType)
	*p = x
	return p
}
func (x Transit_Fetch_VehicleType) String() string {
	return proto.EnumName(Transit_Fetch_VehicleType_name, int32(x))
}
func (x *Transit_Fetch_VehicleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Transit_Fetch_VehicleType_value, data, "Transit_Fetch_VehicleType")
	if err != nil {
		return err
	}
	*x = Transit_Fetch_VehicleType(value)
	return nil
}
func (Transit_Fetch_VehicleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTransitFetch, []int{0, 0}
}

type Transit_Fetch struct {
	Stops            []*Transit_Fetch_Stop     `protobuf:"bytes,1,rep,name=stops" json:"stops,omitempty"`
	StopPairs        []*Transit_Fetch_StopPair `protobuf:"bytes,2,rep,name=stop_pairs,json=stopPairs" json:"stop_pairs,omitempty"`
	Routes           []*Transit_Fetch_Route    `protobuf:"bytes,3,rep,name=routes" json:"routes,omitempty"`
	Shapes           []*Transit_Fetch_Shape    `protobuf:"bytes,4,rep,name=shapes" json:"shapes,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Transit_Fetch) Reset()                    { *m = Transit_Fetch{} }
func (m *Transit_Fetch) String() string            { return proto.CompactTextString(m) }
func (*Transit_Fetch) ProtoMessage()               {}
func (*Transit_Fetch) Descriptor() ([]byte, []int) { return fileDescriptorTransitFetch, []int{0} }

func (m *Transit_Fetch) GetStops() []*Transit_Fetch_Stop {
	if m != nil {
		return m.Stops
	}
	return nil
}

func (m *Transit_Fetch) GetStopPairs() []*Transit_Fetch_StopPair {
	if m != nil {
		return m.StopPairs
	}
	return nil
}

func (m *Transit_Fetch) GetRoutes() []*Transit_Fetch_Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *Transit_Fetch) GetShapes() []*Transit_Fetch_Shape {
	if m != nil {
		return m.Shapes
	}
	return nil
}

type Transit_Fetch_Stop struct {
	Lon                *float32 `protobuf:"fixed32,1,opt,name=lon" json:"lon,omitempty"`
	Lat                *float32 `protobuf:"fixed32,2,opt,name=lat" json:"lat,omitempty"`
	Graphid            *uint64  `protobuf:"varint,3,opt,name=graphid" json:"graphid,omitempty"`
	Name               *string  `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	OnestopId          *string  `protobuf:"bytes,5,opt,name=onestop_id,json=onestopId" json:"onestop_id,omitempty"`
	OsmWayId           *uint64  `protobuf:"varint,6,opt,name=osm_way_id,json=osmWayId" json:"osm_way_id,omitempty"`
	Timezone           *string  `protobuf:"bytes,8,opt,name=timezone" json:"timezone,omitempty"`
	WheelchairBoarding *bool    `protobuf:"varint,9,opt,name=wheelchair_boarding,json=wheelchairBoarding" json:"wheelchair_boarding,omitempty"`
	XXX_unrecognized   []byte   `json:"-"`
}

func (m *Transit_Fetch_Stop) Reset()         { *m = Transit_Fetch_Stop{} }
func (m *Transit_Fetch_Stop) String() string { return proto.CompactTextString(m) }
func (*Transit_Fetch_Stop) ProtoMessage()    {}
func (*Transit_Fetch_Stop) Descriptor() ([]byte, []int) {
	return fileDescriptorTransitFetch, []int{0, 0}
}

func (m *Transit_Fetch_Stop) GetLon() float32 {
	if m != nil && m.Lon != nil {
		return *m.Lon
	}
	return 0
}

func (m *Transit_Fetch_Stop) GetLat() float32 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *Transit_Fetch_Stop) GetGraphid() uint64 {
	if m != nil && m.Graphid != nil {
		return *m.Graphid
	}
	return 0
}

func (m *Transit_Fetch_Stop) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Transit_Fetch_Stop) GetOnestopId() string {
	if m != nil && m.OnestopId != nil {
		return *m.OnestopId
	}
	return ""
}

func (m *Transit_Fetch_Stop) GetOsmWayId() uint64 {
	if m != nil && m.OsmWayId != nil {
		return *m.OsmWayId
	}
	return 0
}

func (m *Transit_Fetch_Stop) GetTimezone() string {
	if m != nil && m.Timezone != nil {
		return *m.Timezone
	}
	return ""
}

func (m *Transit_Fetch_Stop) GetWheelchairBoarding() bool {
	if m != nil && m.WheelchairBoarding != nil {
		return *m.WheelchairBoarding
	}
	return false
}

type Transit_Fetch_StopPair struct {
	BikesAllowed            *bool    `protobuf:"varint,1,opt,name=bikes_allowed,json=bikesAllowed" json:"bikes_allowed,omitempty"`
	BlockId                 *uint32  `protobuf:"varint,2,opt,name=block_id,json=blockId" json:"block_id,omitempty"`
	DestinationArrivalTime  *uint32  `protobuf:"varint,3,opt,name=destination_arrival_time,json=destinationArrivalTime" json:"destination_arrival_time,omitempty"`
	DestinationGraphid      *uint64  `protobuf:"varint,4,opt,name=destination_graphid,json=destinationGraphid" json:"destination_graphid,omitempty"`
	DestinationOnestopId    *string  `protobuf:"bytes,5,opt,name=destination_onestop_id,json=destinationOnestopId" json:"destination_onestop_id,omitempty"`
	OperatedByOnestopId     *string  `protobuf:"bytes,6,opt,name=operated_by_onestop_id,json=operatedByOnestopId" json:"operated_by_onestop_id,omitempty"`
	OriginDepartureTime     *uint32  `protobuf:"varint,7,opt,name=origin_departure_time,json=originDepartureTime" json:"origin_departure_time,omitempty"`
	OriginGraphid           *uint64  `protobuf:"varint,8,opt,name=origin_graphid,json=originGraphid" json:"origin_graphid,omitempty"`
	OriginOnestopId         *string  `protobuf:"bytes,9,opt,name=origin_onestop_id,json=originOnestopId" json:"origin_onestop_id,omitempty"`
	RouteIndex              *uint32  `protobuf:"varint,10,opt,name=route_index,json=routeIndex" json:"route_index,omitempty"`
	ServiceAddedDates       []uint32 `protobuf:"varint,11,rep,name=service_added_dates,json=serviceAddedDates" json:"service_added_dates,omitempty"`
	ServiceDaysOfWeek       []bool   `protobuf:"varint,12,rep,name=service_days_of_week,json=serviceDaysOfWeek" json:"service_days_of_week,omitempty"`
	ServiceEndDate          *uint32  `protobuf:"varint,13,opt,name=service_end_date,json=serviceEndDate" json:"service_end_date,omitempty"`
	ServiceExceptDates      []uint32 `protobuf:"varint,14,rep,name=service_except_dates,json=serviceExceptDates" json:"service_except_dates,omitempty"`
	ServiceStartDate        *uint32  `protobuf:"varint,15,opt,name=service_start_date,json=serviceStartDate" json:"service_start_date,omitempty"`
	TripHeadsign            *string  `protobuf:"bytes,16,opt,name=trip_headsign,json=tripHeadsign" json:"trip_headsign,omitempty"`
	TripId                  *uint32  `protobuf:"varint,17,opt,name=trip_id,json=tripId" json:"trip_id,omitempty"`
	WheelchairAccessible    *bool    `protobuf:"varint,18,opt,name=wheelchair_accessible,json=wheelchairAccessible" json:"wheelchair_accessible,omitempty"`
	ShapeId                 *uint32  `protobuf:"varint,20,opt,name=shape_id,json=shapeId" json:"shape_id,omitempty"`
	OriginDistTraveled      *float32 `protobuf:"fixed32,21,opt,name=origin_dist_traveled,json=originDistTraveled" json:"origin_dist_traveled,omitempty"`
	DestinationDistTraveled *float32 `protobuf:"fixed32,22,opt,name=destination_dist_traveled,json=destinationDistTraveled" json:"destination_dist_traveled,omitempty"`
	FrequencyEndTime        *uint32  `protobuf:"varint,23,opt,name=frequency_end_time,json=frequencyEndTime" json:"frequency_end_time,omitempty"`
	FrequencyHeadwaySeconds *uint32  `protobuf:"varint,24,opt,name=frequency_headway_seconds,json=frequencyHeadwaySeconds" json:"frequency_headway_seconds,omitempty"`
	XXX_unrecognized        []byte   `json:"-"`
}

func (m *Transit_Fetch_StopPair) Reset()         { *m = Transit_Fetch_StopPair{} }
func (m *Transit_Fetch_StopPair) String() string { return proto.CompactTextString(m) }
func (*Transit_Fetch_StopPair) ProtoMessage()    {}
func (*Transit_Fetch_StopPair) Descriptor() ([]byte, []int) {
	return fileDescriptorTransitFetch, []int{0, 1}
}

func (m *Transit_Fetch_StopPair) GetBikesAllowed() bool {
	if m != nil && m.BikesAllowed != nil {
		return *m.BikesAllowed
	}
	return false
}

func (m *Transit_Fetch_StopPair) GetBlockId() uint32 {
	if m != nil && m.BlockId != nil {
		return *m.BlockId
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetDestinationArrivalTime() uint32 {
	if m != nil && m.DestinationArrivalTime != nil {
		return *m.DestinationArrivalTime
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetDestinationGraphid() uint64 {
	if m != nil && m.DestinationGraphid != nil {
		return *m.DestinationGraphid
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetDestinationOnestopId() string {
	if m != nil && m.DestinationOnestopId != nil {
		return *m.DestinationOnestopId
	}
	return ""
}

func (m *Transit_Fetch_StopPair) GetOperatedByOnestopId() string {
	if m != nil && m.OperatedByOnestopId != nil {
		return *m.OperatedByOnestopId
	}
	return ""
}

func (m *Transit_Fetch_StopPair) GetOriginDepartureTime() uint32 {
	if m != nil && m.OriginDepartureTime != nil {
		return *m.OriginDepartureTime
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetOriginGraphid() uint64 {
	if m != nil && m.OriginGraphid != nil {
		return *m.OriginGraphid
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetOriginOnestopId() string {
	if m != nil && m.OriginOnestopId != nil {
		return *m.OriginOnestopId
	}
	return ""
}

func (m *Transit_Fetch_StopPair) GetRouteIndex() uint32 {
	if m != nil && m.RouteIndex != nil {
		return *m.RouteIndex
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetServiceAddedDates() []uint32 {
	if m != nil {
		return m.ServiceAddedDates
	}
	return nil
}

func (m *Transit_Fetch_StopPair) GetServiceDaysOfWeek() []bool {
	if m != nil {
		return m.ServiceDaysOfWeek
	}
	return nil
}

func (m *Transit_Fetch_StopPair) GetServiceEndDate() uint32 {
	if m != nil && m.ServiceEndDate != nil {
		return *m.ServiceEndDate
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetServiceExceptDates() []uint32 {
	if m != nil {
		return m.ServiceExceptDates
	}
	return nil
}

func (m *Transit_Fetch_StopPair) GetServiceStartDate() uint32 {
	if m != nil && m.ServiceStartDate != nil {
		return *m.ServiceStartDate
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetTripHeadsign() string {
	if m != nil && m.TripHeadsign != nil {
		return *m.TripHeadsign
	}
	return ""
}

func (m *Transit_Fetch_StopPair) GetTripId() uint32 {
	if m != nil && m.TripId != nil {
		return *m.TripId
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetWheelchairAccessible() bool {
	if m != nil && m.WheelchairAccessible != nil {
		return *m.WheelchairAccessible
	}
	return false
}

func (m *Transit_Fetch_StopPair) GetShapeId() uint32 {
	if m != nil && m.ShapeId != nil {
		return *m.ShapeId
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetOriginDistTraveled() float32 {
	if m != nil && m.OriginDistTraveled != nil {
		return *m.OriginDistTraveled
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetDestinationDistTraveled() float32 {
	if m != nil && m.DestinationDistTraveled != nil {
		return *m.DestinationDistTraveled
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetFrequencyEndTime() uint32 {
	if m != nil && m.FrequencyEndTime != nil {
		return *m.FrequencyEndTime
	}
	return 0
}

func (m *Transit_Fetch_StopPair) GetFrequencyHeadwaySeconds() uint32 {
	if m != nil && m.FrequencyHeadwaySeconds != nil {
		return *m.FrequencyHeadwaySeconds
	}
	return 0
}

type Transit_Fetch_Route struct {
	Name                *string                    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	OnestopId           *string                    `protobuf:"bytes,2,opt,name=onestop_id,json=onestopId" json:"onestop_id,omitempty"`
	OperatedByName      *string                    `protobuf:"bytes,3,opt,name=operated_by_name,json=operatedByName" json:"operated_by_name,omitempty"`
	OperatedByOnestopId *string                    `protobuf:"bytes,4,opt,name=operated_by_onestop_id,json=operatedByOnestopId" json:"operated_by_onestop_id,omitempty"`
	OperatedByWebsite   *string                    `protobuf:"bytes,5,opt,name=operated_by_website,json=operatedByWebsite" json:"operated_by_website,omitempty"`
	RouteColor          *uint32                    `protobuf:"varint,6,opt,name=route_color,json=routeColor" json:"route_color,omitempty"`
	RouteDesc           *string                    `protobuf:"bytes,7,opt,name=route_desc,json=routeDesc" json:"route_desc,omitempty"`
	RouteLongName       *string                    `protobuf:"bytes,8,opt,name=route_long_name,json=routeLongName" json:"route_long_name,omitempty"`
	RouteTextColor      *uint32                    `protobuf:"varint,9,opt,name=route_text_color,json=routeTextColor" json:"route_text_color,omitempty"`
	VehicleType         *Transit_Fetch_VehicleType `protobuf:"varint,10,opt,name=vehicle_type,json=vehicleType,enum=valhalla.mjolnir.Transit_Fetch_VehicleType" json:"vehicle_type,omitempty"`
	XXX_unrecognized    []byte                     `json:"-"`
}

func (m *Transit_Fetch_Route) Reset()         { *m = Transit_Fetch_Route{} }
func (m *Transit_Fetch_Route) String() string { return proto.CompactTextString(m) }
func (*Transit_Fetch_Route) ProtoMessage()    {}
func (*Transit_Fetch_Route) Descriptor() ([]byte, []int) {
	return fileDescriptorTransitFetch, []int{0, 2}
}

func (m *Transit_Fetch_Route) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Transit_Fetch_Route) GetOnestopId() string {
	if m != nil && m.OnestopId != nil {
		return *m.OnestopId
	}
	return ""
}

func (m *Transit_Fetch_Route) GetOperatedByName() string {
	if m != nil && m.OperatedByName != nil {
		return *m.OperatedByName
	}
	return ""
}

func (m *Transit_Fetch_Route) GetOperatedByOnestopId() string {
	if m != nil && m.OperatedByOnestopId != nil {
		return *m.OperatedByOnestopId
	}
	return ""
}

func (m *Transit_Fetch_Route) GetOperatedByWebsite() string {
	if m != nil && m.OperatedByWebsite != nil {
		return *m.OperatedByWebsite
	}
	return ""
}

func (m *Transit_Fetch_Route) GetRouteColor() uint32 {
	if m != nil && m.RouteColor != nil {
		return *m.RouteColor
	}
	return 0
}

func (m *Transit_Fetch_Route) GetRouteDesc() string {
	if m != nil && m.RouteDesc != nil {
		return *m.RouteDesc
	}
	return ""
}

func (m *Transit_Fetch_Route) GetRouteLongName() string {
	if m != nil && m.RouteLongName != nil {
		return *m.RouteLongName
	}
	return ""
}

func (m *Transit_Fetch_Route) GetRouteTextColor() uint32 {
	if m != nil && m.RouteTextColor != nil {
		return *m.RouteTextColor
	}
	return 0
}

func (m *Transit_Fetch_Route) GetVehicleType() Transit_Fetch_VehicleType {
	if m != nil && m.VehicleType != nil {
		return *m.VehicleType
	}
	return Transit_Fetch_kTram
}

type Transit_Fetch_Shape struct {
	ShapeId          *uint32 `protobuf:"varint,1,opt,name=shape_id,json=shapeId" json:"shape_id,omitempty"`
	EncodedShape     []byte  `protobuf:"bytes,2,opt,name=encoded_shape,json=encodedShape" json:"encoded_shape,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Transit_Fetch_Shape) Reset()         { *m = Transit_Fetch_Shape{} }
func (m *Transit_Fetch_Shape) String() string { return proto.CompactTextString(m) }
func (*Transit_Fetch_Shape) ProtoMessage()    {}
func (*Transit_Fetch_Shape) Descriptor() ([]byte, []int) {
	return fileDescriptorTransitFetch, []int{0, 3}
}

func (m *Transit_Fetch_Shape) GetShapeId() uint32 {
	if m != nil && m.ShapeId != nil {
		return *m.ShapeId
	}
	return 0
}

func (m *Transit_Fetch_Shape) GetEncodedShape() []byte {
	if m != nil {
		return m.EncodedShape
	}
	return nil
}

func init() {
	proto.RegisterType((*Transit_Fetch)(nil), "valhalla.mjolnir.Transit_Fetch")
	proto.RegisterType((*Transit_Fetch_Stop)(nil), "valhalla.mjolnir.Transit_Fetch.Stop")
	proto.RegisterType((*Transit_Fetch_StopPair)(nil), "valhalla.mjolnir.Transit_Fetch.StopPair")
	proto.RegisterType((*Transit_Fetch_Route)(nil), "valhalla.mjolnir.Transit_Fetch.Route")
	proto.RegisterType((*Transit_Fetch_Shape)(nil), "valhalla.mjolnir.Transit_Fetch.Shape")
	proto.RegisterEnum("valhalla.mjolnir.Transit_Fetch_VehicleType", Transit_Fetch_VehicleType_name, Transit_Fetch_VehicleType_value)
}

func init() { proto.RegisterFile("transit_fetch.proto", fileDescriptorTransitFetch) }

var fileDescriptorTransitFetch = []byte{
	// 1083 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xe1, 0x6e, 0x1b, 0x45,
	0x10, 0xc6, 0xb1, 0x93, 0x9c, 0x27, 0xb6, 0x7b, 0x59, 0xa7, 0xcd, 0x35, 0x02, 0x61, 0xb5, 0x14,
	0x59, 0x50, 0x6c, 0xd4, 0x82, 0x84, 0x2a, 0xf1, 0x23, 0x69, 0xda, 0x34, 0x12, 0xb4, 0xe8, 0x62,
	0x51, 0x89, 0x3f, 0xa7, 0xf5, 0xed, 0xc4, 0x5e, 0x7c, 0xbe, 0x3d, 0x76, 0xd7, 0x4e, 0xcc, 0x9b,
	0xf0, 0x0c, 0xbc, 0x15, 0x0f, 0x82, 0xd0, 0xce, 0xdd, 0xd9, 0xe7, 0xa2, 0xaa, 0xfd, 0xb7, 0xfb,
	0x7d, 0xf3, 0xcd, 0xce, 0x8c, 0x67, 0xe6, 0x0c, 0x5d, 0xab, 0x79, 0x6a, 0xa4, 0x8d, 0xae, 0xd1,
	0xc6, 0xd3, 0x41, 0xa6, 0x95, 0x55, 0xcc, 0x5f, 0xf2, 0x64, 0xca, 0x93, 0x84, 0x0f, 0xe6, 0xbf,
	0xab, 0x24, 0x95, 0xfa, 0xc1, 0xbf, 0x3e, 0xb4, 0x47, 0x85, 0xe5, 0x4b, 0x67, 0xc9, 0x9e, 0xc1,
	0xae, 0xb1, 0x2a, 0x33, 0x41, 0xad, 0x57, 0xef, 0x1f, 0x3c, 0xf9, 0x62, 0xf0, 0xae, 0x66, 0xb0,
	0x65, 0x3f, 0xb8, 0xb2, 0x2a, 0x0b, 0x73, 0x09, 0xbb, 0x00, 0x70, 0x87, 0x28, 0xe3, 0x52, 0x9b,
	0x60, 0x87, 0x1c, 0xf4, 0x3f, 0xc6, 0xc1, 0x2f, 0x5c, 0xea, 0xb0, 0x69, 0x8a, 0x93, 0x61, 0x3f,
	0xc2, 0x9e, 0x56, 0x0b, 0x8b, 0x26, 0xa8, 0x93, 0x93, 0x47, 0x1f, 0x72, 0x12, 0x3a, 0xeb, 0xb0,
	0x10, 0x39, 0xb9, 0x99, 0xf2, 0x0c, 0x4d, 0xd0, 0xf8, 0x38, 0xf9, 0x95, 0xb3, 0x0e, 0x0b, 0xd1,
	0xc9, 0x3f, 0x35, 0x68, 0xb8, 0xa8, 0x98, 0x0f, 0xf5, 0x44, 0xa5, 0x41, 0xad, 0x57, 0xeb, 0xef,
	0x84, 0xee, 0x48, 0x08, 0xb7, 0xc1, 0x4e, 0x81, 0x70, 0xcb, 0x02, 0xd8, 0x9f, 0x68, 0x9e, 0x4d,
	0xa5, 0x08, 0xea, 0xbd, 0x5a, 0xbf, 0x11, 0x96, 0x57, 0xc6, 0xa0, 0x91, 0xf2, 0x39, 0x06, 0x8d,
	0x5e, 0xad, 0xdf, 0x0c, 0xe9, 0xcc, 0x3e, 0x03, 0x50, 0x29, 0x52, 0x91, 0xa4, 0x08, 0x76, 0x89,
	0x69, 0x16, 0xc8, 0xa5, 0x60, 0x9f, 0x02, 0x28, 0x33, 0x8f, 0x6e, 0xf8, 0xca, 0xd1, 0x7b, 0xe4,
	0xcf, 0x53, 0x66, 0xfe, 0x96, 0xaf, 0x2e, 0x05, 0x3b, 0x01, 0xcf, 0xca, 0x39, 0xfe, 0xa9, 0x52,
	0x0c, 0x3c, 0x92, 0xae, 0xef, 0x6c, 0x08, 0xdd, 0x9b, 0x29, 0x62, 0x12, 0x4f, 0xb9, 0xd4, 0xd1,
	0x58, 0x71, 0x2d, 0x64, 0x3a, 0x09, 0x9a, 0xbd, 0x5a, 0xdf, 0x0b, 0xd9, 0x86, 0x3a, 0x2b, 0x98,
	0x93, 0xbf, 0x3c, 0xf0, 0xca, 0xd2, 0xb3, 0x87, 0xd0, 0x1e, 0xcb, 0x19, 0x9a, 0x88, 0x27, 0x89,
	0xba, 0x41, 0x41, 0x29, 0x7b, 0x61, 0x8b, 0xc0, 0xd3, 0x1c, 0x63, 0xf7, 0xc1, 0x1b, 0x27, 0x2a,
	0x9e, 0xb9, 0xd0, 0x5c, 0x01, 0xda, 0xe1, 0x3e, 0xdd, 0x2f, 0x05, 0xfb, 0x01, 0x02, 0x81, 0xc6,
	0xca, 0x94, 0x5b, 0xa9, 0xd2, 0x88, 0x6b, 0x2d, 0x97, 0x3c, 0x89, 0x5c, 0x74, 0x54, 0x95, 0x76,
	0x78, 0xaf, 0xc2, 0x9f, 0xe6, 0xf4, 0x48, 0xce, 0x29, 0xee, 0xaa, 0xb2, 0x2c, 0x65, 0x83, 0x52,
	0x67, 0x15, 0xea, 0xa2, 0xa8, 0xea, 0x77, 0x50, 0x75, 0x15, 0xfd, 0xaf, 0x9a, 0x47, 0x15, 0xf6,
	0xcd, 0xba, 0xb0, 0x4f, 0xe1, 0x9e, 0xca, 0x50, 0x73, 0x8b, 0x22, 0x1a, 0xaf, 0xaa, 0xaa, 0x3d,
	0x52, 0x75, 0x4b, 0xf6, 0x6c, 0xb5, 0x11, 0x3d, 0x81, 0xbb, 0x4a, 0xcb, 0x89, 0x4c, 0x23, 0x81,
	0x19, 0xd7, 0x76, 0xa1, 0x31, 0x4f, 0x69, 0x9f, 0x52, 0xea, 0xe6, 0xe4, 0x79, 0xc9, 0x51, 0x3e,
	0x8f, 0xa0, 0x53, 0x68, 0xca, 0x54, 0x3c, 0x4a, 0xa5, 0x9d, 0xa3, 0x65, 0x16, 0x5f, 0xc1, 0x61,
	0x61, 0x56, 0x09, 0xa5, 0x49, 0xa1, 0xdc, 0xc9, 0x89, 0x4d, 0x18, 0x9f, 0xc3, 0x01, 0xf5, 0x75,
	0x24, 0x53, 0x81, 0xb7, 0x01, 0xd0, 0xe3, 0x40, 0xd0, 0xa5, 0x43, 0xd8, 0x00, 0xba, 0x06, 0xf5,
	0x52, 0xc6, 0x18, 0x71, 0x21, 0x50, 0x44, 0x82, 0xbb, 0xd1, 0x39, 0xe8, 0xd5, 0xfb, 0xed, 0xf0,
	0xb0, 0xa0, 0x4e, 0x1d, 0x73, 0xee, 0x08, 0x36, 0x84, 0xa3, 0xd2, 0x5e, 0xf0, 0x95, 0x89, 0xd4,
	0x75, 0x74, 0x83, 0x38, 0x0b, 0x5a, 0xbd, 0x7a, 0xdf, 0x5b, 0x0b, 0xce, 0xf9, 0xca, 0xbc, 0xb9,
	0x7e, 0x8b, 0x38, 0x63, 0x7d, 0xf0, 0x4b, 0x01, 0xa6, 0xb9, 0xfb, 0xa0, 0x4d, 0x61, 0x74, 0x0a,
	0xfc, 0x45, 0x4a, 0xbe, 0xd9, 0xb7, 0x1b, 0xd7, 0x78, 0x1b, 0x63, 0x66, 0x8b, 0x58, 0x3a, 0x14,
	0x0b, 0x2b, 0xad, 0x89, 0xca, 0x83, 0x79, 0x0c, 0x25, 0x1a, 0x19, 0xcb, 0x75, 0x2e, 0x08, 0xee,
	0x90, 0xf7, 0xf2, 0xd5, 0x2b, 0x47, 0x90, 0xff, 0x87, 0xd0, 0xb6, 0x5a, 0x66, 0xd1, 0x14, 0xb9,
	0x30, 0x72, 0x92, 0x06, 0x3e, 0xd5, 0xac, 0xe5, 0xc0, 0x57, 0x05, 0xc6, 0x8e, 0x61, 0x9f, 0x8c,
	0xa4, 0x08, 0x0e, 0xc9, 0xcf, 0x9e, 0xbb, 0x52, 0x17, 0xdc, 0xad, 0x0c, 0x09, 0x8f, 0x63, 0x34,
	0x46, 0x8e, 0x13, 0x0c, 0x18, 0xb5, 0xfb, 0xd1, 0x86, 0x3c, 0x5d, 0x73, 0xae, 0xed, 0x69, 0x2f,
	0x38, 0x77, 0x47, 0x79, 0xdb, 0xd3, 0xfd, 0x52, 0xb8, 0x6c, 0xcb, 0x06, 0x91, 0xc6, 0x46, 0x56,
	0xf3, 0x25, 0x26, 0x28, 0x82, 0xbb, 0xb4, 0x1e, 0x58, 0xd1, 0x1f, 0xd2, 0xd8, 0x51, 0xc1, 0xb0,
	0x67, 0x70, 0xbf, 0xda, 0xbd, 0xdb, 0xb2, 0x7b, 0x24, 0x3b, 0xae, 0x18, 0x6c, 0x69, 0x1f, 0x03,
	0xbb, 0xd6, 0xf8, 0xc7, 0x02, 0xd3, 0x78, 0x45, 0xbf, 0x03, 0xf5, 0xe2, 0x71, 0x5e, 0xa9, 0x35,
	0xf3, 0x22, 0x15, 0xd4, 0x88, 0xcf, 0xe0, 0xfe, 0xc6, 0xda, 0x95, 0xcb, 0x2d, 0x15, 0x83, 0xb1,
	0x4a, 0x85, 0x09, 0x02, 0x12, 0x1d, 0xaf, 0x0d, 0x5e, 0xe5, 0xfc, 0x55, 0x4e, 0x9f, 0xfc, 0x5d,
	0x87, 0x5d, 0xda, 0xa8, 0xeb, 0x1d, 0x56, 0x7b, 0xef, 0x0e, 0xdb, 0x79, 0x77, 0x87, 0xf5, 0xc1,
	0xaf, 0x8e, 0x1a, 0xc9, 0xeb, 0x64, 0xd4, 0xd9, 0x0c, 0xd9, 0x6b, 0xe7, 0xe8, 0xfd, 0x43, 0xd9,
	0x78, 0xff, 0x50, 0x0e, 0xa0, 0x5b, 0x15, 0xdd, 0xe0, 0xd8, 0x48, 0x8b, 0xc5, 0xf0, 0x1f, 0x6e,
	0x14, 0x6f, 0x73, 0x62, 0x33, 0x3d, 0xb1, 0x4a, 0x94, 0xa6, 0x71, 0x2f, 0xa7, 0xe7, 0xb9, 0x43,
	0x5c, 0x3a, 0xb9, 0x81, 0x40, 0x13, 0xd3, 0x68, 0x37, 0xc3, 0x26, 0x21, 0xe7, 0x68, 0x62, 0xf6,
	0x25, 0xdc, 0xc9, 0xe9, 0x44, 0xa5, 0x93, 0x3c, 0x9b, 0x7c, 0xf7, 0xb6, 0x09, 0xfe, 0x49, 0xa5,
	0x13, 0x4a, 0xa6, 0x0f, 0x7e, 0x6e, 0x67, 0xf1, 0xd6, 0x16, 0x8f, 0x35, 0xf3, 0x19, 0x21, 0x7c,
	0x84, 0xb7, 0x36, 0x7f, 0xf0, 0x35, 0xb4, 0x96, 0x38, 0x95, 0x71, 0x82, 0x91, 0x5d, 0x65, 0x48,
	0x03, 0xdd, 0x79, 0xf2, 0xf5, 0x87, 0xbe, 0x51, 0xbf, 0xe6, 0x9a, 0xd1, 0x2a, 0xc3, 0xf0, 0x60,
	0xb9, 0xb9, 0x9c, 0x5c, 0xc0, 0x2e, 0x7d, 0xbf, 0xb6, 0x3a, 0xb5, 0xb6, 0xdd, 0xa9, 0x0f, 0xa1,
	0x8d, 0x69, 0xac, 0xdc, 0x72, 0x20, 0x88, 0x7e, 0xb6, 0x56, 0xd8, 0x2a, 0x40, 0xd2, 0x3f, 0xd0,
	0x70, 0x50, 0x79, 0x84, 0x35, 0x61, 0x77, 0x36, 0xd2, 0x7c, 0xee, 0x7f, 0xc2, 0x00, 0xf6, 0x66,
	0x3f, 0xa3, 0xd5, 0xca, 0xaf, 0x11, 0x1c, 0x72, 0x99, 0xf8, 0x3b, 0xcc, 0x83, 0xc6, 0xec, 0x6c,
	0x61, 0xfc, 0x3a, 0x19, 0xbc, 0x44, 0xad, 0x57, 0x7e, 0x83, 0xb5, 0xa1, 0x39, 0x7b, 0xce, 0xc7,
	0x09, 0x3e, 0xe7, 0xda, 0xdf, 0x65, 0x2d, 0xf0, 0x66, 0x17, 0x2a, 0x15, 0x2a, 0xe1, 0xfe, 0x1e,
	0xeb, 0x00, 0xcc, 0x5e, 0x2e, 0x52, 0x19, 0x2f, 0x12, 0xae, 0xfd, 0xfd, 0xb3, 0xef, 0x7f, 0x7b,
	0x3a, 0x91, 0x76, 0xba, 0x18, 0x0f, 0x62, 0x35, 0x1f, 0x2e, 0x52, 0x69, 0x51, 0x7c, 0x23, 0xb4,
	0x5c, 0xa2, 0x36, 0xc3, 0xb2, 0x22, 0xc3, 0x6c, 0x3c, 0xdc, 0xfa, 0x3f, 0xf3, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x45, 0xcf, 0xe5, 0x62, 0xdf, 0x08, 0x00, 0x00,
}