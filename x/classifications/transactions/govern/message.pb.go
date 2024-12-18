// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/classifications/transactions/govern/message.proto

package govern

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/parameters/base"
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
	From      string          `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Parameter *base.Parameter `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_caddba2240d01829, []int{0}
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

func (m *Message) GetParameter() *base.Parameter {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "AssetMantle.modules.x.classifications.transactions.govern.Message")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/classifications/transactions/govern/message.proto", fileDescriptor_caddba2240d01829)
}

var fileDescriptor_caddba2240d01829 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4e, 0xc2, 0x40,
	0x18, 0xc7, 0xb9, 0x6a, 0x30, 0x54, 0xa7, 0x4e, 0xc4, 0xa1, 0x21, 0x4e, 0x2c, 0xde, 0x25, 0x38,
	0x79, 0x89, 0x26, 0x80, 0xa6, 0x83, 0x69, 0xd2, 0x10, 0x86, 0xc6, 0x74, 0xf9, 0x28, 0x07, 0xd4,
	0x70, 0x3d, 0xd2, 0x3b, 0x0c, 0x8f, 0xe1, 0x33, 0x38, 0xfa, 0x02, 0xbe, 0x82, 0x71, 0x62, 0x74,
	0x34, 0x65, 0xf3, 0x29, 0x0c, 0x5c, 0x95, 0xa3, 0x63, 0xb7, 0xfb, 0x86, 0xdf, 0xef, 0xff, 0xff,
	0xf2, 0x9d, 0xed, 0x75, 0xa5, 0x64, 0xca, 0x87, 0x54, 0xcd, 0x19, 0xe1, 0x62, 0xbc, 0x9c, 0x33,
	0x49, 0x56, 0x24, 0x9e, 0x83, 0x94, 0xc9, 0x24, 0x89, 0x41, 0x25, 0x22, 0x95, 0x44, 0x65, 0x90,
	0x4a, 0x88, 0xf5, 0x30, 0x15, 0xcf, 0x2c, 0x4b, 0x09, 0x67, 0x52, 0xc2, 0x94, 0xe1, 0x45, 0x26,
	0x94, 0x70, 0xae, 0x0d, 0x11, 0x2e, 0x44, 0x78, 0x85, 0x4b, 0x22, 0x6c, 0x8a, 0xb0, 0x16, 0x9d,
	0x77, 0xcc, 0x0e, 0x32, 0x9e, 0x31, 0x0e, 0x64, 0x01, 0x19, 0x70, 0xa6, 0x58, 0x26, 0xc9, 0x08,
	0x24, 0xdb, 0xcf, 0x3a, 0xee, 0xe2, 0xc9, 0x3e, 0xf1, 0x75, 0xbe, 0xe3, 0xd8, 0xc7, 0x93, 0x4c,
	0xf0, 0x26, 0x6a, 0xa1, 0x76, 0x63, 0xb0, 0x7b, 0x3b, 0x0f, 0x76, 0xe3, 0x9f, 0x68, 0x5a, 0x2d,
	0xd4, 0x3e, 0xed, 0x5c, 0x62, 0xb3, 0xa1, 0x8e, 0xc1, 0xfb, 0x18, 0xbc, 0x8d, 0xc1, 0xc1, 0xdf,
	0x3c, 0xd8, 0xf3, 0xbd, 0xf7, 0xa3, 0x8f, 0xdc, 0x45, 0xeb, 0xdc, 0x45, 0xdf, 0xb9, 0x8b, 0x5e,
	0x36, 0x6e, 0x6d, 0xbd, 0x71, 0x6b, 0x5f, 0x1b, 0xb7, 0x66, 0xdf, 0xc4, 0x82, 0xe3, 0xca, 0x9b,
	0xf7, 0xce, 0x8a, 0x1d, 0x82, 0xed, 0x4e, 0x01, 0x7a, 0xbc, 0x9b, 0x26, 0x6a, 0xb6, 0x1c, 0xe1,
	0x58, 0x70, 0x52, 0xf9, 0x30, 0xaf, 0x56, 0xbd, 0xeb, 0x87, 0xfd, 0xa1, 0xf7, 0x66, 0x1d, 0x1c,
	0xc5, 0x2f, 0xaa, 0x85, 0xb8, 0x5f, 0xaa, 0x36, 0x34, 0xab, 0x79, 0x3b, 0xc9, 0xe7, 0x01, 0x1b,
	0x15, 0x6c, 0x14, 0x46, 0x25, 0x36, 0x32, 0xd9, 0x48, 0xb3, 0xb9, 0x75, 0x5f, 0x99, 0x8d, 0xbc,
	0xa0, 0xe7, 0x33, 0x05, 0x63, 0x50, 0xf0, 0x63, 0xdd, 0x1a, 0x1e, 0x4a, 0x0b, 0x11, 0xa5, 0x21,
	0xa5, 0x25, 0x15, 0xa5, 0xa6, 0x8b, 0x52, 0x2d, 0x1b, 0xd5, 0x77, 0x9f, 0xe5, 0xea, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xf4, 0xbc, 0xfe, 0xf7, 0xe6, 0x02, 0x00, 0x00,
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
	if m.Parameter != nil {
		{
			size, err := m.Parameter.MarshalToSizedBuffer(dAtA[:i])
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
	if m.Parameter != nil {
		l = m.Parameter.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
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
			if m.Parameter == nil {
				m.Parameter = &base.Parameter{}
			}
			if err := m.Parameter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
