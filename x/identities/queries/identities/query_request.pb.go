// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/queries/identities/query_request.proto

package identities

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	query "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryRequest struct {
	PageRequest *query.PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd1c6d7331483876, []int{0}
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
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.identities.queries.identities.QueryRequest")
}

func init() {
	proto.RegisterFile("identities/queries/identities/query_request.proto", fileDescriptor_fd1c6d7331483876)
}

var fileDescriptor_fd1c6d7331483876 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x9b, 0x88, 0x0e, 0x69, 0x17, 0x3b, 0x89, 0xc3, 0x21, 0x0e, 0x22, 0x0e, 0x77, 0x44,
	0xe9, 0x72, 0xa0, 0xd0, 0x2e, 0x92, 0x21, 0x90, 0x76, 0x53, 0x02, 0x72, 0x69, 0x3f, 0xe2, 0x41,
	0x93, 0x4b, 0x7b, 0x17, 0xb1, 0xef, 0xe0, 0xe0, 0x33, 0x38, 0xfa, 0x24, 0xe2, 0xd4, 0xd1, 0x51,
	0xd2, 0xcd, 0xa7, 0x90, 0x24, 0x27, 0x77, 0x3a, 0x08, 0x9d, 0x42, 0xfe, 0x7c, 0xff, 0xdf, 0xf7,
	0xbb, 0x3b, 0xcf, 0xe7, 0x33, 0xc8, 0x15, 0x57, 0x1c, 0x24, 0x59, 0x94, 0xb0, 0xac, 0xbf, 0x7f,
	0xa2, 0xd5, 0xdd, 0x12, 0x16, 0x25, 0x48, 0x85, 0x8b, 0xa5, 0x50, 0xa2, 0xef, 0x33, 0x29, 0x41,
	0x65, 0x2c, 0x57, 0x73, 0xc0, 0x99, 0x98, 0x95, 0x73, 0x90, 0xd8, 0x74, 0xb0, 0xc6, 0x58, 0xd1,
	0xe1, 0xd9, 0x54, 0xc8, 0x4c, 0x48, 0x92, 0x30, 0x09, 0x2d, 0x93, 0x3c, 0xf8, 0x09, 0x28, 0xe6,
	0x93, 0x82, 0xa5, 0x3c, 0x67, 0x8a, 0x8b, 0xbc, 0xc5, 0x1f, 0xdf, 0x78, 0xbd, 0x71, 0x3d, 0x31,
	0x69, 0x97, 0xf6, 0x03, 0xaf, 0x57, 0xb0, 0x14, 0x7e, 0x24, 0x0e, 0x9c, 0x23, 0xe7, 0xb4, 0x7b,
	0x7e, 0x82, 0x5b, 0x24, 0xae, 0x91, 0xcd, 0xca, 0x15, 0xd6, 0x48, 0x1c, 0xb1, 0x14, 0x74, 0x7b,
	0xd2, 0x2d, 0xcc, 0xcf, 0xe8, 0x69, 0xe7, 0xad, 0x42, 0xce, 0xba, 0x42, 0xce, 0x67, 0x85, 0x9c,
	0xe7, 0x0d, 0xea, 0xac, 0x37, 0xa8, 0xf3, 0xb1, 0x41, 0x1d, 0x6f, 0x30, 0x15, 0x19, 0xde, 0xfa,
	0x60, 0xa3, 0x7d, 0x5b, 0x35, 0xaa, 0xfd, 0x23, 0xe7, 0xf6, 0x2a, 0xe5, 0xea, 0xbe, 0x4c, 0xf0,
	0x54, 0x64, 0x64, 0x58, 0x23, 0xc3, 0x06, 0x49, 0x34, 0x92, 0x3c, 0x92, 0x7f, 0x2f, 0xfd, 0xc5,
	0xdd, 0x1d, 0x86, 0xc1, 0x38, 0x78, 0x75, 0xfd, 0xa1, 0x25, 0x14, 0x6a, 0xa1, 0xc0, 0x08, 0x8d,
	0xb5, 0x90, 0x89, 0xde, 0x7f, 0x75, 0x62, 0xdd, 0x89, 0xcd, 0x40, 0xac, 0x3b, 0x56, 0x54, 0xb9,
	0x97, 0x5b, 0x77, 0xe2, 0xeb, 0x68, 0x14, 0x82, 0x62, 0x33, 0xa6, 0xd8, 0x97, 0x3b, 0xb0, 0xfa,
	0x94, 0x6a, 0x00, 0xa5, 0x66, 0x9c, 0x52, 0x8d, 0xb0, 0xc3, 0x64, 0xaf, 0x79, 0xf0, 0x8b, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x4b, 0x69, 0x07, 0x84, 0x02, 0x00, 0x00,
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