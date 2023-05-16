// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metas/transactions/reveal/message.proto

package reveal

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/data/base"
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
	From string        `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Data *base.AnyData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_20a8135a6bddd6eb, []int{0}
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

func (m *Message) GetData() *base.AnyData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "assetmantle.modules.metas.transactions.reveal.Message")
}

func init() {
	proto.RegisterFile("metas/transactions/reveal/message.proto", fileDescriptor_20a8135a6bddd6eb)
}

var fileDescriptor_20a8135a6bddd6eb = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0x3b, 0x31,
	0x1c, 0xc7, 0x9b, 0xd0, 0xff, 0x5f, 0x3c, 0x9d, 0x6e, 0x2a, 0x0e, 0xa1, 0x08, 0x62, 0x97, 0x26,
	0x54, 0xa7, 0x66, 0xbb, 0x22, 0x38, 0x05, 0xca, 0xd1, 0x41, 0xe4, 0x40, 0x7e, 0xbd, 0xc6, 0xb6,
	0xd0, 0x5c, 0xe4, 0x92, 0x8a, 0x7d, 0x0b, 0x9f, 0xc1, 0xd1, 0x27, 0x11, 0xa7, 0x8e, 0x8e, 0x72,
	0xdd, 0x5c, 0x7c, 0x05, 0x49, 0xee, 0x86, 0x38, 0x38, 0x74, 0x4b, 0xc8, 0xf7, 0xf3, 0xc9, 0x37,
	0xbf, 0x44, 0xe7, 0x4a, 0x5a, 0x30, 0xcc, 0x96, 0x50, 0x18, 0xc8, 0xed, 0x52, 0x17, 0x86, 0x95,
	0xf2, 0x51, 0xc2, 0x8a, 0x29, 0x69, 0x0c, 0xcc, 0x25, 0x7d, 0x28, 0xb5, 0xd5, 0x71, 0x1f, 0x8c,
	0x91, 0x56, 0x41, 0x61, 0x57, 0x92, 0x2a, 0x3d, 0x5b, 0xaf, 0xa4, 0xa1, 0x1e, 0xa6, 0x21, 0x4c,
	0x6b, 0xf8, 0xa4, 0x33, 0x03, 0x0b, 0x6c, 0x0a, 0x46, 0x32, 0x28, 0x36, 0x77, 0x6e, 0x57, 0x8b,
	0x4e, 0x6f, 0xa2, 0x03, 0x51, 0x9b, 0xe3, 0x38, 0x6a, 0xdf, 0x97, 0x5a, 0x75, 0x50, 0x17, 0xf5,
	0x0e, 0x53, 0xbf, 0x8e, 0x87, 0x51, 0xdb, 0x85, 0x3b, 0xb8, 0x8b, 0x7a, 0x47, 0x17, 0x67, 0x34,
	0xbc, 0xd6, 0xe4, 0x0b, 0xa9, 0x80, 0x7a, 0x99, 0x53, 0xd3, 0xa4, 0xd8, 0x5c, 0x81, 0x85, 0xd4,
	0x23, 0xa3, 0x6f, 0xfc, 0x56, 0x11, 0xb4, 0xad, 0x08, 0xfa, 0xac, 0x08, 0x7a, 0xde, 0x91, 0xd6,
	0x76, 0x47, 0x5a, 0x1f, 0x3b, 0xd2, 0x8a, 0x06, 0xb9, 0x56, 0x74, 0xaf, 0x17, 0x8c, 0x8e, 0x9b,
	0x96, 0x63, 0xd7, 0x7a, 0x8c, 0x6e, 0xf9, 0x7c, 0x69, 0x17, 0xeb, 0x29, 0xcd, 0xb5, 0x62, 0x89,
	0x33, 0x09, 0x6f, 0x62, 0x8d, 0x89, 0x3d, 0xb1, 0x3f, 0x47, 0xf9, 0x82, 0xff, 0x25, 0x42, 0x4c,
	0xd2, 0x57, 0xdc, 0x4f, 0x82, 0x0e, 0xa2, 0xe9, 0x20, 0x7c, 0x87, 0x49, 0xd8, 0x21, 0xf5, 0xdc,
	0xfb, 0xaf, 0x7c, 0xd6, 0xe4, 0x33, 0x9f, 0xcf, 0xc2, 0x7c, 0x56, 0xe7, 0x2b, 0x3c, 0xdc, 0x2b,
	0x9f, 0x5d, 0x8f, 0x47, 0xee, 0xd0, 0x4d, 0xef, 0x0b, 0x0f, 0x02, 0x96, 0xf3, 0x06, 0xe6, 0xdc,
	0xd3, 0x9c, 0x87, 0x38, 0xe7, 0x35, 0x3f, 0xfd, 0xef, 0xbf, 0xf4, 0xf2, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x86, 0x25, 0x57, 0x46, 0x02, 0x00, 0x00,
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
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
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
	if m.Data != nil {
		l = m.Data.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			if m.Data == nil {
				m.Data = &base.AnyData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
