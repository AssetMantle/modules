// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/transactions/revoke/message.proto

package revoke

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
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
	From             string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID           *base.IdentityID       `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ToID             *base.IdentityID       `protobuf:"bytes,3,opt,name=to_i_d,json=toID,proto3" json:"to_i_d,omitempty"`
	ClassificationID *base.ClassificationID `protobuf:"bytes,4,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_e82dcb877ec2bfe2, []int{0}
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

func (m *Message) GetToID() *base.IdentityID {
	if m != nil {
		return m.ToID
	}
	return nil
}

func (m *Message) GetClassificationID() *base.ClassificationID {
	if m != nil {
		return m.ClassificationID
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "assets.transactions.revoke.Message")
}

func init() {
	proto.RegisterFile("x/assets/internal/transactions/revoke/message.proto", fileDescriptor_e82dcb877ec2bfe2)
}

var fileDescriptor_e82dcb877ec2bfe2 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x3b, 0x69, 0xe9, 0xbd, 0x37, 0xf7, 0xc2, 0x95, 0x01, 0xa1, 0x44, 0x19, 0xaa, 0xa0,
	0xd4, 0xcd, 0x0c, 0xd8, 0x5d, 0x76, 0x6d, 0x03, 0x92, 0x45, 0xa1, 0x84, 0xae, 0xa4, 0x50, 0xa6,
	0xc9, 0xb4, 0x0e, 0x26, 0x19, 0xc9, 0x4c, 0xa5, 0xbe, 0x85, 0x2b, 0x1f, 0xc0, 0xa5, 0x6b, 0x1f,
	0x42, 0x5c, 0x75, 0xe9, 0x52, 0xd2, 0x9d, 0x4f, 0x21, 0x99, 0xb1, 0x12, 0x8b, 0x05, 0x57, 0x09,
	0xfc, 0xdf, 0xc7, 0x7f, 0xce, 0x1c, 0xbb, 0xbd, 0x20, 0x54, 0x4a, 0xa6, 0x24, 0xe1, 0xa9, 0x62,
	0x59, 0x4a, 0x63, 0xa2, 0x32, 0x9a, 0x4a, 0x1a, 0x2a, 0x2e, 0x52, 0x49, 0x32, 0x76, 0x2d, 0x2e,
	0x19, 0x49, 0x98, 0x94, 0x74, 0xc6, 0xf0, 0x55, 0x26, 0x94, 0x80, 0x8e, 0x51, 0x70, 0x99, 0xc4,
	0x86, 0x74, 0xf6, 0x16, 0x84, 0x47, 0x92, 0x4c, 0xa8, 0x64, 0x84, 0x47, 0x2c, 0x55, 0x5c, 0xdd,
	0xf8, 0x9e, 0x11, 0x9d, 0x83, 0x52, 0x18, 0xc6, 0x54, 0x4a, 0x3e, 0xe5, 0x21, 0x2d, 0xfc, 0x35,
	0x72, 0xf8, 0x08, 0xec, 0x5f, 0x7d, 0xd3, 0x06, 0xa1, 0x5d, 0x9b, 0x66, 0x22, 0x69, 0x80, 0x26,
	0x68, 0xfd, 0x09, 0xf4, 0x3f, 0x3c, 0xb1, 0x7f, 0x17, 0xdf, 0x31, 0x1f, 0x47, 0x0d, 0xab, 0x09,
	0x5a, 0x7f, 0x4f, 0xff, 0x63, 0x1e, 0x49, 0xec, 0x7f, 0x76, 0x05, 0xf5, 0x02, 0xf0, 0x3d, 0x78,
	0x64, 0xd7, 0x95, 0xd0, 0x60, 0xf5, 0x7b, 0xb0, 0xa6, 0x84, 0xef, 0xc1, 0x9e, 0x0d, 0xbf, 0xce,
	0xa2, 0x95, 0x9a, 0x56, 0x76, 0xb5, 0xd2, 0xdb, 0x18, 0x35, 0xd8, 0xd9, 0x1c, 0xbe, 0x7b, 0x67,
	0x3d, 0xe5, 0x08, 0x2c, 0x73, 0x04, 0x5e, 0x73, 0x04, 0x6e, 0x57, 0xa8, 0xb2, 0x5c, 0xa1, 0xca,
	0xcb, 0x0a, 0x55, 0x6c, 0x14, 0x8a, 0x04, 0x6f, 0x7f, 0xb1, 0xee, 0xbf, 0x8f, 0x75, 0x07, 0xc5,
	0xfe, 0x03, 0x70, 0xee, 0xcd, 0xb8, 0xba, 0x98, 0x4f, 0x70, 0x28, 0x12, 0xd2, 0x29, 0xb4, 0x3e,
	0x4d, 0x55, 0xcc, 0x48, 0x22, 0xa2, 0x79, 0xcc, 0x24, 0xf9, 0xd1, 0xc5, 0xee, 0xad, 0x6a, 0x67,
	0x18, 0x3c, 0x58, 0x4e, 0xc7, 0x14, 0x0f, 0xcb, 0xc5, 0x81, 0x46, 0x9e, 0xd7, 0xe1, 0xa8, 0x1c,
	0x8e, 0x4c, 0x98, 0x5b, 0xc7, 0xdb, 0xc3, 0xd1, 0xd9, 0xa0, 0xdb, 0x67, 0x8a, 0x46, 0x54, 0xd1,
	0x37, 0x6b, 0xdf, 0x80, 0xae, 0x5b, 0x26, 0x5d, 0xd7, 0xa0, 0x93, 0xba, 0x3e, 0x6b, 0xfb, 0x3d,
	0x00, 0x00, 0xff, 0xff, 0x0a, 0x87, 0xfa, 0xd9, 0x69, 0x02, 0x00, 0x00,
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
		dAtA[i] = 0x22
	}
	if m.ToID != nil {
		{
			size, err := m.ToID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.ToID != nil {
		l = m.ToID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field ToID", wireType)
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
			if m.ToID == nil {
				m.ToID = &base.IdentityID{}
			}
			if err := m.ToID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
