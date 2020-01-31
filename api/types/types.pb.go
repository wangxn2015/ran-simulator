// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/types/types.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
// A compilation error at this line likely means your cop	y of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Point struct {
	Lat float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float32 `protobuf:"fixed32,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_386b2e9db948721e, []int{0}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Point.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return m.Size()
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Point) GetLng() float32 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type Route struct {
	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Waypoints []*Point `protobuf:"bytes,2,rep,name=waypoints,proto3" json:"waypoints,omitempty"`
	Color     string   `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_386b2e9db948721e, []int{1}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Route.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return m.Size()
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Route) GetWaypoints() []*Point {
	if m != nil {
		return m.Waypoints
	}
	return nil
}

func (m *Route) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

type Ue struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Position *Point `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Rotation uint32 `protobuf:"varint,5,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Route    string `protobuf:"bytes,6,opt,name=route,proto3" json:"route,omitempty"`
	Tower    string `protobuf:"bytes,7,opt,name=tower,proto3" json:"tower,omitempty"`
	Tower2   string `protobuf:"bytes,8,opt,name=tower2,proto3" json:"tower2,omitempty"`
	Tower3   string `protobuf:"bytes,9,opt,name=tower3,proto3" json:"tower3,omitempty"`
}

func (m *Ue) Reset()         { *m = Ue{} }
func (m *Ue) String() string { return proto.CompactTextString(m) }
func (*Ue) ProtoMessage()    {}
func (*Ue) Descriptor() ([]byte, []int) {
	return fileDescriptor_386b2e9db948721e, []int{2}
}
func (m *Ue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ue.Merge(m, src)
}
func (m *Ue) XXX_Size() int {
	return m.Size()
}
func (m *Ue) XXX_DiscardUnknown() {
	xxx_messageInfo_Ue.DiscardUnknown(m)
}

var xxx_messageInfo_Ue proto.InternalMessageInfo

func (m *Ue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Ue) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Ue) GetRotation() uint32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

func (m *Ue) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *Ue) GetTower() string {
	if m != nil {
		return m.Tower
	}
	return ""
}

func (m *Ue) GetTower2() string {
	if m != nil {
		return m.Tower2
	}
	return ""
}

func (m *Ue) GetTower3() string {
	if m != nil {
		return m.Tower3
	}
	return ""
}

type Tower struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location *Point `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Color    string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (m *Tower) Reset()         { *m = Tower{} }
func (m *Tower) String() string { return proto.CompactTextString(m) }
func (*Tower) ProtoMessage()    {}
func (*Tower) Descriptor() ([]byte, []int) {
	return fileDescriptor_386b2e9db948721e, []int{3}
}
func (m *Tower) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tower) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tower.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tower) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tower.Merge(m, src)
}
func (m *Tower) XXX_Size() int {
	return m.Size()
}
func (m *Tower) XXX_DiscardUnknown() {
	xxx_messageInfo_Tower.DiscardUnknown(m)
}

var xxx_messageInfo_Tower proto.InternalMessageInfo

func (m *Tower) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tower) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Tower) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

type MapLayout struct {
	Center     *Point  `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Zoom       float32 `protobuf:"fixed32,2,opt,name=zoom,proto3" json:"zoom,omitempty"`
	Fade       bool    `protobuf:"varint,3,opt,name=fade,proto3" json:"fade,omitempty"`
	ShowRoutes bool    `protobuf:"varint,4,opt,name=showRoutes,proto3" json:"showRoutes,omitempty"`
}

func (m *MapLayout) Reset()         { *m = MapLayout{} }
func (m *MapLayout) String() string { return proto.CompactTextString(m) }
func (*MapLayout) ProtoMessage()    {}
func (*MapLayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_386b2e9db948721e, []int{4}
}
func (m *MapLayout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MapLayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MapLayout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MapLayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapLayout.Merge(m, src)
}
func (m *MapLayout) XXX_Size() int {
	return m.Size()
}
func (m *MapLayout) XXX_DiscardUnknown() {
	xxx_messageInfo_MapLayout.DiscardUnknown(m)
}

var xxx_messageInfo_MapLayout proto.InternalMessageInfo

func (m *MapLayout) GetCenter() *Point {
	if m != nil {
		return m.Center
	}
	return nil
}

func (m *MapLayout) GetZoom() float32 {
	if m != nil {
		return m.Zoom
	}
	return 0
}

func (m *MapLayout) GetFade() bool {
	if m != nil {
		return m.Fade
	}
	return false
}

