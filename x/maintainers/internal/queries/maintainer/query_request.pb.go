// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/maintainers/internal/queries/maintainer/query_request.proto

package maintainer

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

type QueryRequest struct {
	MaintainerID *base.MaintainerID `protobuf:"bytes,1,opt,name=maintainer_i_d,json=maintainerID,proto3" json:"maintainer_i_d,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21370ca36c06d04, []int{0}
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
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.maintainers.queries.maintainer.QueryRequest")
}

func init() {
	proto.RegisterFile("x/maintainers/internal/queries/maintainer/query_request.proto", fileDescriptor_d21370ca36c06d04)
}

var fileDescriptor_d21370ca36c06d04 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x93, 0xbc, 0xaf, 0x0e, 0xb1, 0x08, 0x16, 0x07, 0x29, 0x12, 0xc5, 0x49, 0x97, 0x3b,
	0xa8, 0xd0, 0xe1, 0x40, 0xa1, 0x45, 0x50, 0x87, 0xc3, 0xb6, 0xa3, 0x04, 0xca, 0xb5, 0x79, 0x68,
	0x0f, 0x7a, 0x39, 0x9b, 0xbb, 0x80, 0x7e, 0x03, 0x47, 0x77, 0x17, 0x71, 0xf4, 0x93, 0x88, 0x53,
	0x47, 0x47, 0x49, 0x37, 0x3f, 0x85, 0x24, 0x3d, 0xf4, 0x71, 0xb3, 0x5b, 0xf2, 0xbf, 0xfc, 0x7e,
	0xcf, 0xff, 0x21, 0x17, 0x9e, 0xdc, 0x52, 0x25, 0x64, 0x6a, 0x85, 0x4c, 0x21, 0x33, 0x54, 0xa6,
	0x16, 0xb2, 0x54, 0x4c, 0xe9, 0x2c, 0x87, 0x4c, 0x82, 0x41, 0x87, 0x55, 0x74, 0x37, 0xc8, 0x60,
	0x96, 0x83, 0xb1, 0xe4, 0x26, 0xd3, 0x56, 0xd7, 0x9b, 0xc2, 0x18, 0xb0, 0x4a, 0xa4, 0x76, 0x0a,
	0x44, 0xe9, 0x24, 0x9f, 0x82, 0x21, 0x48, 0x48, 0x9c, 0x07, 0x65, 0x8d, 0xed, 0xb1, 0x1e, 0xeb,
	0x0a, 0xa7, 0xe5, 0xd3, 0xd2, 0xd4, 0xd8, 0x95, 0x89, 0xa1, 0x43, 0x61, 0x00, 0x8d, 0x1c, 0xc8,
	0x64, 0x79, 0x7a, 0x00, 0x61, 0xad, 0x57, 0x8e, 0xef, 0x2f, 0xa7, 0xd7, 0xaf, 0xc2, 0x4d, 0xfc,
	0xd9, 0x20, 0xd9, 0xf1, 0xf7, 0xfd, 0xc3, 0x8d, 0xe6, 0x11, 0xc1, 0x85, 0xcc, 0x68, 0x02, 0x4a,
	0x10, 0x99, 0x18, 0x52, 0x9a, 0x09, 0xff, 0x46, 0x2e, 0xcf, 0xfa, 0x35, 0x85, 0xde, 0xd8, 0xff,
	0xfb, 0xa7, 0x3d, 0xaf, 0xf3, 0xf8, 0xef, 0xb5, 0x88, 0xfc, 0x79, 0x11, 0xf9, 0x1f, 0x45, 0xe4,
	0x3f, 0x2c, 0x22, 0x6f, 0xbe, 0x88, 0xbc, 0xf7, 0x45, 0xe4, 0x85, 0xad, 0x91, 0x56, 0x64, 0xf5,
	0x6d, 0x3b, 0x5b, 0xb8, 0x77, 0xb7, 0x5c, 0xa6, 0xeb, 0x5f, 0x5f, 0x8c, 0xa5, 0x9d, 0xe4, 0x43,
	0x32, 0xd2, 0x8a, 0xb6, 0x4b, 0x27, 0xaf, 0x9c, 0xd4, 0x39, 0xe9, 0x9f, 0x7f, 0xca, 0x73, 0xb0,
	0xd6, 0xe6, 0xbc, 0xc7, 0x5f, 0x82, 0x66, 0x1b, 0x75, 0xe3, 0xae, 0x1b, 0x47, 0xdd, 0x7a, 0xae,
	0xdb, 0x4f, 0xf6, 0xf6, 0x0b, 0x8a, 0x1d, 0x14, 0x23, 0x28, 0x76, 0x10, 0xca, 0x8a, 0xe0, 0x74,
	0x75, 0x28, 0x3e, 0xef, 0x76, 0x38, 0x58, 0x91, 0x08, 0x2b, 0x3e, 0x83, 0x16, 0x12, 0x30, 0xe6,
	0x0c, 0x8c, 0x21, 0x05, 0x63, 0xce, 0x81, 0xd3, 0xe1, 0x7a, 0x75, 0x17, 0x8e, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0xab, 0x3a, 0x77, 0xb4, 0x02, 0x00, 0x00,
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
	if m.MaintainerID != nil {
		{
			size, err := m.MaintainerID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.MaintainerID != nil {
		l = m.MaintainerID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field MaintainerID", wireType)
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
			if m.MaintainerID == nil {
				m.MaintainerID = &base.MaintainerID{}
			}
			if err := m.MaintainerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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