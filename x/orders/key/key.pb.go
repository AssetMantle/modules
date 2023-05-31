// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/key/key.proto

package key

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

type Key struct {
	OrderID *base.OrderID `protobuf:"bytes,1,opt,name=order_i_d,json=orderID,proto3" json:"order_i_d,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2e237a71fe29dc, []int{0}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return m.Size()
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetOrderID() *base.OrderID {
	if m != nil {
		return m.OrderID
	}
	return nil
}

func init() {
	proto.RegisterType((*Key)(nil), "assetmantle.modules.orders.key.Key")
}

func init() { proto.RegisterFile("orders/key/key.proto", fileDescriptor_5b2e237a71fe29dc) }

var fileDescriptor_5b2e237a71fe29dc = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x2f, 0x4a, 0x49,
	0x2d, 0x2a, 0xd6, 0xcf, 0x4e, 0xad, 0x04, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xb9,
	0xc4, 0xe2, 0xe2, 0xd4, 0x92, 0xdc, 0xc4, 0xbc, 0x92, 0x9c, 0x54, 0xbd, 0xdc, 0xfc, 0x94, 0xd2,
	0x9c, 0xd4, 0x62, 0x3d, 0x88, 0x4a, 0xbd, 0xec, 0xd4, 0x4a, 0x29, 0xf1, 0xcc, 0x94, 0x62, 0xfd,
	0xa4, 0xc4, 0xe2, 0x54, 0x7d, 0xb0, 0x60, 0x7c, 0x66, 0x0a, 0x44, 0xa3, 0x92, 0x3b, 0x17, 0xb3,
	0x77, 0x6a, 0xa5, 0x90, 0x03, 0x17, 0x27, 0x54, 0x22, 0x3e, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x45, 0x0f, 0xd9, 0xcc, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0xbd, 0xcc, 0x94, 0x62,
	0x3d, 0x90, 0x31, 0x7a, 0xfe, 0x20, 0xd5, 0x9e, 0x2e, 0x41, 0xec, 0xf9, 0x10, 0x86, 0x53, 0x3f,
	0xd3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0x29, 0x25, 0xe7, 0xe7, 0xea,
	0xe1, 0x77, 0xa0, 0x13, 0x87, 0x77, 0x6a, 0x65, 0x00, 0xc8, 0x45, 0x01, 0x8c, 0x51, 0xda, 0xe9,
	0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x8e, 0x20, 0x6d, 0xbe, 0x60, 0x6d,
	0xfa, 0x50, 0x6d, 0xfa, 0x15, 0xfa, 0x88, 0x30, 0x58, 0xc4, 0xc4, 0xe2, 0xe8, 0xeb, 0xef, 0xbd,
	0x8a, 0x49, 0xce, 0x11, 0xc9, 0x02, 0x5f, 0xa8, 0x05, 0xfe, 0x10, 0x0b, 0xbc, 0x53, 0x2b, 0x4f,
	0xa1, 0x28, 0x88, 0x81, 0x2a, 0x88, 0x81, 0x28, 0x88, 0xf1, 0x4e, 0xad, 0x7c, 0xc4, 0xa4, 0x85,
	0x5f, 0x41, 0x8c, 0x7b, 0x80, 0x93, 0x6f, 0x6a, 0x49, 0x62, 0x4a, 0x62, 0x49, 0xe2, 0x2b, 0x26,
	0x45, 0x24, 0xc5, 0x56, 0x56, 0x50, 0xd5, 0x56, 0x56, 0x10, 0xe5, 0x56, 0x56, 0xde, 0xa9, 0x95,
	0x49, 0x6c, 0xe0, 0x10, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x0e, 0xa8, 0x62, 0xb2,
	0x01, 0x00, 0x00,
}

func (m *Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintKey(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderID != nil {
		l = m.OrderID.Size()
		n += 1 + l + sovKey(uint64(l))
	}
	return n
}

func sovKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKey(x uint64) (n int) {
	return sovKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKey
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKey
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
			skippy, err := skipKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKey
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
func skipKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKey
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
					return 0, ErrIntOverflowKey
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
					return 0, ErrIntOverflowKey
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
				return 0, ErrInvalidLengthKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKey = fmt.Errorf("proto: unexpected end of group")
)
