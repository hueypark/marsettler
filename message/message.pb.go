// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message/message.proto

package message

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Actor struct {
	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerId int32   `protobuf:"varint,2,opt,name=serverId,proto3" json:"serverId,omitempty"`
	Position *Vector `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Velocity *Vector `protobuf:"bytes,4,opt,name=velocity,proto3" json:"velocity,omitempty"`
}

func (m *Actor) Reset()         { *m = Actor{} }
func (m *Actor) String() string { return proto.CompactTextString(m) }
func (*Actor) ProtoMessage()    {}
func (*Actor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebceca9e8703e37f, []int{0}
}
func (m *Actor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Actor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Actor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Actor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Actor.Merge(m, src)
}
func (m *Actor) XXX_Size() int {
	return m.Size()
}
func (m *Actor) XXX_DiscardUnknown() {
	xxx_messageInfo_Actor.DiscardUnknown(m)
}

var xxx_messageInfo_Actor proto.InternalMessageInfo

func (m *Actor) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Actor) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *Actor) GetPosition() *Vector {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Actor) GetVelocity() *Vector {
	if m != nil {
		return m.Velocity
	}
	return nil
}

type ActorCreate struct {
	Pos *Vector `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (m *ActorCreate) Reset()         { *m = ActorCreate{} }
func (m *ActorCreate) String() string { return proto.CompactTextString(m) }
func (*ActorCreate) ProtoMessage()    {}
func (*ActorCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebceca9e8703e37f, []int{1}
}
func (m *ActorCreate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActorCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActorCreate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActorCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActorCreate.Merge(m, src)
}
func (m *ActorCreate) XXX_Size() int {
	return m.Size()
}
func (m *ActorCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ActorCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ActorCreate proto.InternalMessageInfo

func (m *ActorCreate) GetPos() *Vector {
	if m != nil {
		return m.Pos
	}
	return nil
}

type Actors struct {
	Actors []*Actor `protobuf:"bytes,1,rep,name=actors,proto3" json:"actors,omitempty"`
}

func (m *Actors) Reset()         { *m = Actors{} }
func (m *Actors) String() string { return proto.CompactTextString(m) }
func (*Actors) ProtoMessage()    {}
func (*Actors) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebceca9e8703e37f, []int{2}
}
func (m *Actors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Actors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Actors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Actors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Actors.Merge(m, src)
}
func (m *Actors) XXX_Size() int {
	return m.Size()
}
func (m *Actors) XXX_DiscardUnknown() {
	xxx_messageInfo_Actors.DiscardUnknown(m)
}

var xxx_messageInfo_Actors proto.InternalMessageInfo

func (m *Actors) GetActors() []*Actor {
	if m != nil {
		return m.Actors
	}
	return nil
}

type Vector struct {
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *Vector) Reset()         { *m = Vector{} }
func (m *Vector) String() string { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()    {}
func (*Vector) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebceca9e8703e37f, []int{3}
}
func (m *Vector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector.Merge(m, src)
}
func (m *Vector) XXX_Size() int {
	return m.Size()
}
func (m *Vector) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector.DiscardUnknown(m)
}

var xxx_messageInfo_Vector proto.InternalMessageInfo

