// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/data/base/stringData.proto

package base

import (
	fmt "fmt"
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

type StringData struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StringData) Reset()         { *m = StringData{} }
func (m *StringData) String() string { return proto.CompactTextString(m) }
func (*StringData) ProtoMessage()    {}
func (*StringData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7649f0b4a236d0fc, []int{0}
}
func (m *StringData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringData.Merge(m, src)
}
func (m *StringData) XXX_Size() int {
	return m.Size()
}
func (m *StringData) XXX_DiscardUnknown() {
	xxx_messageInfo_StringData.DiscardUnknown(m)
}

var xxx_messageInfo_StringData proto.InternalMessageInfo

func (m *StringData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringData)(nil), "base.StringData")
}

func init() { proto.RegisterFile("schema/data/base/stringData.proto", fileDescriptor_7649f0b4a236d0fc) }

var fileDescriptor_7649f0b4a236d0fc = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x2e, 0x29,
	0xca, 0xcc, 0x4b, 0x77, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01,
	0x09, 0x2b, 0x29, 0x71, 0x71, 0x05, 0xc3, 0x65, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0xa7, 0x29, 0x8c, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0xc0, 0xc5, 0x91, 0x9c, 0x9f, 0xab, 0x07, 0x32, 0xc8, 0x89,
	0x1f, 0x61, 0x4c, 0x00, 0xc8, 0xfc, 0x00, 0xc6, 0x28, 0xfd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24,
	0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2, 0xd4, 0x12, 0xdf, 0xc4, 0xbc, 0x92, 0x9c, 0x54,
	0xfd, 0xdc, 0xfc, 0x94, 0xd2, 0x9c, 0xd4, 0x62, 0x7d, 0x74, 0x37, 0x2e, 0x62, 0x62, 0x76, 0x8a,
	0x88, 0x58, 0xc5, 0xc4, 0xe2, 0x94, 0x58, 0x9c, 0x7a, 0x0a, 0x42, 0x3d, 0x62, 0x12, 0x00, 0x51,
	0x31, 0xee, 0x01, 0x4e, 0xbe, 0xa9, 0x25, 0x89, 0x20, 0xb5, 0xaf, 0x20, 0x32, 0x49, 0x6c, 0x60,
	0x7f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x99, 0xb5, 0x0a, 0x4c, 0xec, 0x00, 0x00, 0x00,
}

func (m *StringData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintStringData(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStringData(dAtA []byte, offset int, v uint64) int {
	offset -= sovStringData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StringData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovStringData(uint64(l))
	}
	return n
}

func sovStringData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStringData(x uint64) (n int) {
	return sovStringData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStringData
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
			return fmt.Errorf("proto: StringData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStringData
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
				return ErrInvalidLengthStringData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStringData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStringData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStringData
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
func skipStringData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStringData
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
					return 0, ErrIntOverflowStringData
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
					return 0, ErrIntOverflowStringData
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
				return 0, ErrInvalidLengthStringData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStringData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStringData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStringData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStringData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStringData = fmt.Errorf("proto: unexpected end of group")
)