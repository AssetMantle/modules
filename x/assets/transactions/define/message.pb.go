// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/transactions/define/message.proto

package define

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/ids/base"
	base1 "github.com/AssetMantle/schema/go/lists/base"
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

type Message struct {
	From                    string              `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID                  *base.IdentityID    `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ImmutableMetaProperties *base1.PropertyList `protobuf:"bytes,3,opt,name=immutable_meta_properties,json=immutableMetaProperties,proto3" json:"immutable_meta_properties,omitempty"`
	ImmutableProperties     *base1.PropertyList `protobuf:"bytes,4,opt,name=immutable_properties,json=immutableProperties,proto3" json:"immutable_properties,omitempty"`
	MutableMetaProperties   *base1.PropertyList `protobuf:"bytes,5,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3" json:"mutable_meta_properties,omitempty"`
	MutableProperties       *base1.PropertyList `protobuf:"bytes,6,opt,name=mutable_properties,json=mutableProperties,proto3" json:"mutable_properties,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_05473ae6b21cc1ec, []int{0}
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

func (m *Message) GetImmutableMetaProperties() *base1.PropertyList {
	if m != nil {
		return m.ImmutableMetaProperties
	}
	return nil
}

func (m *Message) GetImmutableProperties() *base1.PropertyList {
	if m != nil {
		return m.ImmutableProperties
	}
	return nil
}

func (m *Message) GetMutableMetaProperties() *base1.PropertyList {
	if m != nil {
		return m.MutableMetaProperties
	}
	return nil
}

func (m *Message) GetMutableProperties() *base1.PropertyList {
	if m != nil {
		return m.MutableProperties
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "assetmantle.modules.assets.transactions.define.Message")
}

func init() {
	proto.RegisterFile("assets/transactions/define/message.proto", fileDescriptor_05473ae6b21cc1ec)
}

var fileDescriptor_05473ae6b21cc1ec = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x31, 0x8b, 0xdb, 0x30,
	0x18, 0x86, 0x63, 0xdf, 0x5d, 0xda, 0xaa, 0x5d, 0xea, 0xb6, 0x5c, 0x9a, 0xc1, 0x1c, 0x5d, 0x1a,
	0x28, 0x48, 0x90, 0x6e, 0xea, 0xe4, 0x10, 0x28, 0x81, 0x1a, 0x4c, 0xb8, 0xe9, 0x30, 0x35, 0x8a,
	0xa5, 0xcb, 0x09, 0x22, 0x2b, 0xf8, 0x53, 0xa0, 0xf7, 0x2f, 0xfa, 0x1b, 0x3a, 0xf6, 0x97, 0x94,
	0x4e, 0x37, 0x76, 0x2c, 0xc9, 0x76, 0x7b, 0xf7, 0x62, 0x4b, 0xa4, 0x3a, 0x9a, 0x0c, 0x9e, 0x6c,
	0x24, 0x3f, 0xcf, 0xfb, 0x7e, 0x58, 0x42, 0x23, 0x06, 0x20, 0x0c, 0x10, 0x53, 0xb3, 0x0a, 0x58,
	0x69, 0xa4, 0xae, 0x80, 0x70, 0x71, 0x2d, 0x2b, 0x41, 0x94, 0x00, 0x60, 0x4b, 0x81, 0xd7, 0xb5,
	0x36, 0x3a, 0xc2, 0xed, 0x97, 0x8a, 0x55, 0x66, 0x25, 0xb0, 0xd2, 0x7c, 0xb3, 0x12, 0x60, 0xd7,
	0x00, 0xfb, 0x34, 0xb6, 0xf4, 0x70, 0x28, 0x39, 0x90, 0x05, 0x03, 0x41, 0x24, 0x17, 0x95, 0x91,
	0xe6, 0xb6, 0x90, 0xdc, 0xba, 0x86, 0xf1, 0x4a, 0x82, 0x71, 0xbb, 0xeb, 0x5a, 0xaf, 0x45, 0x6d,
	0x6e, 0x8b, 0x66, 0xcd, 0xee, 0xbf, 0xb9, 0x3f, 0x41, 0x8f, 0x52, 0x9b, 0x1e, 0x45, 0xe8, 0xf4,
	0xba, 0xd6, 0x6a, 0x10, 0x5c, 0x04, 0xa3, 0x27, 0xf3, 0xf6, 0x3d, 0x4a, 0xd0, 0xe3, 0xe6, 0x59,
	0xc8, 0x82, 0x0f, 0xc2, 0x8b, 0x60, 0xf4, 0x74, 0xfc, 0xf6, 0x41, 0x3d, 0x28, 0x6f, 0x84, 0x62,
	0x58, 0x72, 0xc0, 0x4d, 0x06, 0x9e, 0xb9, 0x06, 0xb3, 0xe9, 0xbc, 0xdf, 0x80, 0xb3, 0x69, 0xb4,
	0x44, 0xaf, 0xa5, 0x52, 0x1b, 0xc3, 0x16, 0x2b, 0x51, 0x28, 0x61, 0x58, 0xe1, 0x8a, 0x48, 0x01,
	0x83, 0x93, 0xd6, 0xf9, 0xee, 0x90, 0xb3, 0x6d, 0x6e, 0xad, 0x99, 0x6b, 0xfe, 0x49, 0x82, 0x99,
	0x9f, 0xef, 0x6d, 0xa9, 0x30, 0x2c, 0xdb, 0xbb, 0xa2, 0xcf, 0xe8, 0xe5, 0xbf, 0x20, 0x2f, 0xe3,
	0xb4, 0x7b, 0xc6, 0x8b, 0xbd, 0xc8, 0xf3, 0x97, 0xe8, 0xfc, 0xd8, 0x18, 0x67, 0xdd, 0x23, 0x5e,
	0x1d, 0x1e, 0xe2, 0x0a, 0x45, 0x07, 0x46, 0xe8, 0x77, 0xf7, 0x3f, 0xff, 0x6f, 0x80, 0xc9, 0x9f,
	0xf0, 0xc7, 0x36, 0x0e, 0xee, 0xb6, 0x71, 0xf0, 0x7b, 0x1b, 0x07, 0x5f, 0x77, 0x71, 0xef, 0x6e,
	0x17, 0xf7, 0x7e, 0xed, 0xe2, 0x1e, 0x1a, 0x97, 0x5a, 0x75, 0x3c, 0x77, 0x93, 0x67, 0xee, 0xe0,
	0x64, 0xcd, 0x49, 0xca, 0x82, 0xab, 0x0f, 0x4b, 0x69, 0x6e, 0x36, 0x0b, 0x5c, 0x6a, 0x45, 0x92,
	0x06, 0x4b, 0x5b, 0x15, 0x71, 0x2a, 0xf2, 0x85, 0x1c, 0xbf, 0x02, 0xdf, 0xc2, 0xb3, 0x24, 0x4d,
	0x2e, 0xa7, 0xdf, 0x43, 0x9c, 0x78, 0x2d, 0x52, 0xd7, 0x22, 0xb1, 0x2d, 0x2e, 0xfd, 0x16, 0xd3,
	0x16, 0xfc, 0xf9, 0x00, 0xc8, 0x1d, 0x90, 0x5b, 0x20, 0xf7, 0x81, 0xdc, 0x02, 0xdb, 0x90, 0x76,
	0x03, 0xf2, 0x8f, 0xd9, 0xa4, 0xf9, 0x49, 0x9c, 0x19, 0x76, 0x1f, 0x8e, 0x3d, 0x98, 0x52, 0x47,
	0x53, 0xab, 0x04, 0x4a, 0x7d, 0x9e, 0x52, 0x2b, 0x58, 0xf4, 0xdb, 0xbb, 0xf6, 0xfe, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0xbc, 0x27, 0x4c, 0x03, 0x04, 0x00, 0x00,
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
		dAtA[i] = 0x32
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
		dAtA[i] = 0x2a
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
		dAtA[i] = 0x22
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
				m.ImmutableMetaProperties = &base1.PropertyList{}
			}
			if err := m.ImmutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
				m.ImmutableProperties = &base1.PropertyList{}
			}
			if err := m.ImmutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
				m.MutableMetaProperties = &base1.PropertyList{}
			}
			if err := m.MutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
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
				m.MutableProperties = &base1.PropertyList{}
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
