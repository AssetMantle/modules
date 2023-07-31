// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: splits/queries/supply/query_request.proto

package supply

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/go/ids/base"
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
	OwnableID *base.AnyOwnableID `protobuf:"bytes,1,opt,name=ownable_i_d,json=ownableID,proto3" json:"ownable_i_d,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec9e50cf3f7dbd34, []int{0}
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

func (m *QueryRequest) GetOwnableID() *base.AnyOwnableID {
	if m != nil {
		return m.OwnableID
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.splits.queries.supply.QueryRequest")
}

func init() {
	proto.RegisterFile("splits/queries/supply/query_request.proto", fileDescriptor_ec9e50cf3f7dbd34)
}

var fileDescriptor_ec9e50cf3f7dbd34 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x9b, 0xc0, 0xfb, 0x82, 0xa9, 0x8b, 0x9d, 0x44, 0x30, 0x88, 0x93, 0x19, 0x7a, 0x07,
	0x8a, 0x0e, 0xb7, 0xa5, 0x08, 0xd2, 0x21, 0xd8, 0x36, 0x93, 0x12, 0x08, 0x97, 0xde, 0x61, 0x0f,
	0x92, 0x5c, 0xda, 0xe7, 0x82, 0xe6, 0x5b, 0xf8, 0x19, 0x1c, 0xfd, 0x24, 0xe2, 0xd4, 0xd1, 0xc1,
	0x41, 0xd2, 0xcd, 0x4f, 0x21, 0xc9, 0x9d, 0x12, 0xc1, 0xa1, 0xdb, 0x1d, 0xfc, 0x7e, 0xff, 0xe7,
	0x7f, 0xf7, 0x38, 0x1e, 0x14, 0xa9, 0x50, 0x80, 0x97, 0x25, 0x5f, 0x09, 0x0e, 0x18, 0xca, 0xa2,
	0x48, 0xab, 0xf6, 0x5a, 0xc5, 0x2b, 0xbe, 0x2c, 0x39, 0x28, 0x54, 0xac, 0xa4, 0x92, 0x03, 0x8f,
	0x02, 0x70, 0x95, 0xd1, 0x5c, 0xa5, 0x1c, 0x65, 0x92, 0x95, 0x29, 0x07, 0xa4, 0x75, 0x64, 0x74,
	0xa4, 0xf5, 0x83, 0x43, 0xc1, 0x00, 0x27, 0x14, 0x38, 0xa6, 0x79, 0x15, 0xcb, 0xfb, 0x9c, 0x26,
	0x29, 0x8f, 0x05, 0xd3, 0x49, 0xc7, 0x37, 0xce, 0xee, 0xb4, 0x19, 0x30, 0xd3, 0xf9, 0x83, 0xb1,
	0xd3, 0xff, 0x61, 0x62, 0xb6, 0x6f, 0x1d, 0x59, 0x27, 0xfd, 0x53, 0x0f, 0x75, 0xe7, 0xc1, 0x7c,
	0xc1, 0x33, 0x8a, 0x04, 0x03, 0xd4, 0xe4, 0x22, 0x3f, 0xaf, 0xae, 0xb5, 0x32, 0xbe, 0x9c, 0xed,
	0xc8, 0xef, 0xe3, 0xe8, 0xdd, 0x7e, 0xa9, 0x5d, 0x6b, 0x5d, 0xbb, 0xd6, 0x47, 0xed, 0x5a, 0x8f,
	0x1b, 0xb7, 0xb7, 0xde, 0xb8, 0xbd, 0xb7, 0x8d, 0xdb, 0x73, 0x86, 0x73, 0x99, 0xa1, 0xad, 0xdf,
	0x30, 0xda, 0xeb, 0x56, 0x9c, 0x34, 0xbd, 0x27, 0xd6, 0xed, 0xc5, 0x9d, 0x50, 0x8b, 0x32, 0x41,
	0x73, 0x99, 0x61, 0xbf, 0x89, 0x0a, 0xda, 0x28, 0x6c, 0xa2, 0xf0, 0x03, 0xfe, 0xf3, 0x3f, 0x9f,
	0xec, 0x7f, 0x7e, 0x10, 0x4e, 0xc3, 0x67, 0xdb, 0xf3, 0x3b, 0x05, 0x02, 0x53, 0x20, 0xd4, 0x05,
	0xa6, 0xa6, 0x40, 0xd8, 0x3a, 0xaf, 0xbf, 0xd8, 0xc8, 0xb0, 0x91, 0x66, 0x23, 0xc3, 0x46, 0x9a,
	0xad, 0xed, 0xf3, 0xad, 0xd9, 0xe8, 0x6a, 0x32, 0x0a, 0xb8, 0xa2, 0x8c, 0x2a, 0xfa, 0x69, 0x0f,
	0x3b, 0x1e, 0x21, 0x46, 0x24, 0x44, 0x9b, 0x84, 0x18, 0x95, 0x10, 0xed, 0x26, 0xff, 0xdb, 0x05,
	0x9e, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x60, 0x9e, 0x67, 0x70, 0x37, 0x02, 0x00, 0x00,
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
	if m.OwnableID != nil {
		{
			size, err := m.OwnableID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.OwnableID != nil {
		l = m.OwnableID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
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
			if m.OwnableID == nil {
				m.OwnableID = &base.AnyOwnableID{}
			}
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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