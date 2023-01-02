// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/nub/response.v1.proto

package nub

import (
	fmt "fmt"
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

type Response struct {
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_581fa9a231842dd7, []int{0}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Response)(nil), "identities.transactions.nub.Response")
}

func init() {
	proto.RegisterFile("modules/identities/internal/transactions/nub/response.v1.proto", fileDescriptor_581fa9a231842dd7)
}

var fileDescriptor_581fa9a231842dd7 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xcb, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x04, 0x31, 0xf3, 0x4a,
	0x52, 0x8b, 0xf2, 0x12, 0x73, 0xf4, 0x4b, 0x8a, 0x12, 0xf3, 0x8a, 0x13, 0x93, 0x4b, 0x32, 0xf3,
	0xf3, 0x8a, 0xf5, 0xf3, 0x4a, 0x93, 0xf4, 0x8b, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xf5,
	0xca, 0x0c, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xa4, 0x11, 0xfa, 0xf4, 0x90, 0x95, 0xeb,
	0xe5, 0x95, 0x26, 0x29, 0x71, 0x71, 0x71, 0x04, 0x41, 0x75, 0x38, 0x2d, 0x65, 0x3a, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0xf9, 0xe4, 0xfc, 0x5c, 0x3d, 0x3c, 0xe6, 0x38,
	0xf1, 0xc3, 0x4c, 0x09, 0x33, 0x0c, 0x00, 0xd9, 0x1a, 0xc0, 0x18, 0xe5, 0x9d, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x58, 0x5c, 0x9c, 0x5a, 0xe2, 0x9b, 0x98, 0x57,
	0x92, 0x93, 0xaa, 0x0f, 0xf3, 0x0e, 0x29, 0xde, 0x5a, 0xc4, 0xc4, 0xec, 0x19, 0xe2, 0xb7, 0x8a,
	0x49, 0xda, 0x13, 0xe1, 0x86, 0x10, 0x64, 0x37, 0xf8, 0x95, 0x26, 0x9d, 0x42, 0x96, 0x8d, 0x41,
	0x96, 0x8d, 0xf1, 0x2b, 0x4d, 0x7a, 0xc4, 0xa4, 0x8e, 0x47, 0x36, 0xc6, 0x3d, 0xc0, 0xc9, 0x37,
	0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0xf1, 0x15, 0x93, 0x2c, 0x42, 0xa5, 0x95, 0x15, 0xb2, 0x52,
	0x2b, 0x2b, 0xbf, 0xd2, 0xa4, 0x24, 0x36, 0x70, 0xb8, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x43, 0x41, 0x62, 0x2f, 0x99, 0x01, 0x00, 0x00,
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintResponseV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovResponseV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovResponseV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResponseV1(x uint64) (n int) {
	return sovResponseV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponseV1
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipResponseV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResponseV1
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
func skipResponseV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResponseV1
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
					return 0, ErrIntOverflowResponseV1
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
					return 0, ErrIntOverflowResponseV1
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
				return 0, ErrInvalidLengthResponseV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResponseV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResponseV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResponseV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResponseV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResponseV1 = fmt.Errorf("proto: unexpected end of group")
)