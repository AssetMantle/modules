// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/splits/internal/transactions/unwrap/message.proto

package unwrap

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
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
	From      string             `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID    *base.IdentityID   `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	OwnableID *base.AnyOwnableID `protobuf:"bytes,3,opt,name=ownable_i_d,json=ownableID,proto3" json:"ownable_i_d,omitempty"`
	Value     string             `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_925527723f6a8b34, []int{0}
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
	proto.RegisterType((*Message)(nil), "splits.transactions.unwrap.Message")
}

func init() {
	proto.RegisterFile("x/splits/internal/transactions/unwrap/message.proto", fileDescriptor_925527723f6a8b34)
}

var fileDescriptor_925527723f6a8b34 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x6e, 0xea, 0x30,
	0x14, 0x87, 0x71, 0xe0, 0x72, 0x2f, 0xe1, 0x4a, 0x57, 0x37, 0x62, 0x40, 0x69, 0x1b, 0xa1, 0x0e,
	0x15, 0x5d, 0x6c, 0xb5, 0x6c, 0xd9, 0x40, 0x48, 0x15, 0x03, 0x2a, 0xe2, 0xcf, 0x52, 0x21, 0x21,
	0x87, 0xb8, 0xa9, 0xa5, 0xc4, 0x46, 0xb1, 0x53, 0xe0, 0x2d, 0xda, 0xa5, 0x0f, 0xd0, 0xb1, 0x4f,
	0x52, 0x75, 0x62, 0xec, 0x58, 0x85, 0xad, 0x4f, 0x51, 0xc5, 0x86, 0x2a, 0x0b, 0x52, 0x27, 0x1f,
	0xeb, 0xfb, 0x8e, 0x7f, 0x3a, 0x3e, 0x66, 0x6b, 0x85, 0xc4, 0x22, 0xa4, 0x52, 0x20, 0xca, 0x24,
	0x89, 0x19, 0x0e, 0x91, 0x8c, 0x31, 0x13, 0x78, 0x2e, 0x29, 0x67, 0x02, 0x25, 0x6c, 0x19, 0xe3,
	0x05, 0x8a, 0x88, 0x10, 0x38, 0x20, 0x70, 0x11, 0x73, 0xc9, 0x2d, 0x5b, 0xb7, 0xc0, 0xbc, 0x09,
	0xb5, 0x69, 0xd7, 0x02, 0x1e, 0x70, 0xa5, 0xa1, 0xac, 0xd2, 0x1d, 0xf6, 0xd1, 0x0a, 0x51, 0x5f,
	0x20, 0x0f, 0x0b, 0x82, 0xa8, 0x4f, 0x98, 0xa4, 0x72, 0xdd, 0xeb, 0xee, 0xe0, 0x49, 0x0e, 0x62,
	0xb6, 0xbe, 0x5e, 0x32, 0xec, 0x85, 0x64, 0x8f, 0x4f, 0x1f, 0x81, 0xf9, 0xbb, 0xaf, 0xf3, 0x2d,
	0xcb, 0x2c, 0xdd, 0xc6, 0x3c, 0xaa, 0x83, 0x06, 0x68, 0x56, 0x86, 0xaa, 0xb6, 0xce, 0xcd, 0x3f,
	0xd9, 0x39, 0xa3, 0x33, 0xbf, 0x6e, 0x34, 0x40, 0xb3, 0x7a, 0xf9, 0x0f, 0x52, 0x5f, 0xc0, 0xde,
	0x77, 0xce, 0xb0, 0x9c, 0x09, 0xbd, 0xae, 0x75, 0x61, 0x56, 0xb9, 0x7e, 0x5d, 0xd9, 0x45, 0x65,
	0xff, 0x57, 0x76, 0x3b, 0x17, 0x3c, 0xac, 0xf0, 0x7d, 0x69, 0xd5, 0xcc, 0x5f, 0xf7, 0x38, 0x4c,
	0x48, 0xbd, 0xa4, 0x22, 0xf5, 0xa5, 0xf3, 0x64, 0xbc, 0xa6, 0x0e, 0xd8, 0xa4, 0x0e, 0xf8, 0x48,
	0x1d, 0xf0, 0xb0, 0x75, 0x0a, 0x9b, 0xad, 0x53, 0x78, 0xdf, 0x3a, 0x05, 0xd3, 0x99, 0xf3, 0x08,
	0x1e, 0xfe, 0xa0, 0xce, 0xdf, 0xdd, 0x2c, 0x83, 0x6c, 0xb8, 0x01, 0xb8, 0xe9, 0x06, 0x54, 0xde,
	0x25, 0x1e, 0x9c, 0xf3, 0x08, 0xb5, 0x85, 0x20, 0xb2, 0x8f, 0x99, 0x0c, 0x09, 0x8a, 0xb8, 0x9f,
	0x84, 0x44, 0xa0, 0x1f, 0x2d, 0xe8, 0xd9, 0x28, 0x8e, 0xc6, 0x93, 0x17, 0xc3, 0x1e, 0xe9, 0xe0,
	0x71, 0x3e, 0x78, 0xa2, 0x94, 0xb7, 0x3d, 0x9c, 0xe6, 0xe1, 0x54, 0xc3, 0xd4, 0x38, 0x3b, 0x0c,
	0xa7, 0x57, 0x83, 0x4e, 0x9f, 0x48, 0xec, 0x63, 0x89, 0x3f, 0x8d, 0x63, 0x2d, 0xba, 0x6e, 0xde,
	0x74, 0x5d, 0xad, 0x7a, 0x65, 0xb5, 0xb3, 0xd6, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0xa4,
	0x87, 0x3f, 0x58, 0x02, 0x00, 0x00,
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