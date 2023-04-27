// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/transactions/make/message.proto

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
	return fileDescriptor_59959bcb14043951, []int{0}
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
	proto.RegisterFile("orders/transactions/make/message.proto", fileDescriptor_59959bcb14043951)
}

var fileDescriptor_59959bcb14043951 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6f, 0xd3, 0x3c,
	0x1c, 0xc6, 0x97, 0xbc, 0xdb, 0xba, 0x7a, 0x7b, 0x11, 0xcb, 0x40, 0x0b, 0x93, 0x88, 0x26, 0x24,
	0x60, 0x88, 0xe1, 0x4c, 0x70, 0x81, 0xdc, 0xba, 0x45, 0x82, 0x48, 0x44, 0xab, 0xc2, 0x4e, 0xa3,
	0x22, 0x72, 0x13, 0xaf, 0xb5, 0x16, 0x27, 0x51, 0xec, 0x89, 0xf5, 0x5b, 0xf0, 0x19, 0x38, 0xf2,
	0x49, 0x10, 0xa7, 0x1d, 0x39, 0x21, 0xd4, 0x9e, 0xe0, 0x53, 0x20, 0x3b, 0x56, 0x9a, 0xae, 0xdd,
	0xa4, 0x9c, 0xda, 0xe6, 0xff, 0x3c, 0x3f, 0x3f, 0x4f, 0x1d, 0xfd, 0xc1, 0x93, 0xac, 0x88, 0x71,
	0xc1, 0x6c, 0x5e, 0xa0, 0x94, 0xa1, 0x88, 0x93, 0x2c, 0x65, 0x36, 0x45, 0xe7, 0xd8, 0xa6, 0x98,
	0x31, 0x34, 0xc0, 0x30, 0x2f, 0x32, 0x9e, 0x19, 0xfb, 0x88, 0x31, 0xcc, 0x29, 0x4a, 0x79, 0x82,
	0x21, 0xcd, 0xe2, 0x8b, 0x04, 0x33, 0x58, 0x7a, 0x61, 0xdd, 0x0b, 0x85, 0x77, 0x67, 0x97, 0xc4,
	0xcc, 0xee, 0x23, 0x86, 0xed, 0x28, 0x41, 0x8c, 0x91, 0x33, 0x12, 0x21, 0x31, 0x0e, 0x49, 0x5c,
	0xf2, 0x76, 0x76, 0x2a, 0x05, 0x89, 0x71, 0xca, 0x09, 0x1f, 0x4d, 0x67, 0x0f, 0xab, 0x19, 0x4a,
	0x47, 0x61, 0xf6, 0x39, 0x45, 0xfd, 0x04, 0x4f, 0xc7, 0x56, 0x42, 0x18, 0x57, 0x82, 0xbc, 0xc8,
	0x72, 0x5c, 0xf0, 0x51, 0x28, 0x9e, 0xa9, 0xf9, 0x36, 0x1f, 0xe5, 0x58, 0xcd, 0x87, 0x98, 0x0c,
	0x86, 0x6a, 0xf0, 0xe8, 0x57, 0x0b, 0xb4, 0xfc, 0xb2, 0x95, 0x61, 0x80, 0xe5, 0xb3, 0x22, 0xa3,
	0xa6, 0xb6, 0xab, 0xed, 0xb5, 0x03, 0xf9, 0xdd, 0xe8, 0x80, 0x35, 0xf1, 0x19, 0x92, 0x30, 0x36,
	0xf5, 0x5d, 0x6d, 0x6f, 0xfd, 0xe5, 0x53, 0x58, 0xaf, 0xcd, 0xa2, 0x21, 0xa6, 0x08, 0x92, 0x98,
	0x41, 0x01, 0x87, 0x9e, 0x4a, 0xee, 0xb9, 0xc1, 0xaa, 0x30, 0x7a, 0xae, 0xf1, 0x11, 0x18, 0xd7,
	0x1b, 0x87, 0xb1, 0xf9, 0x9f, 0x84, 0xbd, 0xb8, 0x15, 0x76, 0x34, 0x63, 0xf3, 0xdc, 0xe0, 0x6e,
	0x74, 0xed, 0x89, 0x71, 0x04, 0xda, 0x1c, 0x9d, 0xe3, 0x42, 0x32, 0x97, 0x9b, 0x05, 0x6c, 0x49,
	0xa7, 0xe7, 0x1a, 0x27, 0x60, 0x93, 0x4a, 0x48, 0xf5, 0xbf, 0x86, 0xb1, 0xb9, 0x22, 0x61, 0xcf,
	0x6e, 0x85, 0x75, 0xd2, 0xd1, 0x71, 0x69, 0xf1, 0xdc, 0xe0, 0x8e, 0x64, 0x54, 0xbf, 0x05, 0x95,
	0xcf, 0x51, 0x57, 0x1b, 0x53, 0xf9, 0x2c, 0xd5, 0x05, 0x00, 0x5f, 0xe6, 0xa4, 0xc0, 0x2c, 0x24,
	0xa9, 0xd9, 0x92, 0xb8, 0xc7, 0x8b, 0x70, 0xf2, 0xc6, 0x4b, 0xe0, 0x3b, 0x79, 0xe3, 0x41, 0x5b,
	0x19, 0xbd, 0xd4, 0x80, 0x60, 0x6b, 0xb6, 0x31, 0xcb, 0x13, 0xc2, 0xcd, 0x35, 0x79, 0xf3, 0x9b,
	0xf5, 0x22, 0x1f, 0xc4, 0x40, 0xe8, 0xf9, 0x02, 0x7d, 0xbb, 0xd4, 0xf3, 0x39, 0xfd, 0x00, 0x3c,
	0x20, 0x94, 0x5e, 0x70, 0xa9, 0xa5, 0x98, 0xa3, 0x50, 0xbd, 0x95, 0x04, 0x33, 0x13, 0xc8, 0xd0,
	0xcf, 0x17, 0x85, 0x96, 0xaf, 0x71, 0x19, 0xba, 0xab, 0x5e, 0xe3, 0xf7, 0x84, 0xf1, 0x60, 0xbb,
	0xa2, 0xf9, 0x98, 0xa3, 0x6e, 0xc5, 0x32, 0x3e, 0x81, 0x7b, 0xd3, 0x83, 0x6a, 0x67, 0xac, 0x37,
	0x3f, 0x63, 0xab, 0x02, 0xd5, 0xf8, 0x11, 0xd8, 0xbe, 0xa9, 0xc6, 0x46, 0xf3, 0x23, 0xee, 0x2f,
	0x2e, 0x71, 0x0a, 0x8c, 0x05, 0x15, 0xfe, 0x6f, 0xce, 0xdf, 0x9c, 0x2b, 0x70, 0xf8, 0x47, 0xff,
	0x3e, 0xb6, 0xb4, 0xab, 0xb1, 0xa5, 0xfd, 0x1e, 0x5b, 0xda, 0x97, 0x89, 0xb5, 0x74, 0x35, 0xb1,
	0x96, 0x7e, 0x4e, 0xac, 0x25, 0x70, 0x10, 0x65, 0x14, 0x36, 0xd9, 0x61, 0x87, 0x1b, 0x6a, 0x55,
	0x74, 0xc5, 0xee, 0xe8, 0x6a, 0xa7, 0x6f, 0x06, 0x84, 0x0f, 0x2f, 0xfa, 0x30, 0xca, 0xa8, 0xdd,
	0x11, 0x20, 0x5f, 0x82, 0x6c, 0x05, 0xb2, 0x2f, 0xed, 0x9b, 0x56, 0xe9, 0x57, 0x7d, 0xa5, 0xe3,
	0x1f, 0x9f, 0xf8, 0xdf, 0xf4, 0xfd, 0x4e, 0x2d, 0x81, 0xaf, 0x12, 0x1c, 0x97, 0x09, 0x4e, 0xea,
	0x09, 0x7c, 0x74, 0x8e, 0x7f, 0xcc, 0xc8, 0x7b, 0x4a, 0xde, 0x2b, 0xe5, 0xbd, 0xba, 0xbc, 0x27,
	0xe4, 0x63, 0xfd, 0x75, 0x13, 0x79, 0xef, 0x6d, 0xf7, 0x50, 0x5c, 0x4d, 0x8c, 0x38, 0xfa, 0xab,
	0x1f, 0xd4, 0xac, 0x8e, 0xa3, 0xbc, 0x8e, 0x53, 0x9a, 0x1d, 0xa7, 0xee, 0x76, 0x1c, 0x61, 0xef,
	0xaf, 0xca, 0x9d, 0xfa, 0xea, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xd7, 0xfb, 0xff, 0x41,
	0x06, 0x00, 0x00,
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