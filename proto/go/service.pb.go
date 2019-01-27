// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package open_now

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Context_Situation int32

const (
	Context_UNKNOWN Context_Situation = 0
	Context_VEHICLE Context_Situation = 1
	Context_FOOT    Context_Situation = 2
)

var Context_Situation_name = map[int32]string{
	0: "UNKNOWN",
	1: "VEHICLE",
	2: "FOOT",
}
var Context_Situation_value = map[string]int32{
	"UNKNOWN": 0,
	"VEHICLE": 1,
	"FOOT":    2,
}

func (x Context_Situation) String() string {
	return proto.EnumName(Context_Situation_name, int32(x))
}
func (Context_Situation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{4, 0}
}

type Interest_Type int32

const (
	Interest_UNKNOWN    Interest_Type = 0
	Interest_AUTHORITY  Interest_Type = 1
	Interest_FOOD       Interest_Type = 2
	Interest_SHOPPING   Interest_Type = 3
	Interest_LODGING    Interest_Type = 4
	Interest_ATTRACTION Interest_Type = 5
)

var Interest_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "AUTHORITY",
	2: "FOOD",
	3: "SHOPPING",
	4: "LODGING",
	5: "ATTRACTION",
}
var Interest_Type_value = map[string]int32{
	"UNKNOWN":    0,
	"AUTHORITY":  1,
	"FOOD":       2,
	"SHOPPING":   3,
	"LODGING":    4,
	"ATTRACTION": 5,
}

func (x Interest_Type) String() string {
	return proto.EnumName(Interest_Type_name, int32(x))
}
func (Interest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{6, 0}
}

