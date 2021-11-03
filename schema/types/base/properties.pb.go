// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistenceSDK/schema/types/base/properties.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Properties struct {
	PropertyList []Property `protobuf:"bytes,1,rep,name=propertyList,proto3" json:"propertyList"`
}

func (m *Properties) Reset()         { *m = Properties{} }
func (m *Properties) String() string { return proto.CompactTextString(m) }
func (*Properties) ProtoMessage()    {}
func (*Properties) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb548bbcf36d52c, []int{0}
}
func (m *Properties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Properties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Properties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Properties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Properties.Merge(m, src)
}
func (m *Properties) XXX_Size() int {
	return m.Size()
}
func (m *Properties) XXX_DiscardUnknown() {
	xxx_messageInfo_Properties.DiscardUnknown(m)
}

var xxx_messageInfo_Properties proto.InternalMessageInfo

func (m *Properties) GetPropertyList() []Property {
	if m != nil {
		return m.PropertyList
	}
	return nil
}

func init() {
	proto.RegisterType((*Properties)(nil), "persistenceSDK.schema.types.base.Properties")
}

func init() {
	proto.RegisterFile("persistenceSDK/schema/types/base/properties.proto", fileDescriptor_ccb548bbcf36d52c)
}

var fileDescriptor_ccb548bbcf36d52c = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2c, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x0d, 0x76, 0xf1, 0xd6, 0x2f, 0x4e, 0xce, 0x48, 0xcd,
	0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x28, 0xca,
	0x2f, 0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x40,
	0xd5, 0xa2, 0x07, 0xd1, 0xa2, 0x07, 0xd6, 0xa2, 0x07, 0xd2, 0x22, 0x25, 0x92, 0x9e, 0x9f, 0x9e,
	0x0f, 0x56, 0xac, 0x0f, 0x62, 0x41, 0xf4, 0x49, 0xe9, 0x13, 0x6b, 0x55, 0x25, 0x44, 0x83, 0x52,
	0x12, 0x17, 0x57, 0x00, 0xdc, 0x72, 0xa1, 0x10, 0x2e, 0x1e, 0x98, 0xbc, 0x4f, 0x66, 0x71, 0x89,
	0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x96, 0x1e, 0x21, 0xd7, 0xe8, 0x41, 0xcd, 0xa8, 0x74,
	0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xc5, 0x14, 0xa7, 0x90, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0x45, 0x76, 0xb9, 0x7f, 0x5e, 0x2a, 0x41, 0x8f, 0x24, 0xb1, 0x81, 0x3d, 0x60, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x00, 0xb9, 0xe4, 0x00, 0x5e, 0x01, 0x00, 0x00,
}

func (m *Properties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Properties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Properties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PropertyList) > 0 {
		for iNdEx := len(m.PropertyList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PropertyList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProperties(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProperties(dAtA []byte, offset int, v uint64) int {
	offset -= sovProperties(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Properties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PropertyList) > 0 {
		for _, e := range m.PropertyList {
			l = e.Size()
			n += 1 + l + sovProperties(uint64(l))
		}
	}
	return n
}

func sovProperties(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProperties(x uint64) (n int) {
	return sovProperties(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Properties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProperties
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
			return fmt.Errorf("proto: Properties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Properties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertyList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProperties
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
				return ErrInvalidLengthProperties
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PropertyList = append(m.PropertyList, Property{})
			if err := m.PropertyList[len(m.PropertyList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProperties
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
func skipProperties(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProperties
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
					return 0, ErrIntOverflowProperties
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
					return 0, ErrIntOverflowProperties
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
				return 0, ErrInvalidLengthProperties
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProperties
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProperties
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProperties        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProperties          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProperties = fmt.Errorf("proto: unexpected end of group")
)
