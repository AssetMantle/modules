// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/transactions/immediate/message.proto

package immediate

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
	base2 "github.com/AssetMantle/schema/go/lists/base"
	base1 "github.com/AssetMantle/schema/go/types/base"
	_ "github.com/cosmos/gogoproto/gogoproto"
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
	return fileDescriptor_929de1ad5c8fc810, []int{0}
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
	proto.RegisterType((*Message)(nil), "orders.transactions.immediate.Message")
}

func init() {
	proto.RegisterFile("x/orders/internal/transactions/immediate/message.proto", fileDescriptor_929de1ad5c8fc810)
}

var fileDescriptor_929de1ad5c8fc810 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x9b, 0x6c, 0x6b, 0x57, 0x6f, 0x03, 0x9a, 0x6e, 0x5a, 0x37, 0xd4, 0x6c, 0xe3, 0xb4,
	0x09, 0x64, 0x4b, 0x20, 0x71, 0xa8, 0xb8, 0xac, 0xab, 0x18, 0x11, 0x54, 0xad, 0xc2, 0x4e, 0xa8,
	0x52, 0xe5, 0x36, 0x5e, 0x6b, 0x91, 0xc4, 0x51, 0xfc, 0x55, 0xb4, 0x6f, 0xc1, 0x33, 0x70, 0x42,
	0x3c, 0x09, 0xe2, 0xb4, 0x23, 0x47, 0xd4, 0xde, 0x38, 0xf0, 0x0c, 0x28, 0x4e, 0x16, 0xd2, 0xb5,
	0x9d, 0x38, 0x35, 0xd5, 0xff, 0xf7, 0xfb, 0xfc, 0xf9, 0x93, 0x6d, 0xf4, 0x72, 0x4c, 0x44, 0xe8,
	0xb0, 0x50, 0x12, 0xee, 0x03, 0x0b, 0x7d, 0xea, 0x12, 0x08, 0xa9, 0x2f, 0x69, 0x1f, 0xb8, 0xf0,
	0x25, 0xe1, 0x9e, 0xc7, 0x1c, 0x4e, 0x81, 0x11, 0x8f, 0x49, 0x49, 0x07, 0x0c, 0x07, 0xa1, 0x00,
	0x61, 0x54, 0x63, 0x0b, 0x67, 0x61, 0x9c, 0xc2, 0x87, 0xbb, 0x03, 0x31, 0x10, 0x8a, 0x24, 0xd1,
	0x57, 0x2c, 0x1d, 0x9e, 0x8c, 0x09, 0x77, 0x24, 0xe9, 0x51, 0xc9, 0x48, 0xdf, 0xa5, 0x52, 0xf2,
	0x6b, 0xde, 0xa7, 0x91, 0x6b, 0x35, 0x12, 0xe4, 0x71, 0x06, 0xe1, 0x0e, 0xf3, 0x81, 0xc3, 0x24,
	0x0d, 0xab, 0x99, 0x90, 0xfa, 0x93, 0xd6, 0x27, 0x9f, 0xf6, 0x5c, 0x96, 0xc6, 0x47, 0x63, 0xe2,
	0x72, 0x09, 0x09, 0x10, 0x84, 0x22, 0x60, 0x21, 0x4c, 0xde, 0x71, 0x09, 0x09, 0x70, 0x30, 0x26,
	0x30, 0x09, 0x58, 0x02, 0x0c, 0x19, 0x1f, 0x0c, 0x93, 0xe8, 0xc9, 0x9f, 0x0d, 0x54, 0x68, 0xc6,
	0x3b, 0x34, 0x0c, 0xb4, 0x7e, 0x1d, 0x0a, 0xaf, 0xa2, 0x1d, 0x6b, 0xa7, 0x45, 0x5b, 0x7d, 0x1b,
	0x67, 0x68, 0x33, 0xfa, 0xed, 0xf2, 0xae, 0x53, 0xd1, 0x8f, 0xb5, 0xd3, 0xad, 0xe7, 0x0f, 0x31,
	0x77, 0x24, 0xb6, 0xd2, 0x1e, 0xed, 0x7c, 0x04, 0x58, 0x0d, 0xe3, 0x02, 0x19, 0xf3, 0x9b, 0x53,
	0xd2, 0x9a, 0x92, 0xf6, 0x94, 0x74, 0x71, 0x67, 0xef, 0xf6, 0xa3, 0xbb, 0xd3, 0x30, 0x9e, 0xa2,
	0x22, 0xd0, 0x8f, 0x2c, 0x54, 0xee, 0xfa, 0xf2, 0x05, 0x0b, 0x8a, 0xb0, 0x1a, 0xc6, 0x2b, 0x54,
	0xf2, 0x14, 0x2c, 0xe2, 0x89, 0x28, 0x69, 0x43, 0x49, 0x25, 0x25, 0x9d, 0x67, 0x86, 0x65, 0x3f,
	0x50, 0x6c, 0xfa, 0x3f, 0xb2, 0x61, 0xc1, 0xce, 0xaf, 0xb4, 0x61, 0xde, 0x7e, 0x86, 0x10, 0x1b,
	0x07, 0x3c, 0x64, 0xb2, 0xcb, 0xfd, 0x4a, 0x41, 0x69, 0x3b, 0x58, 0x8d, 0x19, 0xbf, 0x51, 0x13,
	0xb6, 0x8b, 0x09, 0x60, 0xf9, 0x06, 0x46, 0xe5, 0xf9, 0x4e, 0x65, 0xe0, 0x72, 0xa8, 0x6c, 0xaa,
	0x49, 0x97, 0xb2, 0x8d, 0xbd, 0x8f, 0x82, 0x88, 0x87, 0x25, 0x7c, 0x31, 0xe6, 0x61, 0x81, 0x6f,
	0xa1, 0x03, 0xee, 0x79, 0x23, 0x50, 0xac, 0xc7, 0x80, 0x76, 0x93, 0x63, 0xc0, 0x99, 0xac, 0x20,
	0xd5, 0x5c, 0x19, 0xab, 0x43, 0x82, 0xdb, 0x99, 0xf3, 0x61, 0xef, 0xa7, 0x56, 0x93, 0x01, 0x6d,
	0xa7, 0x8e, 0xf1, 0x1a, 0xed, 0xfe, 0x2b, 0x98, 0xa9, 0xb5, 0xb5, 0xba, 0x56, 0x39, 0x15, 0x32,
	0x75, 0xde, 0xa2, 0xfd, 0x55, 0x6d, 0x6d, 0xaf, 0x2e, 0xb5, 0xb7, 0xbc, 0xa9, 0x3a, 0x32, 0x96,
	0xb4, 0xb4, 0xb3, 0xba, 0x4e, 0x69, 0xa1, 0xa1, 0xfa, 0x57, 0xfd, 0xfb, 0xd4, 0xd4, 0x6e, 0xa6,
	0xa6, 0xf6, 0x6b, 0x6a, 0x6a, 0x9f, 0x67, 0x66, 0xee, 0x66, 0x66, 0xe6, 0x7e, 0xce, 0xcc, 0x1c,
	0x3a, 0xe9, 0x0b, 0x0f, 0xdf, 0x7b, 0xbf, 0xeb, 0xdb, 0xc9, 0x5d, 0x69, 0x47, 0x97, 0xa7, 0xad,
	0x7d, 0xb8, 0x1c, 0x70, 0x18, 0x8e, 0x7a, 0xb8, 0x2f, 0x3c, 0x72, 0x2e, 0x25, 0x83, 0x26, 0xf5,
	0xc1, 0x65, 0xc4, 0x13, 0xce, 0xc8, 0x65, 0x92, 0xfc, 0xef, 0x2b, 0xf3, 0x45, 0x5f, 0x6b, 0x5d,
	0x59, 0xdf, 0xf4, 0x6a, 0x2b, 0x5e, 0xfe, 0x2a, 0xbb, 0xbc, 0x75, 0x4b, 0xfd, 0xb8, 0xcd, 0x3b,
	0xd9, 0xbc, 0x93, 0xe6, 0x53, 0xfd, 0xec, 0xde, 0xbc, 0x73, 0xd9, 0xae, 0x47, 0x23, 0x75, 0x28,
	0xd0, 0xdf, 0xfa, 0x51, 0xcc, 0xd6, 0x6a, 0x59, 0xb8, 0x56, 0x4b, 0xe9, 0x5e, 0x5e, 0x3d, 0x11,
	0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x8c, 0x5f, 0xe0, 0x2c, 0x05, 0x00, 0x00,
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