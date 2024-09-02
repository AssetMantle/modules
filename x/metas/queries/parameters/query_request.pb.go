// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/metas/queries/parameters/query_request.proto

package parameters

import (
	fmt "fmt"
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

type QueryRequest struct {
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4aa1936e6a1ea67, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryRequest)(nil), "AssetMantle.modules.x.metas.queries.parameters.QueryRequest")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/metas/queries/parameters/query_request.proto", fileDescriptor_a4aa1936e6a1ea67)
}

var fileDescriptor_a4aa1936e6a1ea67 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd1, 0xb1, 0x4a, 0x03, 0x31,
	0x1c, 0xc7, 0xf1, 0x5e, 0x86, 0x0e, 0x87, 0x08, 0xfa, 0x00, 0x19, 0x7c, 0x80, 0x04, 0xec, 0xe4,
	0x7f, 0xeb, 0x2d, 0x4e, 0x81, 0x9c, 0xd3, 0x21, 0x01, 0x49, 0xdb, 0x3f, 0x5a, 0x68, 0x4c, 0x9b,
	0xe4, 0xa0, 0xbe, 0x85, 0xcf, 0xe0, 0xe8, 0x93, 0x88, 0x53, 0x47, 0x47, 0xb9, 0xdb, 0x7c, 0x04,
	0x27, 0xb9, 0x34, 0x78, 0xd7, 0xf1, 0xd6, 0x84, 0xef, 0x87, 0x5f, 0x48, 0x5e, 0xcc, 0xbd, 0xc7,
	0x20, 0xf4, 0x73, 0xd8, 0x20, 0x37, 0x76, 0x55, 0x6f, 0xd0, 0xf3, 0x3d, 0x37, 0x18, 0xb4, 0xe7,
	0xbb, 0x1a, 0xdd, 0x1a, 0x3d, 0xdf, 0x6a, 0xa7, 0x0d, 0x06, 0x74, 0xc7, 0xa3, 0x97, 0x07, 0x87,
	0xbb, 0x1a, 0x7d, 0x60, 0x5b, 0x67, 0x83, 0xbd, 0x64, 0x03, 0x83, 0x25, 0x83, 0xed, 0x59, 0x34,
	0x58, 0x32, 0x58, 0x6f, 0x5c, 0x9d, 0xe7, 0x67, 0x65, 0xc7, 0xdc, 0x1d, 0x95, 0xe2, 0x97, 0x7c,
	0x34, 0x34, 0x3b, 0x34, 0x34, 0xfb, 0x6e, 0x68, 0xf6, 0xda, 0xd2, 0xc9, 0xa1, 0xa5, 0x93, 0xaf,
	0x96, 0x4e, 0xf2, 0xeb, 0xa5, 0x35, 0x23, 0xf9, 0xe2, 0x62, 0x88, 0xcb, 0x6e, 0xa1, 0xcc, 0xee,
	0x6f, 0x1e, 0xd7, 0xe1, 0xa9, 0x5e, 0xb0, 0xa5, 0x35, 0x7c, 0xdc, 0x93, 0xdf, 0xc8, 0x74, 0x2e,
	0x2a, 0x51, 0xca, 0x77, 0x72, 0xb2, 0x44, 0xa4, 0x25, 0x15, 0x13, 0x71, 0x49, 0x99, 0x96, 0xc8,
	0xff, 0xf2, 0xf3, 0x24, 0x50, 0x29, 0x50, 0x95, 0x8a, 0x81, 0x4a, 0x81, 0xea, 0x83, 0x86, 0xc0,
	0xb8, 0x40, 0xdd, 0xca, 0xa2, 0xbb, 0x5b, 0xe9, 0xa0, 0x7f, 0xc8, 0x6c, 0x10, 0x03, 0xa4, 0x1a,
	0xa0, 0x02, 0x88, 0x3d, 0x40, 0x02, 0x00, 0x7a, 0x61, 0x31, 0x8d, 0x7f, 0x38, 0xfb, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x06, 0x4b, 0x1a, 0x74, 0x09, 0x02, 0x00, 0x00,
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintQueryRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovQueryRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryRequest(x uint64) (n int) {
	return sovQueryRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRequest
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
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRequest
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
func skipQueryRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryRequest
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
					return 0, ErrIntOverflowQueryRequest
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
					return 0, ErrIntOverflowQueryRequest
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
				return 0, ErrInvalidLengthQueryRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryRequest = fmt.Errorf("proto: unexpected end of group")
)