func (m *Vector) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type World struct {
	Left   float64 `protobuf:"fixed64,1,opt,name=left,proto3" json:"left,omitempty"`
	Right  float64 `protobuf:"fixed64,2,opt,name=right,proto3" json:"right,omitempty"`
	Bottom float64 `protobuf:"fixed64,3,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Top    float64 `protobuf:"fixed64,4,opt,name=top,proto3" json:"top,omitempty"`
}

func (m *World) Reset()         { *m = World{} }
func (m *World) String() string { return proto.CompactTextString(m) }
func (*World) ProtoMessage()    {}
func (*World) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebceca9e8703e37f, []int{4}
}
func (m *World) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *World) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_World.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *World) XXX_Merge(src proto.Message) {
	xxx_messageInfo_World.Merge(m, src)
}
func (m *World) XXX_Size() int {
	return m.Size()
}
func (m *World) XXX_DiscardUnknown() {
	xxx_messageInfo_World.DiscardUnknown(m)
}

var xxx_messageInfo_World proto.InternalMessageInfo

func (m *World) GetLeft() float64 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *World) GetRight() float64 {
	if m != nil {
		return m.Right
	}
	return 0
}

func (m *World) GetBottom() float64 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *World) GetTop() float64 {
	if m != nil {
		return m.Top
	}
	return 0
}

func init() {
	proto.RegisterType((*Actor)(nil), "message.Actor")
	proto.RegisterType((*ActorCreate)(nil), "message.ActorCreate")
	proto.RegisterType((*Actors)(nil), "message.Actors")
	proto.RegisterType((*Vector)(nil), "message.Vector")
	proto.RegisterType((*World)(nil), "message.World")
}

func init() { proto.RegisterFile("message/message.proto", fileDescriptor_ebceca9e8703e37f) }

var fileDescriptor_ebceca9e8703e37f = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x3b, 0xdd, 0xee, 0x5a, 0xa6, 0x52, 0x65, 0x50, 0x09, 0x1e, 0xc2, 0xba, 0x88, 0x2c,
	0x08, 0xb5, 0xd4, 0x27, 0x50, 0x4f, 0x5e, 0x73, 0xd0, 0x83, 0xa7, 0xb6, 0x1b, 0x6b, 0xa0, 0x35,
	0x4b, 0x12, 0x4a, 0xfb, 0x10, 0x82, 0x8f, 0xe5, 0xb1, 0x47, 0x8f, 0xd2, 0x7d, 0x11, 0xd9, 0x34,
	0xdd, 0x93, 0x9e, 0xf2, 0xff, 0xcc, 0x37, 0x7f, 0x66, 0x12, 0x3c, 0x5d, 0x48, 0x6b, 0xc7, 0x33,
	0x79, 0x13, 0xce, 0x41, 0x69, 0xb4, 0xd3, 0x74, 0x10, 0x6c, 0xf6, 0x01, 0x18, 0xdf, 0x4d, 0x9d,
	0x36, 0xd4, 0xc7, 0xb6, 0x2a, 0x18, 0xa4, 0x90, 0x47, 0xa2, 0xad, 0x0a, 0x3a, 0xc7, 0xae, 0x95,
	0x66, 0x29, 0xcd, 0x63, 0xc1, 0xda, 0x29, 0xe4, 0xb1, 0x68, 0x3c, 0x5d, 0x63, 0xb7, 0xd4, 0x56,
	0x39, 0xa5, 0xdf, 0x59, 0x94, 0x42, 0xde, 0x1b, 0x1d, 0x0d, 0xf6, 0x17, 0x3c, 0xc9, 0x3a, 0x4e,
	0x34, 0x40, 0x0d, 0x2f, 0xe5, 0x5c, 0x4f, 0x95, 0x5b, 0xb3, 0xce, 0x3f, 0xf0, 0x1e, 0xc8, 0x86,
	0xd8, 0xf3, 0xe3, 0x3c, 0x18, 0x39, 0x76, 0x92, 0x2e, 0x30, 0x2a, 0xb5, 0xf5, 0x53, 0xfd, 0xd1,
	0x56, 0xd7, 0xb2, 0x21, 0x26, 0xbe, 0xc3, 0xd2, 0x15, 0x26, 0x63, 0xaf, 0x18, 0xa4, 0x51, 0xde,
	0x1b, 0xf5, 0x1b, 0xde, 0x03, 0x22, 0x54, 0xb3, 0x4b, 0x4c, 0x76, 0x01, 0x74, 0x88, 0xb0, 0xf2,
	0xe1, 0x20, 0x60, 0x55, 0xbb, 0xb5, 0x5f, 0x15, 0x04, 0xac, 0xb3, 0x17, 0x8c, 0x9f, 0xb5, 0x99,
	0x17, 0x44, 0xd8, 0x99, 0xcb, 0x57, 0x17, 0x38, 0xaf, 0xe9, 0x04, 0x63, 0xa3, 0x66, 0x6f, 0x2e,
	0xe0, 0x3b, 0x43, 0x67, 0x98, 0x4c, 0xb4, 0x73, 0x7a, 0xe1, 0x1f, 0x05, 0x44, 0x70, 0x74, 0x8c,
	0x91, 0xd3, 0xa5, 0x5f, 0x1e, 0x44, 0x2d, 0xef, 0xd9, 0xd7, 0x96, 0xc3, 0x66, 0xcb, 0xe1, 0x67,
	0xcb, 0xe1, 0xb3, 0xe2, 0xad, 0x4d, 0xc5, 0x5b, 0xdf, 0x15, 0x6f, 0x4d, 0x12, 0xff, 0x41, 0xb7,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x46, 0xf0, 0xfd, 0xb9, 0x01, 0x00, 0x00,
}

func (m *Actor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Actor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Actor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Velocity != nil {
		{
			size, err := m.Velocity.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Position != nil {
		{
			size, err := m.Position.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ServerId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.ServerId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ActorCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActorCreate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActorCreate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pos != nil {
		{
			size, err := m.Pos.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Actors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Actors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Actors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Actors) > 0 {
		for iNdEx := len(m.Actors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Vector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Y != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Y))))
		i--
		dAtA[i] = 0x11
	}
	if m.X != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.X))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *World) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *World) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *World) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Top != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Top))))
		i--
		dAtA[i] = 0x21
	}
	if m.Bottom != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Bottom))))
		i--
		dAtA[i] = 0x19
	}
	if m.Right != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Right))))
		i--
		dAtA[i] = 0x11
	}
	if m.Left != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Left))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Actor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMessage(uint64(m.Id))
	}
	if m.ServerId != 0 {
		n += 1 + sovMessage(uint64(m.ServerId))
	}
	if m.Position != nil {
		l = m.Position.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Velocity != nil {
		l = m.Velocity.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *ActorCreate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pos != nil {
		l = m.Pos.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *Actors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Actors) > 0 {
		for _, e := range m.Actors {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	return n
}

func (m *Vector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 9
	}
	if m.Y != 0 {
		n += 9
	}
	return n
}

func (m *World) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Left != 0 {
		n += 9
	}
	if m.Right != 0 {
		n += 9
	}
	if m.Bottom != 0 {
		n += 9
	}
	if m.Top != 0 {
		n += 9
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Actor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Actor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Actor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerId", wireType)
			}
			m.ServerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Position == nil {
				m.Position = &Vector{}
			}
			if err := m.Position.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Velocity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Velocity == nil {
				m.Velocity = &Vector{}
			}
			if err := m.Velocity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *ActorCreate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: ActorCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActorCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pos == nil {
				m.Pos = &Vector{}
			}
			if err := m.Pos.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *Actors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Actors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Actors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actors = append(m.Actors, &Actor{})
			if err := m.Actors[len(m.Actors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *Vector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Vector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.X = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Y = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *World) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: World: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: World: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Left = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Right = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bottom", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Bottom = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Top", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Top = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMessage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMessage
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)