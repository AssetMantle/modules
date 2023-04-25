// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/transactions/make/message.proto

package make

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
	base2 "github.com/AssetMantle/schema/go/lists/base"
	base1 "github.com/AssetMantle/schema/go/types/base"
	proto "github.com/gogo/protobuf/proto"
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

type Message struct {
	From                    string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID                  *base.IdentityID       `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ClassificationID        *base.ClassificationID `protobuf:"bytes,3,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
	TakerID                 *base.IdentityID       `protobuf:"bytes,4,opt,name=taker_i_d,json=takerID,proto3" json:"taker_i_d,omitempty"`
	MakerOwnableID          *base.AnyOwnableID     `protobuf:"bytes,5,opt,name=maker_ownable_i_d,json=makerOwnableID,proto3" json:"maker_ownable_i_d,omitempty"`
	TakerOwnableID          *base.AnyOwnableID     `protobuf:"bytes,6,opt,name=taker_ownable_i_d,json=takerOwnableID,proto3" json:"taker_ownable_i_d,omitempty"`
	ExpiresIn               *base1.Height          `protobuf:"bytes,7,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	MakerOwnableSplit       string                 `protobuf:"bytes,8,opt,name=maker_ownable_split,json=makerOwnableSplit,proto3" json:"maker_ownable_split,omitempty"`
	TakerOwnableSplit       string                 `protobuf:"bytes,9,opt,name=taker_ownable_split,json=takerOwnableSplit,proto3" json:"taker_ownable_split,omitempty"`
	ImmutableMetaProperties *base2.PropertyList    `protobuf:"bytes,10,opt,name=immutable_meta_properties,json=immutableMetaProperties,proto3" json:"immutable_meta_properties,omitempty"`
	ImmutableProperties     *base2.PropertyList    `protobuf:"bytes,11,opt,name=immutable_properties,json=immutableProperties,proto3" json:"immutable_properties,omitempty"`
	MutableMetaProperties   *base2.PropertyList    `protobuf:"bytes,12,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3" json:"mutable_meta_properties,omitempty"`
	MutableProperties       *base2.PropertyList    `protobuf:"bytes,13,opt,name=mutable_properties,json=mutableProperties,proto3" json:"mutable_properties,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d41e3a5e005d9495, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetFromID() *base.IdentityID {
	if m != nil {
		return m.FromID
	}
	return nil
}

func (m *Message) GetClassificationID() *base.ClassificationID {
	if m != nil {
		return m.ClassificationID
	}
	return nil
}

func (m *Message) GetTakerID() *base.IdentityID {
	if m != nil {
		return m.TakerID
	}
	return nil
}

func (m *Message) GetMakerOwnableID() *base.AnyOwnableID {
	if m != nil {
		return m.MakerOwnableID
	}
	return nil
}

func (m *Message) GetTakerOwnableID() *base.AnyOwnableID {
	if m != nil {
		return m.TakerOwnableID
	}
	return nil
}

func (m *Message) GetExpiresIn() *base1.Height {
	if m != nil {
		return m.ExpiresIn
	}
	return nil
}

func (m *Message) GetMakerOwnableSplit() string {
	if m != nil {
		return m.MakerOwnableSplit
	}
	return ""
}

func (m *Message) GetTakerOwnableSplit() string {
	if m != nil {
		return m.TakerOwnableSplit
	}
	return ""
}

func (m *Message) GetImmutableMetaProperties() *base2.PropertyList {
	if m != nil {
		return m.ImmutableMetaProperties
	}
	return nil
}

func (m *Message) GetImmutableProperties() *base2.PropertyList {
	if m != nil {
		return m.ImmutableProperties
	}
	return nil
}

func (m *Message) GetMutableMetaProperties() *base2.PropertyList {
	if m != nil {
		return m.MutableMetaProperties
	}
	return nil
}

func (m *Message) GetMutableProperties() *base2.PropertyList {
	if m != nil {
		return m.MutableProperties
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "assetmantle.modules.orders.transactions.make.Message")
}