func (m *MapLayout) GetShowRoutes() bool {
	if m != nil {
		return m.ShowRoutes
	}
	return false
}

func init() {
	proto.RegisterType((*Point)(nil), "ran.trafficsim.types.Point")
	proto.RegisterType((*Route)(nil), "ran.trafficsim.types.Route")
	proto.RegisterType((*Ue)(nil), "ran.trafficsim.types.Ue")
	proto.RegisterType((*Tower)(nil), "ran.trafficsim.types.Tower")
	proto.RegisterType((*MapLayout)(nil), "ran.trafficsim.types.MapLayout")
}

func init() { proto.RegisterFile("api/types/types.proto", fileDescriptor_386b2e9db948721e) }

var fileDescriptor_386b2e9db948721e = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x6a, 0xe3, 0x40,
	0x14, 0xf4, 0xca, 0x96, 0x4e, 0x7a, 0xc7, 0xc1, 0xb1, 0xf8, 0x8e, 0xe5, 0x0e, 0x84, 0x50, 0x65,
	0x38, 0xd0, 0x81, 0x55, 0x98, 0xb4, 0xa9, 0x13, 0x08, 0x4b, 0xf2, 0x01, 0x1b, 0x45, 0x76, 0x14,
	0x64, 0x3d, 0xb1, 0x5a, 0x63, 0x9c, 0x2f, 0x48, 0x99, 0xcf, 0x4a, 0xe9, 0x32, 0x4d, 0x20, 0xd8,
	0x3f, 0x12, 0xf6, 0x49, 0x71, 0x5c, 0x38, 0x38, 0x8d, 0x98, 0x19, 0xcd, 0xbe, 0x99, 0x95, 0x1e,
	0xfc, 0x52, 0x75, 0xf1, 0xdf, 0xac, 0xea, 0xbc, 0x69, 0x9f, 0x49, 0xad, 0xd1, 0x20, 0x1f, 0x6a,
	0x55, 0x25, 0x46, 0xab, 0xe9, 0xb4, 0xc8, 0x9a, 0x62, 0x9e, 0xd0, 0xbb, 0xf8, 0x1f, 0xb8, 0x17,
	0x58, 0x54, 0x86, 0xff, 0x84, 0x7e, 0xa9, 0x8c, 0x60, 0x11, 0x1b, 0x39, 0xd2, 0x42, 0x52, 0xaa,
	0x99, 0x70, 0x3a, 0xa5, 0x9a, 0xc5, 0x25, 0xb8, 0x12, 0x17, 0x26, 0xe7, 0x1c, 0x06, 0x95, 0x9a,
	0xe7, 0xe4, 0x0e, 0x24, 0x61, 0x7e, 0x02, 0xc1, 0x52, 0xad, 0x6a, 0x3b, 0xac, 0x11, 0x4e, 0xd4,
	0x1f, 0x7d, 0x1f, 0xff, 0x4d, 0x0e, 0x65, 0x26, 0x14, 0x28, 0x3f, 0xdc, 0x7c, 0x08, 0x6e, 0x86,
	0x25, 0x6a, 0xd1, 0xa7, 0x79, 0x2d, 0x89, 0x5f, 0x18, 0x38, 0x57, 0x87, 0xb3, 0x38, 0x0c, 0xec,
	0x28, 0xea, 0x16, 0x48, 0xc2, 0x7c, 0x02, 0x7e, 0x8d, 0x4d, 0x61, 0x0a, 0xac, 0xc4, 0x20, 0x62,
	0xc7, 0xe2, 0x77, 0x66, 0xfe, 0x07, 0x7c, 0x8d, 0x46, 0xd1, 0x41, 0x37, 0x62, 0xa3, 0x1f, 0x72,
	0xc7, 0x6d, 0x33, 0x6d, 0x6f, 0x2c, 0xbc, 0xb6, 0x19, 0x11, 0xab, 0x1a, 0x5c, 0xe6, 0x5a, 0x7c,
	0x6b, 0x55, 0x22, 0xfc, 0x37, 0x78, 0x04, 0xc6, 0xc2, 0x27, 0xb9, 0x63, 0x3b, 0x3d, 0x15, 0xc1,
	0x9e, 0x9e, 0xc6, 0x77, 0xe0, 0x5e, 0xd2, 0xc1, 0x43, 0x37, 0x9c, 0x80, 0x5f, 0x62, 0xd6, 0x96,
	0x72, 0xbe, 0x70, 0x9b, 0x77, 0xf3, 0x27, 0xdf, 0xf2, 0x81, 0x41, 0x70, 0xae, 0xea, 0x33, 0xb5,
	0xc2, 0x85, 0xe1, 0x29, 0x78, 0x59, 0x5e, 0x99, 0x5c, 0x53, 0xe4, 0x91, 0xd1, 0x9d, 0xd5, 0xb6,
	0xbc, 0x47, 0x9c, 0x77, 0xfb, 0x40, 0xd8, 0x6a, 0x53, 0x75, 0x93, 0x53, 0x96, 0x2f, 0x09, 0xf3,
	0x10, 0xa0, 0xb9, 0xc5, 0x25, 0x2d, 0x4a, 0x43, 0x7f, 0xc2, 0x97, 0x7b, 0xca, 0xa9, 0x78, 0xda,
	0x84, 0x6c, 0xbd, 0x09, 0xd9, 0xeb, 0x26, 0x64, 0x8f, 0xdb, 0xb0, 0xb7, 0xde, 0x86, 0xbd, 0xe7,
	0x6d, 0xd8, 0xbb, 0xf6, 0x68, 0x51, 0xd3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xec, 0x11,
	0x8e, 0xc1, 0x02, 0x00, 0x00,
}

