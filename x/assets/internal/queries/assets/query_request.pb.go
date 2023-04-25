// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/queries/assets/query_request.proto

package assets

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
	return fileDescriptor_db700eca3409a8d0, []int{0}
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
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.assets.queries.assets.QueryRequest")
}

func init() {
	proto.RegisterFile("x/assets/internal/queries/assets/query_request.proto", fileDescriptor_db700eca3409a8d0)
}

var fileDescriptor_db700eca3409a8d0 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4a, 0xc3, 0x50,
	0x14, 0xc6, 0x9b, 0x80, 0x0e, 0x69, 0x17, 0x3b, 0x89, 0x43, 0x10, 0x07, 0xb1, 0x42, 0xef, 0xa5,
	0xfe, 0x59, 0xb2, 0xc8, 0xed, 0x22, 0x0e, 0x81, 0xb4, 0x9b, 0x12, 0x90, 0x93, 0xf6, 0x10, 0x03,
	0x4d, 0x6e, 0x9a, 0x7b, 0x23, 0xf6, 0x2d, 0x7c, 0x06, 0x47, 0x9f, 0x44, 0x9c, 0x3a, 0x3a, 0x4a,
	0x8a, 0x8b, 0x4f, 0x21, 0xc9, 0x3d, 0xc5, 0x3a, 0xd9, 0xf1, 0x3b, 0xe7, 0xf7, 0x7d, 0xe7, 0x4b,
	0xae, 0x73, 0xf1, 0xc4, 0x41, 0x29, 0xd4, 0x8a, 0x27, 0x99, 0xc6, 0x22, 0x83, 0x19, 0x9f, 0x97,
	0x58, 0x24, 0xa8, 0xd6, 0xf3, 0x5a, 0x2e, 0xee, 0x0b, 0x9c, 0x97, 0xa8, 0x34, 0xcb, 0x0b, 0xa9,
	0x65, 0xb7, 0xd7, 0xec, 0x52, 0xc8, 0xf4, 0x0c, 0x59, 0x2a, 0xa7, 0xe5, 0x0c, 0x15, 0x33, 0x3c,
	0x23, 0x3b, 0xc9, 0x83, 0xd3, 0x89, 0x54, 0xa9, 0x54, 0x3c, 0x02, 0x85, 0x26, 0x8b, 0x3f, 0x0e,
	0x22, 0xd4, 0x30, 0xe0, 0x39, 0xc4, 0x49, 0x06, 0x3a, 0x91, 0x99, 0x89, 0x3d, 0xba, 0x75, 0x3a,
	0xa3, 0x9a, 0x18, 0x9b, 0x63, 0xdd, 0x1b, 0xa7, 0x93, 0x43, 0x8c, 0xeb, 0xe3, 0xfb, 0xd6, 0xa1,
	0x75, 0xd2, 0x3e, 0x3b, 0x66, 0x26, 0x92, 0xd5, 0x91, 0xcd, 0xb9, 0x05, 0xa3, 0x48, 0x16, 0x40,
	0x8c, 0xe4, 0x1e, 0xb7, 0xf3, 0x5f, 0x31, 0xfc, 0xb2, 0xdf, 0x2a, 0xd7, 0x5a, 0x56, 0xae, 0xf5,
	0x59, 0xb9, 0xd6, 0xf3, 0xca, 0x6d, 0x2d, 0x57, 0x6e, 0xeb, 0x63, 0xe5, 0xb6, 0x9c, 0xfe, 0x44,
	0xa6, 0x6c, 0xeb, 0x0f, 0x1a, 0xee, 0x6d, 0x56, 0x0c, 0xea, 0xde, 0x81, 0x75, 0x77, 0x15, 0x27,
	0xfa, 0xa1, 0x8c, 0xd8, 0x44, 0xa6, 0x5c, 0xd4, 0x9c, 0xdf, 0x44, 0x71, 0x8a, 0xe2, 0xff, 0xfd,
	0xe5, 0x17, 0x7b, 0x47, 0xf8, 0x62, 0x24, 0x5e, 0xed, 0x9e, 0xd8, 0x68, 0xe2, 0x53, 0x13, 0x61,
	0x9a, 0x8c, 0xa8, 0x89, 0x91, 0xef, 0x7f, 0xd8, 0x90, 0xd8, 0xd0, 0x2c, 0x43, 0x62, 0x49, 0x56,
	0xf6, 0xe5, 0xd6, 0x6c, 0x78, 0x1d, 0x0c, 0x7d, 0xd4, 0x30, 0x05, 0x0d, 0xdf, 0x76, 0x7f, 0xc3,
	0xe7, 0x79, 0x64, 0xf4, 0x3c, 0x83, 0x7a, 0x1e, 0x59, 0xd7, 0x83, 0x68, 0xb7, 0x79, 0xc9, 0xf3,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xcb, 0xa3, 0x42, 0x58, 0x02, 0x00, 0x00,
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