func init() {
	proto.RegisterFile("x/orders/internal/transactions/make/message.proto", fileDescriptor_d41e3a5e005d9495)
}

var fileDescriptor_d41e3a5e005d9495 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x97, 0xb0, 0xad, 0xab, 0x37, 0x10, 0xcb, 0x40, 0x0b, 0x93, 0x88, 0x26, 0x24, 0xc4,
	0x10, 0xc3, 0x19, 0x70, 0x41, 0xb9, 0xb5, 0x8b, 0x04, 0x91, 0x88, 0x56, 0x85, 0x9d, 0x46, 0x45,
	0xe4, 0x26, 0x5e, 0x6b, 0x2d, 0x4e, 0xa2, 0xd8, 0x13, 0xeb, 0x5b, 0xf0, 0x0c, 0x1c, 0x79, 0x12,
	0xc4, 0x69, 0x47, 0x4e, 0x08, 0x75, 0x37, 0x8e, 0x3c, 0x01, 0xb2, 0x63, 0xa5, 0xe9, 0x5a, 0x26,
	0x72, 0x6a, 0x9b, 0xff, 0xf7, 0xfd, 0xfc, 0x7d, 0x75, 0xf4, 0x07, 0x2f, 0x2e, 0xec, 0xac, 0x88,
	0x71, 0xc1, 0x6c, 0x92, 0x72, 0x5c, 0xa4, 0x28, 0xb1, 0x79, 0x81, 0x52, 0x86, 0x22, 0x4e, 0xb2,
	0x94, 0xd9, 0x14, 0x9d, 0x61, 0x9b, 0x62, 0xc6, 0xd0, 0x10, 0xc3, 0xbc, 0xc8, 0x78, 0x66, 0xec,
	0x23, 0xc6, 0x30, 0xa7, 0x28, 0xe5, 0x09, 0x86, 0x34, 0x8b, 0xcf, 0x13, 0xcc, 0x60, 0x09, 0x81,
	0x75, 0x2f, 0x14, 0xde, 0x9d, 0x5d, 0x12, 0x33, 0x7b, 0x80, 0x18, 0xb6, 0xa3, 0x04, 0x31, 0x46,
	0x4e, 0x49, 0x84, 0xc4, 0x38, 0x24, 0x71, 0xc9, 0xdb, 0xd9, 0xa9, 0x14, 0x24, 0xc6, 0x29, 0x27,
	0x7c, 0x3c, 0x9d, 0x3d, 0xac, 0x66, 0x28, 0x1d, 0x87, 0xd9, 0xa7, 0x14, 0x0d, 0x12, 0x3c, 0x1d,
	0x5b, 0x09, 0x61, 0x5c, 0x09, 0xf2, 0x22, 0xcb, 0x71, 0xc1, 0xc7, 0xa1, 0x78, 0xa6, 0xe6, 0xdb,
	0x7c, 0x9c, 0x63, 0x35, 0x1f, 0x61, 0x32, 0x1c, 0xa9, 0xc1, 0xa3, 0x9f, 0x2d, 0xd0, 0xf2, 0xcb,
	0x56, 0x86, 0x01, 0x96, 0x4f, 0x8b, 0x8c, 0x9a, 0xda, 0xae, 0xb6, 0xd7, 0x0e, 0xe4, 0x77, 0xa3,
	0x03, 0xd6, 0xc4, 0x67, 0x48, 0xc2, 0xd8, 0xd4, 0x77, 0xb5, 0xbd, 0xf5, 0x97, 0x4f, 0x60, 0xbd,
	0x36, 0x8b, 0x46, 0x98, 0x22, 0x48, 0x62, 0x06, 0x05, 0x1c, 0x7a, 0x2a, 0xb9, 0xe7, 0x06, 0xab,
	0xc2, 0xe8, 0xb9, 0xc6, 0x07, 0x60, 0x5c, 0x6f, 0x1c, 0xc6, 0xe6, 0x2d, 0x09, 0x7b, 0x7e, 0x23,
	0xec, 0x70, 0xc6, 0xe6, 0xb9, 0xc1, 0xdd, 0xe8, 0xda, 0x13, 0xe3, 0x10, 0xb4, 0x39, 0x3a, 0xc3,
	0x85, 0x64, 0x2e, 0x37, 0x0b, 0xd8, 0x92, 0x4e, 0xcf, 0x35, 0x8e, 0xc1, 0x26, 0x95, 0x90, 0xea,
	0x7f, 0x0d, 0x63, 0x73, 0x45, 0xc2, 0x9e, 0xde, 0x08, 0xeb, 0xa4, 0xe3, 0xa3, 0xd2, 0xe2, 0xb9,
	0xc1, 0x1d, 0xc9, 0xa8, 0x7e, 0x0b, 0x2a, 0x9f, 0xa3, 0xae, 0x36, 0xa6, 0xf2, 0x59, 0xaa, 0x0b,
	0x00, 0xbe, 0xc8, 0x49, 0x81, 0x59, 0x48, 0x52, 0xb3, 0x25, 0x71, 0x8f, 0x17, 0xe1, 0xe4, 0x8d,
	0x97, 0xc0, 0xb7, 0xf2, 0xc6, 0x83, 0xb6, 0x32, 0x7a, 0xa9, 0x01, 0xc1, 0xd6, 0x6c, 0x63, 0x96,
	0x27, 0x84, 0x9b, 0x6b, 0xf2, 0xe6, 0x37, 0xeb, 0x45, 0xde, 0x8b, 0x81, 0xd0, 0xf3, 0x05, 0xfa,
	0x76, 0xa9, 0xe7, 0x73, 0xfa, 0x21, 0x78, 0x40, 0x28, 0x3d, 0xe7, 0x52, 0x4b, 0x31, 0x47, 0xa1,
	0x7a, 0x2b, 0x09, 0x66, 0x26, 0x90, 0xa1, 0x9f, 0x2d, 0x0a, 0x2d, 0x5f, 0xe3, 0x32, 0x74, 0x4f,
	0xbd, 0xc6, 0xef, 0x08, 0xe3, 0xc1, 0x76, 0x45, 0xf3, 0x31, 0x47, 0xbd, 0x8a, 0x65, 0x7c, 0x04,
	0xf7, 0xa6, 0x07, 0xd5, 0xce, 0x58, 0x6f, 0x7e, 0xc6, 0x56, 0x05, 0xaa, 0xf1, 0x23, 0xb0, 0xfd,
	0xaf, 0x1a, 0x1b, 0xcd, 0x8f, 0xb8, 0xbf, 0xb8, 0xc4, 0x09, 0x30, 0x16, 0x54, 0xb8, 0xdd, 0x9c,
	0xbf, 0x39, 0x57, 0xa0, 0xfb, 0x47, 0xff, 0x36, 0xb1, 0xb4, 0xcb, 0x89, 0xa5, 0xfd, 0x9a, 0x58,
	0xda, 0xe7, 0x2b, 0x6b, 0xe9, 0xf2, 0xca, 0x5a, 0xfa, 0x71, 0x65, 0x2d, 0x81, 0x83, 0x28, 0xa3,
	0xb0, 0xc9, 0x0e, 0xeb, 0x6e, 0xa8, 0x55, 0xd1, 0x13, 0xbb, 0xa3, 0xa7, 0x9d, 0x74, 0x87, 0x84,
	0x8f, 0xce, 0x07, 0x30, 0xca, 0xa8, 0xdd, 0x11, 0x20, 0x5f, 0x82, 0x6c, 0x05, 0xb2, 0xff, 0x63,
	0xa7, 0x7e, 0xd1, 0x57, 0x3a, 0xfe, 0xd1, 0xb1, 0xff, 0x55, 0xdf, 0xef, 0xd4, 0xa2, 0xf8, 0x2a,
	0xca, 0x51, 0x19, 0xe5, 0xb8, 0x1e, 0xc5, 0x47, 0x67, 0xf8, 0xfb, 0x8c, 0xbc, 0xaf, 0xe4, 0xfd,
	0x52, 0xde, 0xaf, 0xcb, 0xfb, 0x42, 0x3e, 0xd1, 0x5f, 0x37, 0x91, 0xf7, 0xdf, 0xf4, 0xba, 0xe2,
	0x8e, 0x62, 0xc4, 0xd1, 0x6f, 0xfd, 0xa0, 0x66, 0x75, 0x1c, 0xe5, 0x75, 0x9c, 0xd2, 0xec, 0x38,
	0x75, 0xb7, 0xe3, 0x08, 0xfb, 0x60, 0x55, 0x2e, 0xd7, 0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0xba, 0xfb, 0xf8, 0x55, 0x06, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MutableProperties != nil {
		{
			size, err := m.MutableProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if m.MutableMetaProperties != nil {
		{
			size, err := m.MutableMetaProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if m.ImmutableProperties != nil {
		{
			size, err := m.ImmutableProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.ImmutableMetaProperties != nil {
		{
			size, err := m.ImmutableMetaProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.TakerOwnableSplit) > 0 {
		i -= len(m.TakerOwnableSplit)
		copy(dAtA[i:], m.TakerOwnableSplit)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.TakerOwnableSplit)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.MakerOwnableSplit) > 0 {
		i -= len(m.MakerOwnableSplit)
		copy(dAtA[i:], m.MakerOwnableSplit)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.MakerOwnableSplit)))
		i--
		dAtA[i] = 0x42
	}
	if m.ExpiresIn != nil {
		{
			size, err := m.ExpiresIn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.TakerOwnableID != nil {
		{
			size, err := m.TakerOwnableID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.MakerOwnableID != nil {
		{
			size, err := m.MakerOwnableID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.TakerID != nil {
		{
			size, err := m.TakerID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FromID != nil {
		{
			size, err := m.FromID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
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
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.FromID != nil {
		l = m.FromID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.TakerID != nil {
		l = m.TakerID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.MakerOwnableID != nil {
		l = m.MakerOwnableID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.TakerOwnableID != nil {
		l = m.TakerOwnableID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ExpiresIn != nil {
		l = m.ExpiresIn.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.MakerOwnableSplit)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.TakerOwnableSplit)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ImmutableMetaProperties != nil {
		l = m.ImmutableMetaProperties.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ImmutableProperties != nil {
		l = m.ImmutableProperties.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.MutableMetaProperties != nil {
		l = m.MutableMetaProperties.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.MutableProperties != nil {
		l = m.MutableProperties.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
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
			if m.FromID == nil {
				m.FromID = &base.IdentityID{}
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			if m.ClassificationID == nil {
				m.ClassificationID = &base.ClassificationID{}
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerID", wireType)
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
			if m.TakerID == nil {
				m.TakerID = &base.IdentityID{}
			}
			if err := m.TakerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableID", wireType)
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
			if m.MakerOwnableID == nil {
				m.MakerOwnableID = &base.AnyOwnableID{}
			}
			if err := m.MakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableID", wireType)
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
			if m.TakerOwnableID == nil {
				m.TakerOwnableID = &base.AnyOwnableID{}
			}
			if err := m.TakerOwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresIn", wireType)
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
			if m.ExpiresIn == nil {
				m.ExpiresIn = &base1.Height{}
			}
			if err := m.ExpiresIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerOwnableSplit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MakerOwnableSplit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableSplit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TakerOwnableSplit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableMetaProperties", wireType)
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
			if m.ImmutableMetaProperties == nil {
				m.ImmutableMetaProperties = &base2.PropertyList{}
			}
			if err := m.ImmutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableProperties", wireType)
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
			if m.ImmutableProperties == nil {
				m.ImmutableProperties = &base2.PropertyList{}
			}
			if err := m.ImmutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
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
			if m.MutableMetaProperties == nil {
				m.MutableMetaProperties = &base2.PropertyList{}
			}
			if err := m.MutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
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
			if m.MutableProperties == nil {
				m.MutableProperties = &base2.PropertyList{}
			}
			if err := m.MutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
	depth := 0
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
		case 1:
			iNdEx += 8
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
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
