// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/queries/order/query_request.proto

package order

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
	OrderID *base.OrderID `protobuf:"bytes,1,opt,name=order_i_d,json=orderID,proto3" json:"order_i_d,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b3c527835813dd, []int{0}
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

func (m *QueryRequest) GetOrderID() *base.OrderID {
	if m != nil {
		return m.OrderID
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.orders.queries.order.QueryRequest")
}

func init() {
	proto.RegisterFile("orders/queries/order/query_request.proto", fileDescriptor_79b3c527835813dd)
}

var fileDescriptor_79b3c527835813dd = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x4a, 0xf4, 0x40,
	0x18, 0x85, 0x77, 0x02, 0xdf, 0x27, 0x46, 0x1b, 0xb7, 0x51, 0x2c, 0x06, 0x11, 0x8b, 0x14, 0xcb,
	0x0c, 0xf8, 0xd3, 0x4c, 0x65, 0x82, 0x20, 0x16, 0x21, 0xc9, 0x96, 0x12, 0x08, 0x93, 0xcc, 0x8b,
	0x1b, 0xd8, 0x38, 0x6e, 0x66, 0x02, 0x7a, 0x17, 0x5e, 0x83, 0xa5, 0x57, 0x22, 0x56, 0x5b, 0x89,
	0xa5, 0x24, 0x9d, 0x57, 0x21, 0xc9, 0x4c, 0x11, 0xc1, 0x62, 0xbb, 0x9c, 0xf0, 0x9c, 0xe7, 0xbc,
	0x24, 0xae, 0x27, 0x6b, 0x01, 0xb5, 0xa2, 0xab, 0x06, 0xea, 0x12, 0x14, 0x1d, 0xe2, 0x90, 0x9e,
	0xb2, 0x1a, 0x56, 0x0d, 0x28, 0x4d, 0x1e, 0x6a, 0xa9, 0xe5, 0xd4, 0xe3, 0x4a, 0x81, 0xae, 0xf8,
	0xbd, 0x5e, 0x02, 0xa9, 0xa4, 0x68, 0x96, 0xa0, 0x88, 0x69, 0x13, 0xdb, 0x36, 0xf1, 0x70, 0xbf,
	0x14, 0x8a, 0xe6, 0x5c, 0x81, 0xb1, 0x65, 0xa5, 0x30, 0x8a, 0xe3, 0xd8, 0xdd, 0x4d, 0x7a, 0xf3,
	0xdc, 0x88, 0xa7, 0x97, 0xee, 0xb6, 0x25, 0x32, 0x71, 0x80, 0x8e, 0x90, 0xb7, 0x73, 0x7a, 0x42,
	0xc6, 0x33, 0xaa, 0x58, 0x40, 0xc5, 0x49, 0x29, 0x14, 0xe9, 0x7d, 0x24, 0xea, 0xe9, 0x9b, 0xab,
	0xf9, 0x96, 0x34, 0x0f, 0xc1, 0x87, 0xf3, 0xd6, 0x62, 0xb4, 0x6e, 0x31, 0xfa, 0x6a, 0x31, 0x7a,
	0xee, 0xf0, 0x64, 0xdd, 0xe1, 0xc9, 0x67, 0x87, 0x27, 0xee, 0xac, 0x90, 0x15, 0xd9, 0xf4, 0xe6,
	0x60, 0x6f, 0x7c, 0x58, 0xdc, 0x5f, 0x1b, 0xa3, 0xdb, 0x8b, 0xbb, 0x52, 0x2f, 0x9a, 0x9c, 0x14,
	0xb2, 0xa2, 0x7e, 0x6f, 0x0a, 0x07, 0x13, 0xb5, 0x26, 0xfa, 0x48, 0xff, 0xfa, 0x7a, 0x2f, 0xce,
	0x3f, 0x3f, 0x8c, 0x92, 0xe8, 0xd5, 0xf1, 0xfc, 0xd1, 0x7c, 0x68, 0xe7, 0x23, 0x33, 0x9f, 0xd8,
	0xf9, 0x21, 0xbe, 0xff, 0x42, 0x53, 0x8b, 0xa6, 0x06, 0x4d, 0x2d, 0x6a, 0x62, 0xeb, 0x9c, 0x6f,
	0x8a, 0xa6, 0xd7, 0x71, 0x10, 0x82, 0xe6, 0x82, 0x6b, 0xfe, 0xed, 0xcc, 0x46, 0x35, 0xc6, 0x6c,
	0x8f, 0x31, 0x53, 0x64, 0xcc, 0x36, 0xed, 0x8b, 0xfc, 0xff, 0xf0, 0xc7, 0xce, 0x7e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x50, 0xb2, 0xc1, 0x04, 0x20, 0x02, 0x00, 0x00,
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
	if m.OrderID != nil {
		{
			size, err := m.OrderID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.OrderID != nil {
		l = m.OrderID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
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
			if m.OrderID == nil {
				m.OrderID = &base.OrderID{}
			}
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
