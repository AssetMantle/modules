// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/splits/internal/queries/ownable/queryResponse.proto

package ownable

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type QueryResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_633fd91d00bdd449, []int{0}
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
	proto.RegisterType((*QueryResponse)(nil), "splits.queries.ownable.QueryResponse")
}

func init() {
	proto.RegisterFile("modules/splits/internal/queries/ownable/queryResponse.proto", fileDescriptor_633fd91d00bdd449)
}

var fileDescriptor_633fd91d00bdd449 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xce, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x2f, 0x2e, 0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0xd6, 0xcf, 0x2f, 0xcf, 0x4b,
	0x4c, 0xca, 0x49, 0x05, 0xf3, 0x2b, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xf5, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0x20, 0x9a, 0xf4, 0xa0, 0x6a, 0xf5, 0xa0, 0x6a, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4a, 0xf4, 0x41, 0x2c, 0x88, 0x6a, 0xa5, 0x68, 0x2e, 0xde, 0x40,
	0x64, 0x43, 0x84, 0x24, 0xb8, 0xd8, 0x8b, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x82, 0x60, 0x5c, 0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xa2, 0xa2, 0xfc, 0x22, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x07, 0x24, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1,
	0x0c, 0x11, 0x05, 0x73, 0xac, 0x58, 0x3a, 0x16, 0xc8, 0x33, 0x38, 0x75, 0x31, 0x9d, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78,
	0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x03, 0x97, 0x54, 0x72, 0x7e, 0xae, 0x1e, 0x76, 0x97, 0x3a,
	0x09, 0xa1, 0xb8, 0x28, 0x00, 0xe4, 0xce, 0x00, 0xc6, 0x28, 0xb7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2,
	0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2, 0xd4, 0x12, 0xdf, 0xc4, 0xbc, 0x92, 0x9c,
	0x54, 0x7d, 0x58, 0x58, 0x11, 0x19, 0x66, 0x8b, 0x98, 0x98, 0x83, 0x03, 0xfd, 0x57, 0x31, 0x89,
	0x05, 0x43, 0x2c, 0x0f, 0x84, 0x5a, 0xee, 0x0f, 0x91, 0x3e, 0x05, 0x93, 0x88, 0x81, 0x4a, 0xc4,
	0x40, 0x25, 0x1e, 0x31, 0x29, 0x61, 0x97, 0x88, 0x71, 0x0f, 0x70, 0xf2, 0x4d, 0x2d, 0x49, 0x4c,
	0x49, 0x2c, 0x49, 0x7c, 0xc5, 0x24, 0x01, 0x51, 0x64, 0x65, 0x05, 0x55, 0x65, 0x65, 0x05, 0x55,
	0x96, 0xc4, 0x06, 0x0e, 0x70, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x4a, 0xb1, 0xf5,
	0xdd, 0x01, 0x00, 0x00,
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
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintQueryResponse(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintQueryResponse(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
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
	if m.Success {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovQueryResponse(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovQueryResponse(uint64(l))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
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