type Status struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Coordinates struct {
	Latitude             float64  `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinates) Reset()         { *m = Coordinates{} }
func (m *Coordinates) String() string { return proto.CompactTextString(m) }
func (*Coordinates) ProtoMessage()    {}
func (*Coordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{2}
}
func (m *Coordinates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinates.Unmarshal(m, b)
}
func (m *Coordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinates.Marshal(b, m, deterministic)
}
func (dst *Coordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinates.Merge(dst, src)
}
func (m *Coordinates) XXX_Size() int {
	return xxx_messageInfo_Coordinates.Size(m)
}
func (m *Coordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinates.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinates proto.InternalMessageInfo

func (m *Coordinates) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Coordinates) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type Position struct {
	ClientId             string       `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Situation            *Context     `protobuf:"bytes,2,opt,name=situation,proto3" json:"situation,omitempty"`
	Direction            float32      `protobuf:"fixed32,3,opt,name=direction,proto3" json:"direction,omitempty"`
	Coordinates          *Coordinates `protobuf:"bytes,4,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{3}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (dst *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(dst, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Position) GetSituation() *Context {
	if m != nil {
		return m.Situation
	}
	return nil
}

func (m *Position) GetDirection() float32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *Position) GetCoordinates() *Coordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

type Context struct {
	Situation            Context_Situation `protobuf:"varint,1,opt,name=situation,proto3,enum=open_now.Context_Situation" json:"situation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{4}
}
func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (dst *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(dst, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetSituation() Context_Situation {
	if m != nil {
		return m.Situation
	}
	return Context_UNKNOWN
}

type PointsOfInterest struct {
	Interests            []*Interest `protobuf:"bytes,1,rep,name=interests,proto3" json:"interests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PointsOfInterest) Reset()         { *m = PointsOfInterest{} }
func (m *PointsOfInterest) String() string { return proto.CompactTextString(m) }
func (*PointsOfInterest) ProtoMessage()    {}
func (*PointsOfInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{5}
}
func (m *PointsOfInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointsOfInterest.Unmarshal(m, b)
}
func (m *PointsOfInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointsOfInterest.Marshal(b, m, deterministic)
}
func (dst *PointsOfInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointsOfInterest.Merge(dst, src)
}
func (m *PointsOfInterest) XXX_Size() int {
	return xxx_messageInfo_PointsOfInterest.Size(m)
}
func (m *PointsOfInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_PointsOfInterest.DiscardUnknown(m)
}

var xxx_messageInfo_PointsOfInterest proto.InternalMessageInfo

func (m *PointsOfInterest) GetInterests() []*Interest {
	if m != nil {
		return m.Interests
	}
	return nil
}

type Interest struct {
	InterestId           string        `protobuf:"bytes,1,opt,name=interest_id,json=interestId,proto3" json:"interest_id,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PhotoRef             string        `protobuf:"bytes,4,opt,name=photo_ref,json=photoRef,proto3" json:"photo_ref,omitempty"`
	Type                 Interest_Type `protobuf:"varint,6,opt,name=type,proto3,enum=open_now.Interest_Type" json:"type,omitempty"`
	Coordinates          *Coordinates  `protobuf:"bytes,7,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Interest) Reset()         { *m = Interest{} }
func (m *Interest) String() string { return proto.CompactTextString(m) }
func (*Interest) ProtoMessage()    {}
func (*Interest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{6}
}
func (m *Interest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Interest.Unmarshal(m, b)
}
func (m *Interest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Interest.Marshal(b, m, deterministic)
}
func (dst *Interest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interest.Merge(dst, src)
}
func (m *Interest) XXX_Size() int {
	return xxx_messageInfo_Interest.Size(m)
}
func (m *Interest) XXX_DiscardUnknown() {
	xxx_messageInfo_Interest.DiscardUnknown(m)
}

var xxx_messageInfo_Interest proto.InternalMessageInfo

func (m *Interest) GetInterestId() string {
	if m != nil {
		return m.InterestId
	}
	return ""
}

func (m *Interest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Interest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Interest) GetPhotoRef() string {
	if m != nil {
		return m.PhotoRef
	}
	return ""
}

func (m *Interest) GetType() Interest_Type {
	if m != nil {
		return m.Type
	}
	return Interest_UNKNOWN
}

func (m *Interest) GetCoordinates() *Coordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

type DirectionsReq struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	InterestId           string   `protobuf:"bytes,2,opt,name=interest_id,json=interestId,proto3" json:"interest_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectionsReq) Reset()         { *m = DirectionsReq{} }
func (m *DirectionsReq) String() string { return proto.CompactTextString(m) }
func (*DirectionsReq) ProtoMessage()    {}
func (*DirectionsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{7}
}
func (m *DirectionsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectionsReq.Unmarshal(m, b)
}
func (m *DirectionsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectionsReq.Marshal(b, m, deterministic)
}
func (dst *DirectionsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectionsReq.Merge(dst, src)
}
func (m *DirectionsReq) XXX_Size() int {
	return xxx_messageInfo_DirectionsReq.Size(m)
}
func (m *DirectionsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectionsReq.DiscardUnknown(m)
}

var xxx_messageInfo_DirectionsReq proto.InternalMessageInfo

func (m *DirectionsReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *DirectionsReq) GetInterestId() string {
	if m != nil {
		return m.InterestId
	}
	return ""
}

type DirectionsResp struct {
	InterestId           string   `protobuf:"bytes,1,opt,name=interest_id,json=interestId,proto3" json:"interest_id,omitempty"`
	Legs                 [][]byte `protobuf:"bytes,2,rep,name=legs,proto3" json:"legs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectionsResp) Reset()         { *m = DirectionsResp{} }
func (m *DirectionsResp) String() string { return proto.CompactTextString(m) }
func (*DirectionsResp) ProtoMessage()    {}
func (*DirectionsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{8}
}
func (m *DirectionsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectionsResp.Unmarshal(m, b)
}
func (m *DirectionsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectionsResp.Marshal(b, m, deterministic)
}
func (dst *DirectionsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectionsResp.Merge(dst, src)
}
func (m *DirectionsResp) XXX_Size() int {
	return xxx_messageInfo_DirectionsResp.Size(m)
}
func (m *DirectionsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectionsResp.DiscardUnknown(m)
}

var xxx_messageInfo_DirectionsResp proto.InternalMessageInfo

func (m *DirectionsResp) GetInterestId() string {
	if m != nil {
		return m.InterestId
	}
	return ""
}

func (m *DirectionsResp) GetLegs() [][]byte {
	if m != nil {
		return m.Legs
	}
	return nil
}

type TransitStops struct {
	Stops                []*TransitStop `protobuf:"bytes,1,rep,name=stops,proto3" json:"stops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransitStops) Reset()         { *m = TransitStops{} }
func (m *TransitStops) String() string { return proto.CompactTextString(m) }
func (*TransitStops) ProtoMessage()    {}
func (*TransitStops) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{9}
}
func (m *TransitStops) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitStops.Unmarshal(m, b)
}
func (m *TransitStops) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitStops.Marshal(b, m, deterministic)
}
func (dst *TransitStops) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitStops.Merge(dst, src)
}
func (m *TransitStops) XXX_Size() int {
	return xxx_messageInfo_TransitStops.Size(m)
}
func (m *TransitStops) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitStops.DiscardUnknown(m)
}