func (m *Point) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Point) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Point) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Lng != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Lng))))
		i--
		dAtA[i] = 0x15
	}
	if m.Lat != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Lat))))
		i--
		dAtA[i] = 0xd
	}
	return len(dAtA) - i, nil
}

func (m *Route) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Route) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Route) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Color) > 0 {
		i -= len(m.Color)
		copy(dAtA[i:], m.Color)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Color)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Waypoints) > 0 {
		for iNdEx := len(m.Waypoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Waypoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Ue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tower3) > 0 {
		i -= len(m.Tower3)
		copy(dAtA[i:], m.Tower3)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tower3)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Tower2) > 0 {
		i -= len(m.Tower2)
		copy(dAtA[i:], m.Tower2)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tower2)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Tower) > 0 {
		i -= len(m.Tower)
		copy(dAtA[i:], m.Tower)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tower)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Route) > 0 {
		i -= len(m.Route)
		copy(dAtA[i:], m.Route)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Route)))
		i--
		dAtA[i] = 0x32
	}
	if m.Rotation != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Rotation))
		i--
		dAtA[i] = 0x28
	}
	if m.Position != nil {
		{
			size, err := m.Position.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Tower) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tower) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tower) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Color) > 0 {
		i -= len(m.Color)
		copy(dAtA[i:], m.Color)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Color)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Location != nil {
		{
			size, err := m.Location.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MapLayout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapLayout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MapLayout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShowRoutes {
		i--
		if m.ShowRoutes {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Fade {
		i--
		if m.Fade {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Zoom != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Zoom))))
		i--
		dAtA[i] = 0x15
	}
	if m.Center != nil {
		{
			size, err := m.Center.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Point) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Lat != 0 {
		n += 5
	}
	if m.Lng != 0 {
		n += 5
	}
	return n
}

func (m *Route) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Waypoints) > 0 {
		for _, e := range m.Waypoints {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Color)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Ue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Position != nil {
		l = m.Position.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Rotation != 0 {
		n += 1 + sovTypes(uint64(m.Rotation))
	}
	l = len(m.Route)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Tower)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Tower2)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Tower3)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Tower) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Location != nil {
		l = m.Location.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Color)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *MapLayout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Center != nil {
		l = m.Center.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Zoom != 0 {
		n += 5
	}
	if m.Fade {
		n += 2
	}
	if m.ShowRoutes {
		n += 2
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Point) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Point: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Point: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lat", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Lat = float32(math.Float32frombits(v))
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lng", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Lng = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Route) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Route: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Route: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Waypoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Waypoints = append(m.Waypoints, &Point{})
			if err := m.Waypoints[len(m.Waypoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Color", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Color = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Ue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Ue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Position == nil {
				m.Position = &Point{}
			}
			if err := m.Position.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rotation", wireType)
			}
			m.Rotation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rotation |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Route = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tower = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tower2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tower2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tower3", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tower3 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Tower) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Tower: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tower: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Location == nil {
				m.Location = &Point{}
			}
			if err := m.Location.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Color", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Color = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *MapLayout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MapLayout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapLayout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Center", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Center == nil {
				m.Center = &Point{}
			}
			if err := m.Center.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zoom", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Zoom = float32(math.Float32frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fade", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.Fade = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShowRoutes", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.ShowRoutes = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
