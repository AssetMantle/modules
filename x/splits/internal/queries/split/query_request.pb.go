// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/splits/internal/queries/split/query_request.proto

package split

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
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
	SplitID *base.SplitID `protobuf:"bytes,1,opt,name=split_i_d,json=splitID,proto3" json:"split_i_d,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_752004a4d72d6d4a, []int{0}
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

func (m *QueryRequest) GetSplitID() *base.SplitID {
	if m != nil {
		return m.SplitID
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.splits.queries.split.QueryRequest")
}

func init() {
	proto.RegisterFile("x/splits/internal/queries/split/query_request.proto", fileDescriptor_752004a4d72d6d4a)
}

var fileDescriptor_752004a4d72d6d4a = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x7b, 0x01, 0x15, 0xa3, 0x8b, 0x5d, 0x14, 0x87, 0x43, 0xc4, 0xa1, 0x43, 0xb9, 0x03,
	0xeb, 0x74, 0x83, 0xd8, 0x22, 0x88, 0x43, 0x20, 0x6d, 0x36, 0x09, 0x84, 0x4b, 0xee, 0xb0, 0x07,
	0xf9, 0xd3, 0xe6, 0xbd, 0x80, 0x7e, 0x0b, 0x3f, 0x83, 0xa3, 0x9f, 0x44, 0x9c, 0x3a, 0xba, 0x29,
	0xc9, 0xe6, 0xa7, 0x90, 0xe4, 0x6e, 0x88, 0x93, 0xdd, 0xf2, 0xbc, 0x3c, 0xcf, 0x2f, 0x3f, 0x12,
	0x77, 0xf2, 0x44, 0x61, 0x95, 0x2a, 0x0d, 0x54, 0xe5, 0x5a, 0x96, 0x39, 0x4f, 0xe9, 0xba, 0x92,
	0xa5, 0x92, 0x60, 0xee, 0x5d, 0x7a, 0x8e, 0x4a, 0xb9, 0xae, 0x24, 0x68, 0xb2, 0x2a, 0x0b, 0x5d,
	0x0c, 0x47, 0x1c, 0x40, 0xea, 0x8c, 0xe7, 0x3a, 0x95, 0x24, 0x2b, 0x44, 0x95, 0x4a, 0x20, 0x06,
	0x43, 0xec, 0xda, 0xc4, 0xd3, 0x63, 0x25, 0x80, 0xc6, 0x1c, 0xa4, 0xa1, 0x45, 0x4a, 0x18, 0xc4,
	0xb9, 0xef, 0x1e, 0xce, 0x5b, 0xf2, 0xc2, 0x80, 0x87, 0x37, 0xee, 0xbe, 0x6d, 0x44, 0xe2, 0x04,
	0x9d, 0xa1, 0xd1, 0xc1, 0xe5, 0x05, 0xe9, 0xbf, 0x06, 0x92, 0xa5, 0xcc, 0x38, 0x51, 0x02, 0x48,
	0xcb, 0x23, 0x41, 0xdb, 0xbe, 0xbf, 0x5d, 0xec, 0x81, 0x79, 0x98, 0x7d, 0x39, 0xef, 0x35, 0x46,
	0x9b, 0x1a, 0xa3, 0xef, 0x1a, 0xa3, 0x97, 0x06, 0x0f, 0x36, 0x0d, 0x1e, 0x7c, 0x36, 0x78, 0xe0,
	0x8e, 0x93, 0x22, 0x23, 0xdb, 0x3a, 0xcf, 0x8e, 0xfa, 0x62, 0x7e, 0x6b, 0xeb, 0xa3, 0x87, 0xeb,
	0x47, 0xa5, 0x97, 0x55, 0x4c, 0x92, 0x22, 0xa3, 0xd3, 0x96, 0xe4, 0x75, 0x24, 0x6a, 0x49, 0xf4,
	0x9f, 0xcf, 0xf8, 0xea, 0xec, 0x4c, 0xbd, 0x60, 0x1e, 0xbc, 0x39, 0xa3, 0x69, 0xcf, 0xc3, 0xb3,
	0x1e, 0x81, 0xf1, 0x98, 0x5b, 0x8f, 0x2e, 0x7e, 0xfc, 0xa9, 0x86, 0xb6, 0x1a, 0x9a, 0x6a, 0x68,
	0xab, 0x26, 0xd6, 0xce, 0xd5, 0xb6, 0xd5, 0xf0, 0xce, 0x9f, 0x79, 0x52, 0x73, 0xc1, 0x35, 0xff,
	0x71, 0xc6, 0xbd, 0x19, 0x63, 0x76, 0xc7, 0x98, 0x19, 0x32, 0x66, 0x97, 0xf6, 0x10, 0xef, 0x76,
	0xbf, 0x6e, 0xf2, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xde, 0x64, 0xe7, 0x6a, 0x34, 0x02, 0x00, 0x00,
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
	if m.SplitID != nil {
		{
			size, err := m.SplitID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.SplitID != nil {
		l = m.SplitID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field SplitID", wireType)
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
			if m.SplitID == nil {
				m.SplitID = &base.SplitID{}
			}
			if err := m.SplitID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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