var xxx_messageInfo_TransitStops proto.InternalMessageInfo

func (m *TransitStops) GetStops() []*TransitStop {
	if m != nil {
		return m.Stops
	}
	return nil
}

type TransitStop struct {
	Coordinates          *Coordinates `protobuf:"bytes,1,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TransitStop) Reset()         { *m = TransitStop{} }
func (m *TransitStop) String() string { return proto.CompactTextString(m) }
func (*TransitStop) ProtoMessage()    {}
func (*TransitStop) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_de05588f15297485, []int{10}
}
func (m *TransitStop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitStop.Unmarshal(m, b)
}
func (m *TransitStop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitStop.Marshal(b, m, deterministic)
}
func (dst *TransitStop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitStop.Merge(dst, src)
}
func (m *TransitStop) XXX_Size() int {
	return xxx_messageInfo_TransitStop.Size(m)
}
func (m *TransitStop) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitStop.DiscardUnknown(m)
}

var xxx_messageInfo_TransitStop proto.InternalMessageInfo

func (m *TransitStop) GetCoordinates() *Coordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "open_now.Status")
	proto.RegisterType((*Empty)(nil), "open_now.Empty")
	proto.RegisterType((*Coordinates)(nil), "open_now.Coordinates")
	proto.RegisterType((*Position)(nil), "open_now.Position")
	proto.RegisterType((*Context)(nil), "open_now.Context")
	proto.RegisterType((*PointsOfInterest)(nil), "open_now.PointsOfInterest")
	proto.RegisterType((*Interest)(nil), "open_now.Interest")
	proto.RegisterType((*DirectionsReq)(nil), "open_now.DirectionsReq")
	proto.RegisterType((*DirectionsResp)(nil), "open_now.DirectionsResp")
	proto.RegisterType((*TransitStops)(nil), "open_now.TransitStops")
	proto.RegisterType((*TransitStop)(nil), "open_now.TransitStop")
	proto.RegisterEnum("open_now.Context_Situation", Context_Situation_name, Context_Situation_value)
	proto.RegisterEnum("open_now.Interest_Type", Interest_Type_name, Interest_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreClient interface {
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
	GetPointsOfInterest(ctx context.Context, in *Position, opts ...grpc.CallOption) (*PointsOfInterest, error)
	GetDirections(ctx context.Context, in *DirectionsReq, opts ...grpc.CallOption) (*DirectionsResp, error)
	GetTransitStops(ctx context.Context, in *Position, opts ...grpc.CallOption) (*TransitStops, error)
}

type coreClient struct {
	cc *grpc.ClientConn
}

func NewCoreClient(cc *grpc.ClientConn) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/open_now.Core/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetPointsOfInterest(ctx context.Context, in *Position, opts ...grpc.CallOption) (*PointsOfInterest, error) {
	out := new(PointsOfInterest)
	err := c.cc.Invoke(ctx, "/open_now.Core/GetPointsOfInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetDirections(ctx context.Context, in *DirectionsReq, opts ...grpc.CallOption) (*DirectionsResp, error) {
	out := new(DirectionsResp)
	err := c.cc.Invoke(ctx, "/open_now.Core/GetDirections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetTransitStops(ctx context.Context, in *Position, opts ...grpc.CallOption) (*TransitStops, error) {
	out := new(TransitStops)
	err := c.cc.Invoke(ctx, "/open_now.Core/GetTransitStops", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
type CoreServer interface {
	GetStatus(context.Context, *Empty) (*Status, error)
	GetPointsOfInterest(context.Context, *Position) (*PointsOfInterest, error)
	GetDirections(context.Context, *DirectionsReq) (*DirectionsResp, error)
	GetTransitStops(context.Context, *Position) (*TransitStops, error)
}

func RegisterCoreServer(s *grpc.Server, srv CoreServer) {
	s.RegisterService(&_Core_serviceDesc, srv)
}

func _Core_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_now.Core/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetPointsOfInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetPointsOfInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_now.Core/GetPointsOfInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetPointsOfInterest(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetDirections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetDirections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_now.Core/GetDirections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetDirections(ctx, req.(*DirectionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetTransitStops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetTransitStops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_now.Core/GetTransitStops",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetTransitStops(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

var _Core_serviceDesc = grpc.ServiceDesc{
	ServiceName: "open_now.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Core_GetStatus_Handler,
		},
		{
			MethodName: "GetPointsOfInterest",
			Handler:    _Core_GetPointsOfInterest_Handler,
		},
		{
			MethodName: "GetDirections",
			Handler:    _Core_GetDirections_Handler,
		},
		{
			MethodName: "GetTransitStops",
			Handler:    _Core_GetTransitStops_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_de05588f15297485) }

var fileDescriptor_service_de05588f15297485 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x1d, 0xb7, 0xb1, 0xc7, 0x4d, 0x6b, 0x16, 0x01, 0x51, 0x8a, 0x44, 0xb4, 0xa7, 0x48,
	0x95, 0x52, 0x14, 0x0e, 0x08, 0x71, 0x40, 0x55, 0x92, 0xba, 0x16, 0x25, 0x8e, 0x36, 0x2e, 0x15,
	0xa7, 0x2a, 0xc4, 0xdb, 0xb2, 0x52, 0xea, 0x35, 0xde, 0x0d, 0xd0, 0x3f, 0xe0, 0x6f, 0xf8, 0x02,
	0xfe, 0x0d, 0x79, 0x13, 0xdb, 0x9b, 0xb4, 0x42, 0xbd, 0x8d, 0xdf, 0xbc, 0x99, 0x1d, 0xbf, 0x79,
	0x1a, 0x68, 0x0a, 0x9a, 0xfd, 0x60, 0x73, 0xda, 0x4b, 0x33, 0x2e, 0x39, 0xb2, 0x79, 0x4a, 0x93,
	0xab, 0x84, 0xff, 0xc4, 0x36, 0xec, 0x4e, 0xe5, 0x4c, 0x2e, 0x05, 0x6e, 0xc0, 0xce, 0xe8, 0x36,
	0x95, 0x77, 0xd8, 0x07, 0x77, 0xc0, 0x79, 0x16, 0xb3, 0x64, 0x26, 0xa9, 0x40, 0x6d, 0xb0, 0x17,
	0x33, 0xc9, 0xe4, 0x32, 0xa6, 0x2d, 0xa3, 0x63, 0x74, 0x0d, 0x52, 0x7e, 0xa3, 0x97, 0xe0, 0x2c,
	0x78, 0x72, 0xb3, 0x4a, 0x9a, 0x2a, 0x59, 0x01, 0xf8, 0x8f, 0x01, 0xf6, 0x84, 0x0b, 0x26, 0x19,
	0x4f, 0xd0, 0x21, 0x38, 0xf3, 0x05, 0xa3, 0x89, 0xbc, 0x62, 0xb1, 0xea, 0xe3, 0x10, 0x7b, 0x05,
	0x04, 0x31, 0x3a, 0x06, 0x47, 0x30, 0xb9, 0x9c, 0xe5, 0x4c, 0xd5, 0xc7, 0xed, 0x3f, 0xe9, 0x15,
	0x33, 0xf6, 0x06, 0x3c, 0x91, 0xf4, 0x97, 0x24, 0x15, 0x27, 0x7f, 0x38, 0x66, 0x19, 0x9d, 0xab,
	0x82, 0x7a, 0xc7, 0xe8, 0x9a, 0xa4, 0x02, 0xd0, 0x5b, 0x70, 0xe7, 0xd5, 0x1f, 0xb4, 0x2c, 0xd5,
	0xf0, 0x99, 0xde, 0xb0, 0x4c, 0x12, 0x9d, 0x89, 0x97, 0xd0, 0x58, 0x3f, 0x86, 0xde, 0xe9, 0x23,
	0xe5, 0xf3, 0xee, 0xf7, 0x0f, 0xef, 0x8d, 0xd4, 0x9b, 0x16, 0x14, 0x6d, 0x38, 0x7c, 0x0c, 0x4e,
	0x89, 0x23, 0x17, 0x1a, 0x17, 0xe3, 0x8f, 0xe3, 0xf0, 0x72, 0xec, 0xd5, 0xf2, 0x8f, 0xcf, 0xa3,
	0xb3, 0x60, 0x70, 0x3e, 0xf2, 0x0c, 0x64, 0x83, 0x75, 0x1a, 0x86, 0x91, 0x67, 0xe2, 0x21, 0x78,
	0x13, 0xce, 0x12, 0x29, 0xc2, 0xeb, 0x20, 0x91, 0x34, 0xa3, 0x42, 0xa2, 0xd7, 0xe0, 0xb0, 0x75,
	0x2c, 0x5a, 0x46, 0xa7, 0xde, 0x75, 0xfb, 0xa8, 0x7a, 0xbf, 0xa0, 0x91, 0x8a, 0x84, 0xff, 0x9a,
	0x60, 0x97, 0xe5, 0xaf, 0xc0, 0x2d, 0x32, 0x95, 0xe0, 0x50, 0x40, 0x41, 0x8c, 0x10, 0x58, 0xc9,
	0xec, 0x76, 0xb5, 0x35, 0x87, 0xa8, 0x18, 0x75, 0xc0, 0x8d, 0xa9, 0x98, 0x67, 0x2c, 0x2d, 0x75,
	0x75, 0x88, 0x0e, 0xe5, 0x5b, 0x4c, 0xbf, 0x71, 0xc9, 0xaf, 0x32, 0x7a, 0xad, 0x74, 0x75, 0x88,
	0xad, 0x00, 0x42, 0xaf, 0xd1, 0x11, 0x58, 0xf2, 0x2e, 0xa5, 0xad, 0x5d, 0xa5, 0xd6, 0x8b, 0xfb,
	0xd3, 0xf6, 0xa2, 0xbb, 0x94, 0x12, 0x45, 0xda, 0xde, 0x51, 0xe3, 0xd1, 0x3b, 0xba, 0x04, 0x2b,
	0x6f, 0xb3, 0x29, 0x6c, 0x13, 0x9c, 0x93, 0x8b, 0xe8, 0x2c, 0x24, 0x41, 0xf4, 0xa5, 0x94, 0x76,
	0xe8, 0x99, 0x68, 0x0f, 0xec, 0xe9, 0x59, 0x38, 0x99, 0x04, 0x63, 0xdf, 0xab, 0xe7, 0x35, 0xe7,
	0xe1, 0xd0, 0xcf, 0x3f, 0x2c, 0xb4, 0x0f, 0x70, 0x12, 0x45, 0xe4, 0x64, 0x10, 0x05, 0xe1, 0xd8,
	0xdb, 0xc1, 0x9f, 0xa0, 0x39, 0x2c, 0x2c, 0x24, 0x08, 0xfd, 0xfe, 0x7f, 0xcb, 0x6e, 0x09, 0x6c,
	0x6e, 0x0b, 0x8c, 0x47, 0xb0, 0xaf, 0xb7, 0x13, 0xe9, 0xa3, 0x76, 0xb2, 0xa0, 0x37, 0xa2, 0x65,
	0x76, 0xea, 0xdd, 0x3d, 0xa2, 0x62, 0xfc, 0x1e, 0xf6, 0xa2, 0x6c, 0x96, 0x08, 0x26, 0xa7, 0x92,
	0xa7, 0x02, 0x1d, 0xc1, 0x8e, 0xc8, 0x83, 0xb5, 0x27, 0x34, 0xc5, 0x34, 0x1a, 0x59, 0x71, 0xf0,
	0x29, 0xb8, 0x1a, 0xba, 0xad, 0xb9, 0xf1, 0x58, 0xcd, 0xfb, 0xbf, 0x4d, 0xb0, 0x06, 0x3c, 0xa3,
	0xb9, 0x2b, 0x7d, 0x2a, 0x57, 0x17, 0x03, 0x1d, 0x54, 0x95, 0xea, 0x72, 0xb4, 0xbd, 0x0a, 0x58,
	0x1f, 0x95, 0x1a, 0xf2, 0xe1, 0xa9, 0x4f, 0xe5, 0x3d, 0x7b, 0x6b, 0x5e, 0x2e, 0x4e, 0x44, 0xbb,
	0xad, 0x63, 0x9b, 0x7c, 0x5c, 0x43, 0x43, 0x68, 0xfa, 0x54, 0x56, 0x92, 0x22, 0xcd, 0x60, 0x1b,
	0x7b, 0x6b, 0xb7, 0x1e, 0x4e, 0x88, 0x14, 0xd7, 0xd0, 0x07, 0x38, 0xf0, 0xa9, 0xdc, 0x50, 0xf4,
	0xa1, 0x51, 0x9e, 0x3f, 0x28, 0xab, 0xc0, 0xb5, 0xaf, 0xbb, 0xea, 0x82, 0xbe, 0xf9, 0x17, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0x5d, 0x6f, 0x80, 0x52, 0x05, 0x00, 0x00,
}
