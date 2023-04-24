// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/queries/identities/queryRequest.proto

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
	return fileDescriptor_ae3877e46cf633b6, []int{0}
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
	proto.RegisterFile("x/identities/internal/queries/identities/queryRequest.proto", fileDescriptor_ae3877e46cf633b6)
}

var fileDescriptor_ae3877e46cf633b6 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4f, 0xb3, 0x40,
	0x18, 0xc7, 0x0b, 0x6f, 0x5e, 0x07, 0xda, 0xc5, 0x4e, 0xc6, 0x81, 0x18, 0x07, 0x63, 0x1c, 0xee,
	0x82, 0xa6, 0x0b, 0xc6, 0xa1, 0x5d, 0x1a, 0x06, 0x12, 0xda, 0x4d, 0x43, 0x62, 0x8e, 0xf6, 0x09,
	0x5e, 0x52, 0x38, 0xca, 0x3d, 0x18, 0xfb, 0x2d, 0x5c, 0xfc, 0x02, 0x8e, 0x7e, 0x12, 0xe3, 0xd4,
	0xd1, 0xd1, 0xd0, 0xcd, 0x4f, 0x61, 0x80, 0x4b, 0xee, 0xdc, 0xec, 0xf8, 0xfc, 0x79, 0x7e, 0xbf,
	0xe7, 0x0f, 0x38, 0xd7, 0x4f, 0x94, 0x2f, 0x21, 0x47, 0x8e, 0x1c, 0x24, 0xe5, 0x39, 0x42, 0x99,
	0xb3, 0x15, 0x5d, 0x57, 0x50, 0xb6, 0x81, 0x7e, 0xd6, 0x44, 0x9b, 0x39, 0xac, 0x2b, 0x90, 0x48,
	0x8a, 0x52, 0xa0, 0x18, 0x7a, 0x4c, 0x4a, 0xc0, 0x8c, 0xe5, 0xb8, 0x02, 0x92, 0x89, 0x65, 0xb5,
	0x02, 0x49, 0x34, 0x42, 0x94, 0xc5, 0x88, 0x8e, 0x2f, 0x16, 0x42, 0x66, 0x42, 0xd2, 0x84, 0x49,
	0xe8, 0x94, 0xf4, 0xd1, 0x4b, 0x00, 0x99, 0x47, 0x0b, 0x96, 0xf2, 0x9c, 0x21, 0x17, 0x79, 0xa7,
	0x3f, 0xbd, 0x75, 0x06, 0x33, 0xe3, 0xe8, 0x30, 0x70, 0x06, 0x05, 0x4b, 0xe1, 0xbe, 0xec, 0xe6,
	0x23, 0xeb, 0xc4, 0x3a, 0xef, 0x5f, 0x9e, 0x91, 0x4e, 0x49, 0x1a, 0x65, 0x7b, 0x72, 0x43, 0x94,
	0x92, 0x44, 0x2c, 0x05, 0x45, 0xcf, 0xfb, 0x85, 0x1e, 0x26, 0x2f, 0xff, 0xde, 0x6b, 0xd7, 0xda,
	0xd6, 0xae, 0xf5, 0x55, 0xbb, 0xd6, 0xf3, 0xce, 0xed, 0x6d, 0x77, 0x6e, 0xef, 0x73, 0xe7, 0xf6,
	0x9c, 0xd1, 0x42, 0x64, 0x64, 0xef, 0x17, 0x9b, 0x1c, 0x9a, 0x55, 0xa3, 0xa6, 0x7f, 0x64, 0xdd,
	0x4d, 0x53, 0x8e, 0x0f, 0x55, 0x42, 0x16, 0x22, 0xa3, 0xe3, 0x46, 0x19, 0xb6, 0x4a, 0xaa, 0x94,
	0xf4, 0xaf, 0x1f, 0xff, 0xd5, 0xfe, 0x3f, 0x0e, 0x83, 0x59, 0xf0, 0x66, 0x7b, 0x63, 0xa3, 0x59,
	0xa8, 0x9a, 0x05, 0xba, 0xd9, 0x4c, 0x35, 0xd3, 0xd1, 0xc7, 0x2f, 0x26, 0x56, 0x4c, 0xac, 0x17,
	0x62, 0xc5, 0x18, 0x51, 0x6d, 0xdf, 0xec, 0xcd, 0xc4, 0xd3, 0x68, 0x12, 0x02, 0xb2, 0x25, 0x43,
	0xf6, 0x6d, 0x8f, 0x0c, 0xde, 0xf7, 0x95, 0xc0, 0xf7, 0xf5, 0xba, 0xef, 0x2b, 0x85, 0x19, 0x26,
	0x07, 0xed, 0x9f, 0xbf, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xaf, 0x10, 0xf5, 0x97, 0x02,
	0x00, 0x00,
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
