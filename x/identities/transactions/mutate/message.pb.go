// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/transactions/mutate/message.proto

package mutate

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
	base1 "github.com/AssetMantle/schema/go/lists/base"
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
	From                  string              `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID                *base.IdentityID    `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	IdentityID            *base.IdentityID    `protobuf:"bytes,3,opt,name=identity_i_d,json=identityID,proto3" json:"identity_i_d,omitempty"`
	MutableMetaProperties *base1.PropertyList `protobuf:"bytes,4,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3" json:"mutable_meta_properties,omitempty"`
	MutableProperties     *base1.PropertyList `protobuf:"bytes,5,opt,name=mutable_properties,json=mutableProperties,proto3" json:"mutable_properties,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_e145789846e40bfe, []int{0}
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

func (m *Message) GetIdentityID() *base.IdentityID {
	if m != nil {
		return m.IdentityID
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
	proto.RegisterType((*Message)(nil), "assetmantle.modules.identities.transactions.mutate.Message")
}

func init() {
	proto.RegisterFile("identities/transactions/mutate/message.proto", fileDescriptor_e145789846e40bfe)
}

var fileDescriptor_e145789846e40bfe = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd3, 0xcf, 0x8a, 0x13, 0x31,
	0x1c, 0x07, 0xf0, 0xce, 0xec, 0x1f, 0x35, 0xee, 0xc5, 0x80, 0x58, 0x7a, 0x18, 0x16, 0x2f, 0x2e,
	0x28, 0x09, 0x54, 0xd8, 0x43, 0x0e, 0x4a, 0xcb, 0x82, 0x0c, 0x38, 0x30, 0x94, 0x3d, 0x2d, 0x03,
	0x43, 0x3a, 0x89, 0xbb, 0x81, 0xa6, 0x29, 0x93, 0x5f, 0xc1, 0x7d, 0x08, 0xc1, 0x67, 0xf0, 0xe8,
	0x93, 0x88, 0xa7, 0x3d, 0x7a, 0x94, 0xe9, 0x4d, 0x7c, 0x08, 0xc9, 0x4c, 0xec, 0x44, 0x90, 0x85,
	0x39, 0x75, 0x68, 0xfa, 0xfd, 0xe4, 0x9b, 0xfc, 0x3a, 0xe8, 0x95, 0x12, 0x72, 0x0d, 0x0a, 0x94,
	0xb4, 0x14, 0x6a, 0xbe, 0xb6, 0xbc, 0x02, 0x65, 0xd6, 0x96, 0xea, 0x2d, 0x70, 0x90, 0x54, 0x4b,
	0x6b, 0xf9, 0xb5, 0x24, 0x9b, 0xda, 0x80, 0xc1, 0x53, 0x6e, 0xad, 0x04, 0xcd, 0xd7, 0xb0, 0x92,
	0x44, 0x1b, 0xb1, 0x5d, 0x49, 0x4b, 0x7a, 0x81, 0x84, 0x02, 0xe9, 0x84, 0xc9, 0x44, 0x09, 0x4b,
	0x97, 0xdc, 0x4a, 0xea, 0x7f, 0x78, 0x5b, 0x2a, 0xd1, 0x79, 0x93, 0x64, 0xa5, 0x2c, 0xf8, 0xd5,
	0x4d, 0x6d, 0x36, 0xb2, 0x86, 0xdb, 0xd2, 0x7d, 0xd7, 0xad, 0x3f, 0xff, 0x1d, 0xa3, 0x07, 0x59,
	0xd7, 0x00, 0x63, 0x74, 0xf8, 0xa1, 0x36, 0x7a, 0x1c, 0x9d, 0x46, 0x67, 0x8f, 0x16, 0xed, 0x33,
	0x9e, 0xa1, 0x87, 0xee, 0xb3, 0x54, 0xa5, 0x18, 0xc7, 0xa7, 0xd1, 0xd9, 0xe3, 0xe9, 0x0b, 0x12,
	0x56, 0xb4, 0xd5, 0x8d, 0xd4, 0x9c, 0x28, 0x61, 0x89, 0xdb, 0x83, 0xa4, 0xbe, 0x41, 0x7a, 0xb1,
	0x38, 0x76, 0xc1, 0xf4, 0x02, 0xa7, 0xe8, 0xa4, 0xef, 0x55, 0x8a, 0xf1, 0xc1, 0x30, 0x06, 0xa9,
	0xfd, 0x33, 0xae, 0xd0, 0x33, 0x77, 0xe6, 0xe5, 0x4a, 0x96, 0x5a, 0x02, 0x2f, 0xfd, 0x89, 0x94,
	0xb4, 0xe3, 0xc3, 0x56, 0x7d, 0xf9, 0x3f, 0xb5, 0xbd, 0x82, 0xce, 0xcd, 0xfd, 0x15, 0xbc, 0x57,
	0x16, 0x16, 0x4f, 0xbd, 0x95, 0x49, 0xe0, 0xf9, 0x5e, 0xc2, 0x57, 0x08, 0xff, 0xdd, 0x24, 0xf0,
	0x8f, 0x86, 0xfb, 0x4f, 0x3c, 0xd3, 0xdb, 0xf3, 0x4f, 0x07, 0xdf, 0x9a, 0x24, 0xba, 0x6b, 0x92,
	0xe8, 0x67, 0x93, 0x44, 0x9f, 0x77, 0xc9, 0xe8, 0x6e, 0x97, 0x8c, 0x7e, 0xec, 0x92, 0x11, 0x3a,
	0xaf, 0x8c, 0x26, 0xc3, 0xa7, 0x3f, 0x3f, 0xf1, 0xe3, 0xcb, 0xdd, 0x3c, 0xf3, 0xe8, 0xea, 0xed,
	0xb5, 0x82, 0x9b, 0xed, 0x92, 0x54, 0x46, 0xd3, 0x99, 0xe3, 0xb2, 0x96, 0xa3, 0x9e, 0xa3, 0x1f,
	0xe9, 0xfd, 0x7f, 0xc8, 0x2f, 0xf1, 0xd1, 0x2c, 0x4b, 0x2f, 0xb3, 0xaf, 0xf1, 0x74, 0x16, 0xb4,
	0xc9, 0x7c, 0x9b, 0xb4, 0x6f, 0x73, 0x19, 0xb6, 0xc9, 0xda, 0xf0, 0xf7, 0x7f, 0x42, 0x85, 0x0f,
	0x15, 0x7d, 0xa8, 0x08, 0x43, 0x45, 0x17, 0x6a, 0xe2, 0x37, 0xc3, 0x43, 0xc5, 0xbb, 0x7c, 0xee,
	0x06, 0x27, 0x38, 0xf0, 0x5f, 0xf1, 0x79, 0x00, 0x30, 0xe6, 0x05, 0xc6, 0x7a, 0x82, 0xb1, 0xd0,
	0x60, 0xac, 0x43, 0x96, 0xc7, 0xed, 0x5b, 0xf0, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x67, 0xe9, 0x7b, 0xa5, 0x03, 0x00, 0x00,
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
		dAtA[i] = 0x2a
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
		dAtA[i] = 0x22
	}
	if m.IdentityID != nil {
		{
			size, err := m.IdentityID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.IdentityID != nil {
		l = m.IdentityID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
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
			if m.IdentityID == nil {
				m.IdentityID = &base.IdentityID{}
			}
			if err := m.IdentityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
		case 5:
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