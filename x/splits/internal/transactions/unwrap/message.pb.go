// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/splits/internal/transactions/unwrap/message.proto

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
	proto.RegisterType((*Message)(nil), "assetmantle.modules.splits.transactions.unwrap.Message")
}

func init() {
	proto.RegisterFile("x/splits/internal/transactions/unwrap/message.proto", fileDescriptor_925527723f6a8b34)
}

var fileDescriptor_925527723f6a8b34 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0xaa, 0xdb, 0x30,
	0x14, 0x87, 0x23, 0xf7, 0xde, 0xdb, 0x5e, 0xdf, 0x4e, 0xa6, 0x83, 0x09, 0xd4, 0x84, 0x2e, 0x4d,
	0x17, 0x09, 0x92, 0x4d, 0x9b, 0x83, 0xa1, 0x78, 0x30, 0x0d, 0xf9, 0xb3, 0x14, 0x83, 0x91, 0x2d,
	0x35, 0x11, 0xd8, 0x52, 0xb0, 0xe4, 0xa6, 0x79, 0x8b, 0x3e, 0x43, 0xc7, 0xbe, 0x45, 0xb7, 0xd2,
	0x29, 0x63, 0xc7, 0xe2, 0x6c, 0x7d, 0x88, 0x52, 0x6c, 0xb9, 0xc1, 0x59, 0xca, 0xcd, 0xa4, 0x23,
	0x0e, 0xdf, 0x77, 0x7e, 0x1c, 0xc9, 0x9e, 0x7e, 0x42, 0x6a, 0x97, 0x73, 0xad, 0x10, 0x17, 0x9a,
	0x95, 0x82, 0xe4, 0x48, 0x97, 0x44, 0x28, 0x92, 0x69, 0x2e, 0x85, 0x42, 0x95, 0xd8, 0x97, 0x64,
	0x87, 0x0a, 0xa6, 0x14, 0xd9, 0x30, 0xb8, 0x2b, 0xa5, 0x96, 0x0e, 0x24, 0x4a, 0x31, 0x5d, 0x10,
	0xa1, 0x73, 0x06, 0x0b, 0x49, 0xab, 0x9c, 0x29, 0x68, 0x34, 0xb0, 0x4f, 0x43, 0x43, 0x0f, 0x87,
	0x9c, 0x2a, 0x94, 0x12, 0xc5, 0x10, 0xa7, 0x4c, 0x68, 0xae, 0x0f, 0x09, 0xa7, 0xc6, 0x35, 0x7c,
	0x79, 0xee, 0x11, 0x71, 0x48, 0xe4, 0x5e, 0x90, 0x34, 0x67, 0xe7, 0xf6, 0xab, 0x6f, 0xc0, 0x7e,
	0x1a, 0x99, 0xe1, 0x8e, 0x63, 0xdf, 0x7c, 0x28, 0x65, 0xe1, 0x82, 0x11, 0x18, 0xdf, 0x2f, 0xda,
	0xda, 0xf1, 0xed, 0x67, 0xcd, 0x99, 0xf0, 0x84, 0xba, 0xd6, 0x08, 0x8c, 0x1f, 0x26, 0xaf, 0x2f,
	0xd2, 0xa9, 0x6c, 0xcb, 0x0a, 0x02, 0x39, 0x55, 0xb0, 0x19, 0x02, 0xc3, 0x2e, 0x40, 0x18, 0x2c,
	0xee, 0x1a, 0x30, 0x0c, 0x9c, 0xd0, 0x7e, 0x38, 0x8f, 0x4d, 0xa8, 0xfb, 0xa4, 0xb5, 0xbc, 0xf9,
	0xaf, 0xc5, 0x17, 0x87, 0x77, 0x06, 0x09, 0x83, 0xc5, 0xbd, 0xfc, 0x57, 0x3a, 0x2f, 0xec, 0xdb,
	0x8f, 0x24, 0xaf, 0x98, 0x7b, 0xd3, 0x46, 0x34, 0x97, 0xd9, 0x1f, 0xeb, 0x7b, 0xed, 0x81, 0x63,
	0xed, 0x81, 0x5f, 0xb5, 0x07, 0x3e, 0x9f, 0xbc, 0xc1, 0xf1, 0xe4, 0x0d, 0x7e, 0x9e, 0xbc, 0x81,
	0x3d, 0xc9, 0x64, 0x71, 0xe5, 0x36, 0x67, 0xcf, 0xbb, 0x7d, 0xcc, 0x9b, 0x05, 0xcd, 0xc1, 0xfb,
	0x60, 0xc3, 0xf5, 0xb6, 0x4a, 0x61, 0x26, 0x0b, 0xe4, 0x37, 0xaa, 0xa8, 0x55, 0xa1, 0x4e, 0x85,
	0x1e, 0xf5, 0xc2, 0x5f, 0xac, 0x5b, 0x3f, 0x5a, 0xae, 0xd6, 0x5f, 0x2d, 0xe8, 0xf7, 0xe2, 0x44,
	0x5d, 0x9c, 0xa5, 0x89, 0xb3, 0xea, 0xc7, 0x59, 0xb7, 0xe0, 0x8f, 0x0b, 0x20, 0xee, 0x80, 0xd8,
	0x00, 0x71, 0x1f, 0x88, 0x0d, 0x50, 0x5b, 0xf8, 0x3a, 0x20, 0x7e, 0x3b, 0x9f, 0x45, 0x4c, 0x13,
	0x4a, 0x34, 0xf9, 0x6d, 0x4d, 0x7a, 0x30, 0xc6, 0x1d, 0x8d, 0xb1, 0xc1, 0x31, 0xee, 0xf3, 0x18,
	0x1b, 0x41, 0x7a, 0xd7, 0xfe, 0xa5, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0xda, 0x18,
	0xac, 0xed, 0x02, 0x00, 0x00,
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
