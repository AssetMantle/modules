// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/lists/base/propertyList.v1.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/properties/base"
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

type PropertyList struct {
	List []*base.Property `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (m *PropertyList) Reset()         { *m = PropertyList{} }
func (m *PropertyList) String() string { return proto.CompactTextString(m) }
func (*PropertyList) ProtoMessage()    {}
func (*PropertyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18fec15f46dd73c, []int{0}
}
func (m *PropertyList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PropertyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PropertyList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PropertyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropertyList.Merge(m, src)
}
func (m *PropertyList) XXX_Size() int {
	return m.Size()
}
func (m *PropertyList) XXX_DiscardUnknown() {
	xxx_messageInfo_PropertyList.DiscardUnknown(m)
}

var xxx_messageInfo_PropertyList proto.InternalMessageInfo

func (m *PropertyList) GetList() []*base.Property {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*PropertyList)(nil), "lists.PropertyList")
}

func init() {
	proto.RegisterFile("schema/lists/base/propertyList.v1.proto", fileDescriptor_a18fec15f46dd73c)
}

var fileDescriptor_a18fec15f46dd73c = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0xcf, 0xc9, 0x2c, 0x2e, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x28,
	0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xa9, 0xf4, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x33, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xab, 0x90, 0xd2, 0x80, 0xaa, 0x87, 0x2a, 0xca, 0x4c, 0x45,
	0xd3, 0x04, 0xd7, 0xa0, 0x64, 0xc1, 0xc5, 0x13, 0x80, 0x64, 0x92, 0x90, 0x06, 0x17, 0x0b, 0xc8,
	0x08, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x11, 0x3d, 0x84, 0x09, 0x7a, 0x30, 0x75, 0x41,
	0x60, 0x15, 0x4e, 0xf3, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x8b,
	0x33, 0x39, 0x3f, 0x57, 0x0f, 0xec, 0x12, 0x27, 0x61, 0x64, 0xd3, 0xc3, 0x0c, 0x03, 0x40, 0x96,
	0x06, 0x30, 0x46, 0x19, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b,
	0x16, 0x17, 0xa7, 0x96, 0xf8, 0x26, 0xe6, 0x95, 0xe4, 0xa4, 0xea, 0xe7, 0xe6, 0xa7, 0x94, 0xe6,
	0xa4, 0x16, 0xeb, 0x63, 0xf8, 0x77, 0x11, 0x13, 0xb3, 0x4f, 0x44, 0xc4, 0x2a, 0x26, 0x56, 0x90,
	0x39, 0xc5, 0xa7, 0xa0, 0xf4, 0x23, 0x26, 0x41, 0x30, 0x1d, 0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a,
	0x92, 0x98, 0x92, 0x58, 0x92, 0xf8, 0x0a, 0x2a, 0x97, 0xc4, 0x06, 0xf6, 0xa1, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0xf4, 0x38, 0xfb, 0x3d, 0x01, 0x00, 0x00,
}

func (m *PropertyList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PropertyList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PropertyList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintPropertyListV1(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPropertyListV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovPropertyListV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PropertyList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovPropertyListV1(uint64(l))
		}
	}
	return n
}

func sovPropertyListV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPropertyListV1(x uint64) (n int) {
	return sovPropertyListV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PropertyList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPropertyListV1
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
			return fmt.Errorf("proto: PropertyList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PropertyList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPropertyListV1
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
				return ErrInvalidLengthPropertyListV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPropertyListV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &base.Property{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPropertyListV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPropertyListV1
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
func skipPropertyListV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPropertyListV1
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
					return 0, ErrIntOverflowPropertyListV1
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
					return 0, ErrIntOverflowPropertyListV1
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
				return 0, ErrInvalidLengthPropertyListV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPropertyListV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPropertyListV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPropertyListV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPropertyListV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPropertyListV1 = fmt.Errorf("proto: unexpected end of group")
)