// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/properties/base/property.proto

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

type Property struct {
	// Types that are valid to be assigned to Impl:
	//	*Property_MesaProperty
	//	*Property_MetaProperty
	Impl isProperty_Impl `protobuf_oneof:"impl"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_909d89e9aab7ad23, []int{0}
}
func (m *Property) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Property.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return m.Size()
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

type isProperty_Impl interface {
	isProperty_Impl()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Property_MesaProperty struct {
	MesaProperty *MesaProperty `protobuf:"bytes,1,opt,name=mesaProperty,proto3,oneof" json:"mesaProperty,omitempty"`
}
type Property_MetaProperty struct {
	MetaProperty *MetaProperty `protobuf:"bytes,2,opt,name=metaProperty,proto3,oneof" json:"metaProperty,omitempty"`
}

func (*Property_MesaProperty) isProperty_Impl() {}
func (*Property_MetaProperty) isProperty_Impl() {}

func (m *Property) GetImpl() isProperty_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (m *Property) GetMesaProperty() *MesaProperty {
	if x, ok := m.GetImpl().(*Property_MesaProperty); ok {
		return x.MesaProperty
	}
	return nil
}

func (m *Property) GetMetaProperty() *MetaProperty {
	if x, ok := m.GetImpl().(*Property_MetaProperty); ok {
		return x.MetaProperty
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Property) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Property_MesaProperty)(nil),
		(*Property_MetaProperty)(nil),
	}
}

func init() {
	proto.RegisterType((*Property)(nil), "base.Property")
}

func init() {
	proto.RegisterFile("schema/properties/base/property.proto", fileDescriptor_909d89e9aab7ad23)
}

var fileDescriptor_909d89e9aab7ad23 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x2d, 0xd6, 0x4f, 0x4a,
	0x2c, 0x4e, 0x85, 0xf1, 0x2b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0x40, 0x82, 0x52,
	0x9a, 0x38, 0x14, 0xe7, 0xa6, 0x16, 0x27, 0x06, 0xa0, 0x68, 0xc0, 0xa3, 0xb4, 0x04, 0x4d, 0xa9,
	0x52, 0x1b, 0x23, 0x17, 0x07, 0x4c, 0x48, 0xc8, 0x82, 0x8b, 0x07, 0xd9, 0x34, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x21, 0x3d, 0x90, 0x66, 0x3d, 0x5f, 0x24, 0x19, 0x0f, 0x86, 0x20, 0x14,
	0x95, 0x10, 0x9d, 0x08, 0xc3, 0x25, 0x98, 0x50, 0x75, 0x96, 0xa0, 0xe9, 0x44, 0xf0, 0x9d, 0xd8,
	0xb8, 0x58, 0x32, 0x73, 0x0b, 0x72, 0x9c, 0x66, 0x30, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x03, 0x17, 0x47, 0x72, 0x7e, 0x2e, 0xd8, 0x24, 0x27, 0x5e, 0x98, 0xb6, 0x00, 0x90,
	0xe3, 0x03, 0x18, 0xa3, 0x4c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x1d, 0x8b, 0x8b, 0x53, 0x4b, 0x7c, 0x13, 0xf3, 0x4a, 0x72, 0x52, 0xf5, 0x73, 0xf3, 0x53, 0x4a,
	0x73, 0x52, 0x8b, 0xf5, 0xb1, 0x07, 0xc4, 0x22, 0x26, 0x66, 0xa7, 0x88, 0x88, 0x55, 0x4c, 0x2c,
	0x4e, 0x89, 0xc5, 0xa9, 0xa7, 0x20, 0xd4, 0x23, 0x26, 0x01, 0x10, 0x15, 0xe3, 0x1e, 0xe0, 0x04,
	0x72, 0x6f, 0x4a, 0x62, 0x49, 0xe2, 0x2b, 0x88, 0x4c, 0x12, 0x1b, 0x38, 0xa8, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xed, 0x0a, 0x36, 0xe0, 0xaf, 0x01, 0x00, 0x00,
}

func (m *Property) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Property) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Property) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Impl != nil {
		{
			size := m.Impl.Size()
			i -= size
			if _, err := m.Impl.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Property_MesaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Property_MesaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MesaProperty != nil {
		{
			size, err := m.MesaProperty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProperty(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Property_MetaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Property_MetaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MetaProperty != nil {
		{
			size, err := m.MetaProperty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProperty(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintProperty(dAtA []byte, offset int, v uint64) int {
	offset -= sovProperty(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Property) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Impl != nil {
		n += m.Impl.Size()
	}
	return n
}

func (m *Property_MesaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MesaProperty != nil {
		l = m.MesaProperty.Size()
		n += 1 + l + sovProperty(uint64(l))
	}
	return n
}
func (m *Property_MetaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MetaProperty != nil {
		l = m.MetaProperty.Size()
		n += 1 + l + sovProperty(uint64(l))
	}
	return n
}

func sovProperty(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProperty(x uint64) (n int) {
	return sovProperty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Property) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProperty
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
			return fmt.Errorf("proto: Property: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Property: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MesaProperty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProperty
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
				return ErrInvalidLengthProperty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProperty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MesaProperty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &Property_MesaProperty{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaProperty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProperty
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
				return ErrInvalidLengthProperty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProperty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MetaProperty{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &Property_MetaProperty{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProperty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProperty
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
func skipProperty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProperty
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
					return 0, ErrIntOverflowProperty
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
					return 0, ErrIntOverflowProperty
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
				return 0, ErrInvalidLengthProperty
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProperty
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProperty
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProperty        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProperty          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProperty = fmt.Errorf("proto: unexpected end of group")
)