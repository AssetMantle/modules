// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/queries/identity/queryResponse.proto

package identity

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	mappable "github.com/AssetMantle/modules/x/identities/internal/mappable"
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

type QueryResponse struct {
	List []*mappable.Mappable `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8bd48bf09628c56, []int{0}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryResponse)(nil), "assetmantle.modules.identities.queries.identity.QueryResponse")
}

func init() {
	proto.RegisterFile("x/identities/internal/queries/identity/queryResponse.proto", fileDescriptor_a8bd48bf09628c56)
}

var fileDescriptor_a8bd48bf09628c56 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0xfb, 0x40,
	0x00, 0x87, 0x93, 0xb4, 0xff, 0xff, 0x10, 0x71, 0x29, 0x0e, 0xd2, 0xe1, 0x14, 0xa7, 0x2e, 0xde,
	0x81, 0xdd, 0x4e, 0x97, 0x16, 0x44, 0x3a, 0x04, 0xda, 0xba, 0x49, 0x96, 0x4b, 0x73, 0xc4, 0x83,
	0x24, 0x17, 0x73, 0x17, 0xf0, 0xde, 0x40, 0x9c, 0x7c, 0x04, 0x71, 0xf4, 0x49, 0xc4, 0xa9, 0xa3,
	0xa3, 0x24, 0x9b, 0x4f, 0x21, 0xb9, 0x5c, 0x9b, 0x16, 0x04, 0xed, 0x76, 0xfc, 0x92, 0xef, 0xe3,
	0x3b, 0xce, 0xc5, 0xf7, 0x88, 0x85, 0x34, 0x95, 0x4c, 0x32, 0x2a, 0x10, 0x4b, 0x25, 0xcd, 0x53,
	0x12, 0xa3, 0xbb, 0x82, 0xe6, 0x7a, 0x68, 0xbe, 0x29, 0x3d, 0xa8, 0x39, 0x15, 0x19, 0x4f, 0x05,
	0x85, 0x59, 0xce, 0x25, 0xef, 0x21, 0x22, 0x04, 0x95, 0x09, 0x49, 0x65, 0x4c, 0x61, 0xc2, 0xc3,
	0x22, 0xa6, 0x02, 0xb6, 0x36, 0x68, 0x24, 0xab, 0x49, 0xf5, 0x0f, 0x22, 0x1e, 0x71, 0xcd, 0xa2,
	0xfa, 0xd4, 0x68, 0xfa, 0xa7, 0x3f, 0x27, 0x24, 0x24, 0xcb, 0x48, 0x10, 0xd3, 0xf5, 0xa1, 0xf9,
	0xfd, 0xe4, 0xda, 0xdd, 0x9f, 0x6d, 0xc6, 0xf4, 0x2e, 0xdc, 0x6e, 0xcc, 0x84, 0x3c, 0xb4, 0x8f,
	0x3b, 0x83, 0xbd, 0xb3, 0x01, 0xfc, 0xa5, 0xca, 0x33, 0xba, 0xb9, 0xa6, 0x70, 0xf7, 0xe1, 0xf9,
	0xc8, 0x1a, 0x3f, 0x76, 0xde, 0x4a, 0x60, 0x2f, 0x4b, 0x60, 0x7f, 0x96, 0xc0, 0x7e, 0xaa, 0x80,
	0xb5, 0xac, 0x80, 0xf5, 0x51, 0x01, 0xcb, 0x1d, 0x2e, 0x78, 0x02, 0x77, 0xbc, 0xe9, 0xb8, 0xb7,
	0x95, 0x38, 0xad, 0xc3, 0xa7, 0xf6, 0xcd, 0x65, 0xc4, 0xe4, 0x6d, 0x11, 0xc0, 0x05, 0x4f, 0xd0,
	0xa8, 0x36, 0x7a, 0xda, 0x88, 0x8c, 0x11, 0xfd, 0xed, 0x2d, 0x5e, 0x9c, 0x7f, 0x23, 0x6f, 0x32,
	0x9b, 0xbc, 0x3a, 0x8d, 0xc4, 0x64, 0x79, 0x26, 0x6b, 0xd2, 0x66, 0xcd, 0x4c, 0x96, 0x99, 0xd4,
	0xfb, 0x16, 0xe1, 0x1b, 0xc2, 0x6f, 0x09, 0xdf, 0x10, 0xab, 0x49, 0x95, 0xce, 0xf9, 0x8e, 0x84,
	0x7f, 0x35, 0x1d, 0x7b, 0x54, 0x92, 0x90, 0x48, 0xf2, 0xe5, 0x0c, 0x37, 0x68, 0x8c, 0x0d, 0x8e,
	0x71, 0xcb, 0x63, 0x6c, 0x04, 0xeb, 0x51, 0x05, 0xff, 0xf5, 0x43, 0x0f, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x82, 0x2b, 0x2c, 0x14, 0x9c, 0x02, 0x00, 0x00,
}

func (m *QueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryResponse(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovQueryResponse(uint64(l))
		}
	}
	return n
}

func sovQueryResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryResponse(x uint64) (n int) {
	return sovQueryResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryResponse
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
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponse
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
				return ErrInvalidLengthQueryResponse
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &mappable.Mappable{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryResponse
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
func skipQueryResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryResponse
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
					return 0, ErrIntOverflowQueryResponse
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
					return 0, ErrIntOverflowQueryResponse
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
				return 0, ErrInvalidLengthQueryResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryResponse = fmt.Errorf("proto: unexpected end of group")
)
