// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: splits/transactions/unwrap/message.proto

package unwrap

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
	From      string             `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID    *base.IdentityID   `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	OwnableID *base.AnyOwnableID `protobuf:"bytes,3,opt,name=ownable_i_d,json=ownableID,proto3" json:"ownable_i_d,omitempty"`
	Value     string             `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b614ad2dc1566b2, []int{0}
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

func (m *Message) GetOwnableID() *base.AnyOwnableID {
	if m != nil {
		return m.OwnableID
	}
	return nil
}

func (m *Message) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "assetmantle.modules.splits.transactions.unwrap.Message")
}

func init() {
	proto.RegisterFile("splits/transactions/unwrap/message.proto", fileDescriptor_4b614ad2dc1566b2)
}

var fileDescriptor_4b614ad2dc1566b2 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x3f, 0xcb, 0xd3, 0x40,
	0x1c, 0x07, 0xf0, 0x5e, 0x7c, 0x9e, 0x6a, 0x53, 0xa7, 0xe0, 0x10, 0x0a, 0x86, 0xe2, 0x62, 0x5d,
	0xee, 0xa0, 0x6e, 0xe7, 0x94, 0x52, 0x90, 0x0c, 0xc1, 0xd2, 0x3f, 0x8b, 0x04, 0xc2, 0x25, 0x77,
	0xb6, 0x07, 0xc9, 0x5d, 0xc9, 0x5d, 0xac, 0x7d, 0x17, 0xbe, 0x06, 0x47, 0xdf, 0x85, 0x9b, 0x38,
	0x75, 0x74, 0x94, 0x74, 0x73, 0x77, 0x97, 0xe4, 0x62, 0x49, 0x07, 0x1f, 0xe8, 0x94, 0x0b, 0xc7,
	0xe7, 0xfb, 0xfb, 0xf2, 0x4b, 0xec, 0x89, 0xda, 0x67, 0x5c, 0x2b, 0xa4, 0x0b, 0x22, 0x14, 0x49,
	0x35, 0x97, 0x42, 0xa1, 0x52, 0x1c, 0x0a, 0xb2, 0x47, 0x39, 0x53, 0x8a, 0x6c, 0x19, 0xdc, 0x17,
	0x52, 0x4b, 0x07, 0x12, 0xa5, 0x98, 0xce, 0x89, 0xd0, 0x19, 0x83, 0xb9, 0xa4, 0x65, 0xc6, 0x14,
	0x34, 0x1a, 0x76, 0x35, 0x34, 0x7a, 0x34, 0xe2, 0x54, 0xa1, 0x84, 0x28, 0x86, 0x38, 0x65, 0x42,
	0x73, 0x7d, 0x8c, 0x39, 0x35, 0x59, 0xa3, 0xe7, 0x97, 0x3b, 0x22, 0x8e, 0xb1, 0x3c, 0x08, 0x92,
	0x64, 0xec, 0x72, 0xfd, 0xe2, 0x1b, 0xb0, 0x1f, 0x87, 0x66, 0xb8, 0xe3, 0xd8, 0x77, 0x1f, 0x0a,
	0x99, 0xbb, 0x60, 0x0c, 0x26, 0x83, 0x65, 0x73, 0x76, 0x7c, 0xfb, 0x49, 0xfd, 0x8c, 0x79, 0x4c,
	0x5d, 0x6b, 0x0c, 0x26, 0xc3, 0xe9, 0xcb, 0xab, 0x76, 0x2a, 0xdd, 0xb1, 0x9c, 0x40, 0x4e, 0x15,
	0xac, 0x87, 0xc0, 0xa0, 0x2d, 0x10, 0xcc, 0x97, 0xfd, 0x1a, 0x06, 0x73, 0x27, 0xb0, 0x87, 0x97,
	0xb1, 0x31, 0x75, 0x1f, 0x35, 0x29, 0xaf, 0x1e, 0x4c, 0xf1, 0xc5, 0xf1, 0x9d, 0x21, 0xc1, 0x7c,
	0x39, 0x90, 0xff, 0x8e, 0xce, 0x33, 0xfb, 0xfe, 0x23, 0xc9, 0x4a, 0xe6, 0xde, 0x35, 0x15, 0xcd,
	0xcb, 0xec, 0x8f, 0xf5, 0xbd, 0xf2, 0xc0, 0xa9, 0xf2, 0xc0, 0xaf, 0xca, 0x03, 0x9f, 0xcf, 0x5e,
	0xef, 0x74, 0xf6, 0x7a, 0x3f, 0xcf, 0x5e, 0xcf, 0x9e, 0xa6, 0x32, 0xbf, 0x71, 0x9b, 0xb3, 0xa7,
	0xed, 0x3e, 0x16, 0xf5, 0x82, 0x16, 0xe0, 0xfd, 0x9b, 0x2d, 0xd7, 0xbb, 0x32, 0x81, 0xa9, 0xcc,
	0x91, 0x5f, 0x47, 0x85, 0x4d, 0x14, 0x6a, 0xa3, 0xd0, 0x27, 0xf4, 0xff, 0x0f, 0xfb, 0xc5, 0xba,
	0xf7, 0xc3, 0xd5, 0x7a, 0xf3, 0xd5, 0x82, 0x7e, 0xa7, 0x45, 0xd8, 0xb6, 0x58, 0x99, 0x16, 0xeb,
	0x6e, 0x8b, 0x4d, 0x03, 0x7f, 0x5c, 0x81, 0xa8, 0x05, 0x91, 0x01, 0x51, 0x17, 0x44, 0x06, 0x54,
	0x16, 0xbe, 0x0d, 0x44, 0x6f, 0x17, 0xb3, 0x90, 0x69, 0x42, 0x89, 0x26, 0xbf, 0xad, 0x69, 0x07,
	0x63, 0xdc, 0x6a, 0x8c, 0x0d, 0xc7, 0xb8, 0xeb, 0x31, 0x36, 0x01, 0x49, 0xbf, 0xf9, 0x85, 0x5e,
	0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xc8, 0x05, 0xb3, 0xd9, 0x02, 0x00, 0x00,
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
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if m.OwnableID != nil {
		{
			size, err := m.OwnableID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.OwnableID != nil {
		l = m.OwnableID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
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
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
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
			if m.OwnableID == nil {
				m.OwnableID = &base.AnyOwnableID{}
			}
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			m.Value = string(dAtA[iNdEx:postIndex])
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