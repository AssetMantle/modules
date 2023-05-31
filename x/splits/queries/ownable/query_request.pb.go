// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: splits/queries/ownable/query_request.proto

package ownable

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
	return fileDescriptor_7c34f57fa245e1f3, []int{0}
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
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.splits.queries.ownable.QueryRequest")
}

func init() {
	proto.RegisterFile("splits/queries/ownable/query_request.proto", fileDescriptor_7c34f57fa245e1f3)
}

var fileDescriptor_7c34f57fa245e1f3 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x9b, 0x80, 0x82, 0xa9, 0x8b, 0x9d, 0x44, 0x30, 0x88, 0x93, 0x76, 0x78, 0x0f, 0x14,
	0x14, 0x6e, 0x4b, 0x11, 0xa4, 0x43, 0xe8, 0xbf, 0x49, 0x09, 0x84, 0x4b, 0xef, 0xb0, 0x07, 0x49,
	0xae, 0xed, 0x5d, 0xd0, 0x7e, 0x0b, 0x3f, 0x83, 0xa3, 0x9f, 0x44, 0x9c, 0x3a, 0x3a, 0x4a, 0xba,
	0xf9, 0x29, 0x24, 0xb9, 0x57, 0x89, 0xe0, 0xd0, 0x2d, 0x39, 0x7e, 0xbf, 0xe7, 0x7d, 0xde, 0x3b,
	0xaf, 0xab, 0xe7, 0xa9, 0x34, 0x9a, 0x2c, 0x0a, 0xb1, 0x94, 0x42, 0x13, 0xf5, 0x98, 0xb3, 0x24,
	0x15, 0xf5, 0xff, 0x2a, 0x5e, 0x8a, 0x45, 0x21, 0xb4, 0x81, 0xf9, 0x52, 0x19, 0xd5, 0xe9, 0x32,
	0xad, 0x85, 0xc9, 0x58, 0x6e, 0x52, 0x01, 0x99, 0xe2, 0x45, 0x2a, 0x34, 0x58, 0x1f, 0xd0, 0x07,
	0xf4, 0x8f, 0x8e, 0x25, 0xd7, 0x24, 0x61, 0x5a, 0x10, 0x96, 0xaf, 0x62, 0x3c, 0x8d, 0x25, 0xb7,
	0x51, 0xa7, 0x77, 0xde, 0xfe, 0xa8, 0x9a, 0x30, 0xb6, 0x03, 0x3a, 0x7d, 0xaf, 0xfd, 0xcb, 0xc4,
	0xfc, 0xd0, 0x39, 0x71, 0xce, 0xda, 0x17, 0xe7, 0xd0, 0x1c, 0xa8, 0xa7, 0x33, 0x91, 0x31, 0x90,
	0x5c, 0x43, 0x95, 0x0b, 0x41, 0xbe, 0x1a, 0x58, 0xa5, 0x7f, 0x33, 0xde, 0x53, 0x3f, 0x9f, 0xbd,
	0xd2, 0x7d, 0x2b, 0x7d, 0x67, 0x5d, 0xfa, 0xce, 0x67, 0xe9, 0x3b, 0xcf, 0x1b, 0xbf, 0xb5, 0xde,
	0xf8, 0xad, 0x8f, 0x8d, 0xdf, 0xf2, 0x60, 0xaa, 0x32, 0xd8, 0x7e, 0x89, 0xde, 0x41, 0xb3, 0xe3,
	0xb0, 0x2a, 0x3e, 0x74, 0xee, 0xaf, 0x1f, 0xa4, 0x99, 0x15, 0x49, 0x95, 0x45, 0x82, 0x2a, 0x2b,
	0xac, 0xb3, 0x08, 0x66, 0x91, 0x27, 0xf2, 0xff, 0x95, 0xbe, 0xb8, 0x3b, 0x41, 0x38, 0x19, 0x0d,
	0x5e, 0xdd, 0x6e, 0xd0, 0xa8, 0x10, 0x62, 0x85, 0x89, 0xad, 0x30, 0xc2, 0x0a, 0xb8, 0xda, 0xfb,
	0x1f, 0x38, 0x42, 0x38, 0xb2, 0x70, 0x84, 0x70, 0x84, 0x70, 0xe9, 0x5e, 0x6d, 0x0f, 0x47, 0xb7,
	0xc3, 0x5e, 0x28, 0x0c, 0xe3, 0xcc, 0xb0, 0x2f, 0x17, 0x1a, 0x22, 0xa5, 0x68, 0x52, 0x6a, 0x55,
	0x4a, 0xd1, 0xa5, 0x14, 0xe5, 0x64, 0xb7, 0x7e, 0xc6, 0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1a, 0x9d, 0x83, 0x9a, 0x3f, 0x02, 0x00, 0x00,
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
