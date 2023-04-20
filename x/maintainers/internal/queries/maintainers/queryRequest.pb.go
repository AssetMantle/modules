// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/maintainers/internal/queries/maintainers/queryRequest.proto

package maintainers

import (
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryRequest struct {
	PageRequest *query.PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_236e9fcc26f690b0, []int{0}
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

func (m *QueryRequest) GetPageRequest() *query.PageRequest {
	if m != nil {
		return m.PageRequest
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "maintainers.queries.maintainers.QueryRequest")
}

func init() {
	proto.RegisterFile("x/maintainers/internal/queries/maintainers/queryRequest.proto", fileDescriptor_236e9fcc26f690b0)
}

var fileDescriptor_236e9fcc26f690b0 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x7b, 0x11, 0x1c, 0xd2, 0x2e, 0x76, 0x12, 0x87, 0xab, 0x7f, 0x40, 0x44, 0xe1, 0x8e,
	0xea, 0x16, 0x70, 0xb0, 0x8b, 0x74, 0x08, 0xa4, 0xdd, 0x94, 0x82, 0xbc, 0x69, 0x5f, 0xe2, 0x41,
	0x73, 0x97, 0xe6, 0x2e, 0x62, 0xbf, 0x85, 0x9f, 0xc1, 0xd1, 0x6f, 0xe0, 0x37, 0x10, 0xa7, 0x8e,
	0x8e, 0x92, 0x6c, 0x7e, 0x0a, 0x49, 0x73, 0xe0, 0x75, 0xaa, 0xe3, 0xfb, 0x3c, 0xcf, 0xfb, 0x7b,
	0x9f, 0xe4, 0xfc, 0xeb, 0x67, 0x9e, 0x82, 0x90, 0x06, 0x84, 0xc4, 0x5c, 0x73, 0x21, 0x0d, 0xe6,
	0x12, 0xe6, 0x7c, 0x51, 0x60, 0x2e, 0x50, 0x6f, 0x98, 0xb5, 0xb6, 0x1c, 0xe3, 0xa2, 0x40, 0x6d,
	0x58, 0x96, 0x2b, 0xa3, 0xba, 0x3d, 0xc7, 0x67, 0x76, 0x87, 0x39, 0xda, 0xc1, 0xf9, 0x54, 0xe9,
	0x54, 0x69, 0x1e, 0x83, 0xc6, 0x06, 0xc0, 0x9f, 0xfa, 0x31, 0x1a, 0xe8, 0xf3, 0x0c, 0x12, 0x21,
	0xc1, 0x08, 0x25, 0x1b, 0xd8, 0xf1, 0x9d, 0xdf, 0x19, 0x39, 0x27, 0xba, 0x43, 0xbf, 0x93, 0x41,
	0x82, 0x0f, 0x79, 0x33, 0xef, 0x93, 0x43, 0x72, 0xd6, 0xbe, 0x3c, 0x65, 0x0d, 0x92, 0xd5, 0xc8,
	0xf5, 0xcd, 0x25, 0xb3, 0x48, 0x16, 0x41, 0x82, 0x76, 0x7b, 0xdc, 0xce, 0xfe, 0x86, 0xc1, 0xbb,
	0xf7, 0x51, 0x52, 0xb2, 0x2a, 0x29, 0xf9, 0x2e, 0x29, 0x79, 0xa9, 0x68, 0x6b, 0x55, 0xd1, 0xd6,
	0x57, 0x45, 0x5b, 0xfe, 0xc9, 0x54, 0xa5, 0x6c, 0xcb, 0x67, 0x0c, 0xf6, 0xdc, 0x62, 0x51, 0xdd,
	0x36, 0x22, 0xf7, 0xc3, 0x44, 0x98, 0xc7, 0x22, 0x66, 0x53, 0x95, 0xf2, 0x1b, 0xad, 0xd1, 0x84,
	0x20, 0xcd, 0x1c, 0x79, 0xaa, 0x66, 0xc5, 0x1c, 0x35, 0xff, 0xff, 0xaf, 0x7d, 0xf5, 0x76, 0xc2,
	0x51, 0xf8, 0xe6, 0xf5, 0x42, 0xa7, 0xc7, 0xc8, 0xf6, 0x70, 0xb4, 0xcf, 0x8d, 0xc4, 0xc4, 0x26,
	0x26, 0x8e, 0x56, 0x7a, 0x17, 0x5b, 0x12, 0x93, 0xdb, 0x68, 0x10, 0xa2, 0x81, 0x19, 0x18, 0xf8,
	0xf1, 0x8e, 0x1c, 0x27, 0x08, 0x6c, 0x3c, 0x08, 0x1c, 0x35, 0xde, 0x5d, 0xbf, 0xce, 0xd5, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xc5, 0xc7, 0x9a, 0x2b, 0x02, 0x00, 0x00,
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
	if m.PageRequest != nil {
		{
			size, err := m.PageRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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
	if m.PageRequest != nil {
		l = m.PageRequest.Size()
		n += 1 + l + sovQueryRequest(uint64(l))
	}
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRequest
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
				return ErrInvalidLengthQueryRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PageRequest == nil {
				m.PageRequest = &query.PageRequest{}
			}
			if err := m.PageRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
