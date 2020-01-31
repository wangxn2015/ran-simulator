// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/trafficsim/trafficsim.proto

package trafficsim

import (
	context "context"
	fmt "fmt"
	types "github.com/onosproject/ran-simulator/api/types"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Change event type
type Type int32

const (
	// NONE indicates this response does not represent a modification of the Change
	Type_NONE Type = 0
	// ADDED is an event which occurs when a Change is added to the topology
	Type_ADDED Type = 1
	// UPDATED is an event which occurs when a Change is updated
	Type_UPDATED Type = 2
	// REMOVED is an event which occurs when a Change is removed from the configuration
	Type_REMOVED Type = 3
)

var Type_name = map[int32]string{
	0: "NONE",
	1: "ADDED",
	2: "UPDATED",
	3: "REMOVED",
}

var Type_value = map[string]int32{
	"NONE":    0,
	"ADDED":   1,
	"UPDATED": 2,
	"REMOVED": 3,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{0}
}

type MapLayoutRequest struct {
}

func (m *MapLayoutRequest) Reset()         { *m = MapLayoutRequest{} }
func (m *MapLayoutRequest) String() string { return proto.CompactTextString(m) }
func (*MapLayoutRequest) ProtoMessage()    {}
func (*MapLayoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{0}
}
func (m *MapLayoutRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MapLayoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MapLayoutRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MapLayoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapLayoutRequest.Merge(m, src)
}
func (m *MapLayoutRequest) XXX_Size() int {
	return m.Size()
}
func (m *MapLayoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MapLayoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MapLayoutRequest proto.InternalMessageInfo

type ListTowersRequest struct {
}

func (m *ListTowersRequest) Reset()         { *m = ListTowersRequest{} }
func (m *ListTowersRequest) String() string { return proto.CompactTextString(m) }
func (*ListTowersRequest) ProtoMessage()    {}
func (*ListTowersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{1}
}
func (m *ListTowersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListTowersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListTowersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListTowersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTowersRequest.Merge(m, src)
}
func (m *ListTowersRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListTowersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTowersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTowersRequest proto.InternalMessageInfo

type ListRoutesRequest struct {
	// subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and REMOVE) that occur
	// after all routes have been streamed to the client
	Subscribe bool `protobuf:"varint,1,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
	// option to request only changes that happen after the call
	WithoutReplay bool `protobuf:"varint,2,opt,name=withoutReplay,proto3" json:"withoutReplay,omitempty"`
}

func (m *ListRoutesRequest) Reset()         { *m = ListRoutesRequest{} }
func (m *ListRoutesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRoutesRequest) ProtoMessage()    {}
func (*ListRoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{2}
}
func (m *ListRoutesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListRoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListRoutesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListRoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRoutesRequest.Merge(m, src)
}
func (m *ListRoutesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListRoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRoutesRequest proto.InternalMessageInfo

func (m *ListRoutesRequest) GetSubscribe() bool {
	if m != nil {
		return m.Subscribe
	}
	return false
}

func (m *ListRoutesRequest) GetWithoutReplay() bool {
	if m != nil {
		return m.WithoutReplay
	}
	return false
}

type ListRoutesResponse struct {
	// route is the route change on which the event occurred
	Route *types.Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	// type is a qualification of the type of change being made
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=ran.trafficsim.Type" json:"type,omitempty"`
}

func (m *ListRoutesResponse) Reset()         { *m = ListRoutesResponse{} }
func (m *ListRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRoutesResponse) ProtoMessage()    {}
func (*ListRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{3}
}
func (m *ListRoutesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListRoutesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRoutesResponse.Merge(m, src)
}
func (m *ListRoutesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRoutesResponse proto.InternalMessageInfo

func (m *ListRoutesResponse) GetRoute() *types.Route {
	if m != nil {
		return m.Route
	}
	return nil
}

func (m *ListRoutesResponse) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_NONE
}

type ListUesRequest struct {
	// subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and REMOVE) that occur
	// after all routes have been streamed to the client
	Subscribe bool `protobuf:"varint,1,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
	// option to request only changes that happen after the call
	WithoutReplay bool `protobuf:"varint,2,opt,name=withoutReplay,proto3" json:"withoutReplay,omitempty"`
}

func (m *ListUesRequest) Reset()         { *m = ListUesRequest{} }
func (m *ListUesRequest) String() string { return proto.CompactTextString(m) }
func (*ListUesRequest) ProtoMessage()    {}
func (*ListUesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{4}
}
func (m *ListUesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListUesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListUesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListUesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUesRequest.Merge(m, src)
}
func (m *ListUesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListUesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUesRequest proto.InternalMessageInfo

func (m *ListUesRequest) GetSubscribe() bool {
	if m != nil {
		return m.Subscribe
	}
	return false
}

func (m *ListUesRequest) GetWithoutReplay() bool {
	if m != nil {
		return m.WithoutReplay
	}
	return false
}

type ListUesResponse struct {
	// Ue is the UserEquipment change on which the event occurred
	Ue *types.Ue `protobuf:"bytes,1,opt,name=ue,proto3" json:"ue,omitempty"`
	// type is a qualification of the type of change being made
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=ran.trafficsim.Type" json:"type,omitempty"`
}

func (m *ListUesResponse) Reset()         { *m = ListUesResponse{} }
func (m *ListUesResponse) String() string { return proto.CompactTextString(m) }
func (*ListUesResponse) ProtoMessage()    {}
func (*ListUesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{5}
}
func (m *ListUesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListUesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListUesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListUesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUesResponse.Merge(m, src)
}
func (m *ListUesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListUesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUesResponse proto.InternalMessageInfo

func (m *ListUesResponse) GetUe() *types.Ue {
	if m != nil {
		return m.Ue
	}
	return nil
}

func (m *ListUesResponse) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_NONE
}

func init() {
	proto.RegisterEnum("ran.trafficsim.Type", Type_name, Type_value)
	proto.RegisterType((*MapLayoutRequest)(nil), "ran.trafficsim.MapLayoutRequest")
	proto.RegisterType((*ListTowersRequest)(nil), "ran.trafficsim.ListTowersRequest")
	proto.RegisterType((*ListRoutesRequest)(nil), "ran.trafficsim.ListRoutesRequest")
	proto.RegisterType((*ListRoutesResponse)(nil), "ran.trafficsim.ListRoutesResponse")
	proto.RegisterType((*ListUesRequest)(nil), "ran.trafficsim.ListUesRequest")
	proto.RegisterType((*ListUesResponse)(nil), "ran.trafficsim.ListUesResponse")
}

func init() { proto.RegisterFile("api/trafficsim/trafficsim.proto", fileDescriptor_47869854f8356ea4) }

var fileDescriptor_47869854f8356ea4 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0xad, 0x43, 0x47, 0xb7, 0x6f, 0x50, 0x82, 0xe1, 0x50, 0x15, 0x94, 0x8d, 0x88, 0x43, 0x85,
	0x44, 0x5a, 0x8a, 0xf8, 0x01, 0x9b, 0x12, 0xb8, 0x74, 0x2d, 0x44, 0x29, 0x9c, 0xdd, 0xe2, 0x75,
	0x16, 0x34, 0xf6, 0x62, 0x5b, 0x55, 0xff, 0x05, 0x3f, 0x8b, 0xe3, 0xb8, 0x71, 0x44, 0xed, 0x1f,
	0x41, 0xb1, 0xd7, 0x65, 0x4d, 0x03, 0x12, 0x12, 0x97, 0x28, 0x7e, 0x7a, 0xdf, 0xfb, 0xfc, 0xf4,
	0x9e, 0xe1, 0x88, 0x08, 0xd6, 0x55, 0x19, 0x39, 0x3f, 0x67, 0x53, 0xc9, 0xe6, 0xb7, 0x7e, 0x03,
	0x91, 0x71, 0xc5, 0x71, 0x33, 0x23, 0x69, 0x50, 0xa0, 0xed, 0xd3, 0x19, 0x53, 0x17, 0x7a, 0x12,
	0x4c, 0xf9, 0xbc, 0x3b, 0x12, 0x34, 0x1d, 0x52, 0xb5, 0xe0, 0xd9, 0x17, 0x96, 0xce, 0xde, 0x72,
	0x9d, 0x7e, 0x26, 0x8a, 0xf1, 0xb4, 0x3b, 0x9b, 0x13, 0xf1, 0x32, 0x23, 0x69, 0xd7, 0xa8, 0x2f,
	0x05, 0x95, 0xf6, 0x6b, 0x35, 0x7d, 0x0c, 0xee, 0x19, 0x11, 0x03, 0xb2, 0xe4, 0x5a, 0xc5, 0xf4,
	0x52, 0x53, 0xa9, 0xfc, 0x47, 0xf0, 0x70, 0xc0, 0xa4, 0x4a, 0xf8, 0x82, 0x66, 0x72, 0x03, 0x7e,
	0xb2, 0x60, 0xcc, 0xb5, 0xa2, 0x1b, 0x10, 0x3f, 0x85, 0x03, 0xa9, 0x27, 0x72, 0x9a, 0xb1, 0x09,
	0x6d, 0xa1, 0x63, 0xd4, 0xd9, 0x8f, 0x0b, 0x00, 0x3f, 0x87, 0xfb, 0x0b, 0xa6, 0x2e, 0x8c, 0xb2,
	0xf8, 0x4a, 0x96, 0x2d, 0xc7, 0x30, 0xb6, 0x41, 0xff, 0x12, 0xf0, 0x6d, 0x61, 0x29, 0x78, 0x2a,
	0x29, 0x7e, 0x05, 0x7b, 0x59, 0x8e, 0x18, 0xd5, 0xc3, 0xfe, 0x93, 0x60, 0xdb, 0x7b, 0x60, 0x3d,
	0x98, 0xa1, 0xd8, 0x32, 0x71, 0x07, 0xea, 0x39, 0x6a, 0xb6, 0x34, 0xfb, 0x8f, 0xcb, 0x13, 0xc9,
	0x52, 0xd0, 0xd8, 0x30, 0xfc, 0x04, 0x9a, 0xf9, 0xca, 0xf1, 0xff, 0x35, 0x42, 0xe1, 0xc1, 0x8d,
	0xea, 0xb5, 0x8b, 0x0e, 0x38, 0x7a, 0x63, 0xa1, 0x55, 0x6d, 0x61, 0x4c, 0x63, 0x47, 0xff, 0xc3,
	0xe5, 0x5f, 0xbc, 0x81, 0x7a, 0x7e, 0xc2, 0xfb, 0x50, 0x1f, 0x8e, 0x86, 0x91, 0x5b, 0xc3, 0x07,
	0xb0, 0x77, 0x12, 0x86, 0x51, 0xe8, 0x22, 0x7c, 0x08, 0x8d, 0xf1, 0xfb, 0xf0, 0x24, 0x89, 0x42,
	0xd7, 0xc9, 0x0f, 0x71, 0x74, 0x36, 0xfa, 0x18, 0x85, 0xee, 0x9d, 0xfe, 0x0f, 0x07, 0x1a, 0x89,
	0x15, 0xc4, 0x1f, 0xe0, 0xde, 0x3b, 0xaa, 0x6e, 0x72, 0xc7, 0xc7, 0xe5, 0x75, 0xe5, 0x4a, 0xb4,
	0x8f, 0xaa, 0x2f, 0x5f, 0x48, 0x0c, 0x01, 0x8a, 0xce, 0xe0, 0x67, 0x65, 0xfa, 0x4e, 0x9f, 0xda,
	0x7f, 0x48, 0xd4, 0x90, 0x7a, 0x08, 0x8f, 0xad, 0x9e, 0x6d, 0x45, 0xb5, 0xde, 0x56, 0x15, 0xdb,
	0xfe, 0xdf, 0x28, 0x36, 0x8e, 0x1e, 0xc2, 0x03, 0x68, 0x5c, 0x67, 0x84, 0xbd, 0xaa, 0x81, 0xa2,
	0x12, 0xbb, 0x96, 0x4b, 0xe1, 0xf6, 0xd0, 0x69, 0xeb, 0xfb, 0xca, 0x43, 0x57, 0x2b, 0x0f, 0xfd,
	0x5a, 0x79, 0xe8, 0xdb, 0xda, 0xab, 0x5d, 0xad, 0xbd, 0xda, 0xcf, 0xb5, 0x57, 0x9b, 0xdc, 0x35,
	0xaf, 0xeb, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xb9, 0x04, 0x5f, 0xd4, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrafficClient is the client API for Traffic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrafficClient interface {
	GetMapLayout(ctx context.Context, in *MapLayoutRequest, opts ...grpc.CallOption) (*types.MapLayout, error)
	ListTowers(ctx context.Context, in *ListTowersRequest, opts ...grpc.CallOption) (Traffic_ListTowersClient, error)
	ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (Traffic_ListRoutesClient, error)
	ListUes(ctx context.Context, in *ListUesRequest, opts ...grpc.CallOption) (Traffic_ListUesClient, error)
}

type trafficClient struct {
	cc *grpc.ClientConn
}

func NewTrafficClient(cc *grpc.ClientConn) TrafficClient {
	return &trafficClient{cc}
}

func (c *trafficClient) GetMapLayout(ctx context.Context, in *MapLayoutRequest, opts ...grpc.CallOption) (*types.MapLayout, error) {
	out := new(types.MapLayout)
	err := c.cc.Invoke(ctx, "/ran.trafficsim.Traffic/GetMapLayout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trafficClient) ListTowers(ctx context.Context, in *ListTowersRequest, opts ...grpc.CallOption) (Traffic_ListTowersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Traffic_serviceDesc.Streams[0], "/ran.trafficsim.Traffic/ListTowers", opts...)
	if err != nil {
		return nil, err
	}
	x := &trafficListTowersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Traffic_ListTowersClient interface {
	Recv() (*types.Tower, error)
	grpc.ClientStream
}

type trafficListTowersClient struct {
	grpc.ClientStream
}

func (x *trafficListTowersClient) Recv() (*types.Tower, error) {
	m := new(types.Tower)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trafficClient) ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (Traffic_ListRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Traffic_serviceDesc.Streams[1], "/ran.trafficsim.Traffic/ListRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &trafficListRoutesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Traffic_ListRoutesClient interface {
	Recv() (*ListRoutesResponse, error)
	grpc.ClientStream
}

type trafficListRoutesClient struct {
	grpc.ClientStream
}

func (x *trafficListRoutesClient) Recv() (*ListRoutesResponse, error) {
	m := new(ListRoutesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trafficClient) ListUes(ctx context.Context, in *ListUesRequest, opts ...grpc.CallOption) (Traffic_ListUesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Traffic_serviceDesc.Streams[2], "/ran.trafficsim.Traffic/ListUes", opts...)
	if err != nil {
		return nil, err
	}
	x := &trafficListUesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Traffic_ListUesClient interface {
	Recv() (*ListUesResponse, error)
	grpc.ClientStream
}

type trafficListUesClient struct {
	grpc.ClientStream
}

func (x *trafficListUesClient) Recv() (*ListUesResponse, error) {
	m := new(ListUesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrafficServer is the server API for Traffic service.
type TrafficServer interface {
	GetMapLayout(context.Context, *MapLayoutRequest) (*types.MapLayout, error)
	ListTowers(*ListTowersRequest, Traffic_ListTowersServer) error
	ListRoutes(*ListRoutesRequest, Traffic_ListRoutesServer) error
	ListUes(*ListUesRequest, Traffic_ListUesServer) error
}

// UnimplementedTrafficServer can be embedded to have forward compatible implementations.
type UnimplementedTrafficServer struct {
}

func (*UnimplementedTrafficServer) GetMapLayout(ctx context.Context, req *MapLayoutRequest) (*types.MapLayout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMapLayout not implemented")
}
func (*UnimplementedTrafficServer) ListTowers(req *ListTowersRequest, srv Traffic_ListTowersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTowers not implemented")
}
func (*UnimplementedTrafficServer) ListRoutes(req *ListRoutesRequest, srv Traffic_ListRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListRoutes not implemented")
}
func (*UnimplementedTrafficServer) ListUes(req *ListUesRequest, srv Traffic_ListUesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUes not implemented")
}

func RegisterTrafficServer(s *grpc.Server, srv TrafficServer) {
	s.RegisterService(&_Traffic_serviceDesc, srv)
}

func _Traffic_GetMapLayout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapLayoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficServer).GetMapLayout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ran.trafficsim.Traffic/GetMapLayout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficServer).GetMapLayout(ctx, req.(*MapLayoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Traffic_ListTowers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTowersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrafficServer).ListTowers(m, &trafficListTowersServer{stream})
}

type Traffic_ListTowersServer interface {
	Send(*types.Tower) error
	grpc.ServerStream
}

type trafficListTowersServer struct {
	grpc.ServerStream
}

func (x *trafficListTowersServer) Send(m *types.Tower) error {
	return x.ServerStream.SendMsg(m)
}

func _Traffic_ListRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRoutesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrafficServer).ListRoutes(m, &trafficListRoutesServer{stream})
}

type Traffic_ListRoutesServer interface {
	Send(*ListRoutesResponse) error
	grpc.ServerStream
}

type trafficListRoutesServer struct {
	grpc.ServerStream
}

func (x *trafficListRoutesServer) Send(m *ListRoutesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Traffic_ListUes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrafficServer).ListUes(m, &trafficListUesServer{stream})
}

type Traffic_ListUesServer interface {
	Send(*ListUesResponse) error
	grpc.ServerStream
}

type trafficListUesServer struct {
	grpc.ServerStream
}

func (x *trafficListUesServer) Send(m *ListUesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Traffic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ran.trafficsim.Traffic",
	HandlerType: (*TrafficServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMapLayout",
			Handler:    _Traffic_GetMapLayout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTowers",
			Handler:       _Traffic_ListTowers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListRoutes",
			Handler:       _Traffic_ListRoutes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListUes",
			Handler:       _Traffic_ListUes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/trafficsim/trafficsim.proto",
}

func (m *MapLayoutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapLayoutRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MapLayoutRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListTowersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListTowersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListTowersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListRoutesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListRoutesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListRoutesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithoutReplay {
		i--
		if m.WithoutReplay {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Subscribe {
		i--
		if m.Subscribe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListRoutesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListRoutesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListRoutesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintTrafficsim(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Route != nil {
		{
			size, err := m.Route.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTrafficsim(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListUesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListUesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListUesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithoutReplay {
		i--
		if m.WithoutReplay {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Subscribe {
		i--
		if m.Subscribe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListUesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListUesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListUesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintTrafficsim(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Ue != nil {
		{
			size, err := m.Ue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTrafficsim(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTrafficsim(dAtA []byte, offset int, v uint64) int {
	offset -= sovTrafficsim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MapLayoutRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListTowersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListRoutesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Subscribe {
		n += 2
	}
	if m.WithoutReplay {
		n += 2
	}
	return n
}

func (m *ListRoutesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Route != nil {
		l = m.Route.Size()
		n += 1 + l + sovTrafficsim(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTrafficsim(uint64(m.Type))
	}
	return n
}

func (m *ListUesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Subscribe {
		n += 2
	}
	if m.WithoutReplay {
		n += 2
	}
	return n
}

func (m *ListUesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ue != nil {
		l = m.Ue.Size()
		n += 1 + l + sovTrafficsim(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTrafficsim(uint64(m.Type))
	}
	return n
}

func sovTrafficsim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTrafficsim(x uint64) (n int) {
	return sovTrafficsim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MapLayoutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MapLayoutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapLayoutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListTowersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListTowersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListTowersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListRoutesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListRoutesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRoutesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Subscribe = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithoutReplay", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithoutReplay = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListRoutesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListRoutesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRoutesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrafficsim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Route == nil {
				m.Route = &types.Route{}
			}
			if err := m.Route.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListUesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListUesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListUesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Subscribe = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithoutReplay", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithoutReplay = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListUesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListUesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListUesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrafficsim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ue == nil {
				m.Ue = &types.Ue{}
			}
			if err := m.Ue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTrafficsim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTrafficsim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTrafficsim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTrafficsim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTrafficsim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrafficsim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTrafficsim = fmt.Errorf("proto: unexpected end of group")
)
