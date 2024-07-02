// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/provision/message.proto

package provision

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/ids/base"
	proto "github.com/cosmos/gogoproto/proto"
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
	From       string           `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To         string           `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	IdentityID *base.IdentityID `protobuf:"bytes,3,opt,name=identity_i_d,json=identityID,proto3" json:"identity_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_dddda0e3e861dfd5, []int{0}
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

func (m *Message) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Message) GetIdentityID() *base.IdentityID {
	if m != nil {
		return m.IdentityID
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "AssetMantle.modules.x.identities.transactions.provision.Message")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/provision/message.proto", fileDescriptor_dddda0e3e861dfd5)
}

var fileDescriptor_dddda0e3e861dfd5 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x9b, 0xab, 0x54, 0x3c, 0x8b, 0x43, 0xa6, 0xe2, 0x10, 0x8a, 0x8b, 0x5d, 0xbc, 0x03,
	0x1d, 0x84, 0xd3, 0xa5, 0x51, 0x91, 0x0c, 0x81, 0x50, 0x3a, 0x04, 0x09, 0x94, 0x6b, 0xee, 0x6c,
	0x0f, 0x9a, 0x5c, 0xc9, 0x5d, 0xa5, 0xbe, 0x85, 0xcf, 0xe0, 0xa8, 0x2f, 0x22, 0x4e, 0x1d, 0x1d,
	0x25, 0xdd, 0x7c, 0x0a, 0x69, 0x9b, 0x26, 0xd7, 0xb5, 0xdb, 0x47, 0xe0, 0xf7, 0xfb, 0xff, 0xbf,
	0x7c, 0x07, 0x1f, 0xba, 0x4a, 0x71, 0xed, 0xd3, 0x54, 0x4f, 0x38, 0x4e, 0x24, 0x9b, 0x4d, 0xb8,
	0xc2, 0x73, 0x2c, 0x18, 0x4f, 0xb5, 0xd0, 0x82, 0x2b, 0xac, 0x33, 0x9a, 0x2a, 0x1a, 0x6b, 0x21,
	0x53, 0x85, 0xa7, 0x99, 0x7c, 0x11, 0x4a, 0xc8, 0x14, 0x27, 0x5c, 0x29, 0x3a, 0xe2, 0x68, 0x9a,
	0x49, 0x2d, 0xed, 0x6b, 0x43, 0x83, 0x0a, 0x0d, 0x9a, 0xa3, 0x4a, 0x83, 0x4c, 0x0d, 0x2a, 0x35,
	0xa7, 0x17, 0x66, 0xbe, 0x8a, 0xc7, 0x3c, 0xa1, 0x58, 0x30, 0x85, 0x87, 0x54, 0xf1, 0x6d, 0x8b,
	0xd7, 0x81, 0x60, 0x9b, 0x9c, 0xb3, 0x39, 0x3c, 0xf4, 0x37, 0xc1, 0xb6, 0x0d, 0x0f, 0x9e, 0x33,
	0x99, 0xb4, 0xac, 0xb6, 0xd5, 0x39, 0xea, 0xad, 0x67, 0xfb, 0x04, 0x02, 0x2d, 0x5b, 0x60, 0xfd,
	0x05, 0x68, 0x69, 0x7b, 0xb0, 0x59, 0x39, 0x06, 0xac, 0x55, 0x6f, 0x5b, 0x9d, 0xe3, 0xcb, 0x73,
	0x64, 0xb6, 0xdd, 0x84, 0x22, 0xc1, 0x14, 0x5a, 0x85, 0x22, 0xaf, 0x00, 0xbc, 0xfb, 0x1e, 0x14,
	0xe5, 0xec, 0x7e, 0xd6, 0xbf, 0x72, 0xc7, 0x5a, 0xe4, 0x8e, 0xf5, 0x9b, 0x3b, 0xd6, 0xdb, 0xd2,
	0xa9, 0x2d, 0x96, 0x4e, 0xed, 0x67, 0xe9, 0xd4, 0xe0, 0x4d, 0x2c, 0x13, 0xb4, 0xe7, 0x0f, 0x70,
	0x9b, 0xc5, 0x3e, 0xc1, 0x6a, 0xbf, 0xc0, 0x7a, 0x72, 0x47, 0x42, 0x8f, 0x67, 0x43, 0x14, 0xcb,
	0x04, 0xef, 0x79, 0x9b, 0x77, 0xd0, 0xe8, 0xfa, 0xa1, 0xd7, 0x0f, 0x3e, 0xc0, 0xce, 0x5d, 0xfc,
	0xa2, 0x56, 0xb8, 0xdd, 0x71, 0x55, 0xab, 0x6f, 0xd6, 0x0a, 0xb6, 0x8a, 0xef, 0x1d, 0x32, 0x2a,
	0xc8, 0x28, 0x8c, 0x2a, 0x32, 0x32, 0xc9, 0xa8, 0x24, 0x73, 0x70, 0xb7, 0x27, 0x19, 0x3d, 0x06,
	0xae, 0xcf, 0x35, 0x65, 0x54, 0xd3, 0x3f, 0x70, 0x6b, 0x58, 0x08, 0x29, 0x34, 0x84, 0x84, 0x84,
	0x54, 0x22, 0x42, 0x4c, 0x13, 0x21, 0xa5, 0x6a, 0xd8, 0x58, 0x3f, 0x97, 0xab, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x51, 0x85, 0x1a, 0xdf, 0x02, 0x00, 0x00,
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
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.To)))
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
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.IdentityID != nil {
		l = m.IdentityID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			m.To = string(dAtA[iNdEx:postIndex])
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
