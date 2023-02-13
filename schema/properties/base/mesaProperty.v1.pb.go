// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/properties/base/mesaProperty.v1.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/ids/base"
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

type MesaProperty struct {
	ID     *base.PropertyID `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3" json:"i_d,omitempty"`
	DataID *base.DataID     `protobuf:"bytes,2,opt,name=data_i_d,json=dataID,proto3" json:"data_i_d,omitempty"`
}

func (m *MesaProperty) Reset()         { *m = MesaProperty{} }
func (m *MesaProperty) String() string { return proto.CompactTextString(m) }
func (*MesaProperty) ProtoMessage()    {}
func (*MesaProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_68ec296767b2fdde, []int{0}
}
func (m *MesaProperty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MesaProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MesaProperty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MesaProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MesaProperty.Merge(m, src)
}
func (m *MesaProperty) XXX_Size() int {
	return m.Size()
}
func (m *MesaProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_MesaProperty.DiscardUnknown(m)
}

var xxx_messageInfo_MesaProperty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MesaProperty)(nil), "properties.MesaProperty")
}

func init() {
	proto.RegisterFile("schema/properties/base/mesaProperty.v1.proto", fileDescriptor_68ec296767b2fdde)
}

var fileDescriptor_68ec296767b2fdde = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x2d, 0xd6, 0x4f, 0x4a,
	0x2c, 0x4e, 0xd5, 0xcf, 0x4d, 0x2d, 0x4e, 0x0c, 0x80, 0x88, 0x55, 0xea, 0x95, 0x19, 0xea, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x21, 0x94, 0x49, 0x29, 0x43, 0x75, 0x66, 0xa6, 0x40, 0xb5,
	0x40, 0xe5, 0x2a, 0x3d, 0x5d, 0xe0, 0x1a, 0xa4, 0xe4, 0xd1, 0x15, 0xa5, 0x24, 0x96, 0x24, 0x22,
	0x2b, 0x10, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0x52, 0x2c, 0x17,
	0x8f, 0x2f, 0x92, 0x03, 0x84, 0x14, 0xb8, 0x98, 0x33, 0xe3, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35,
	0xb8, 0x8d, 0xf8, 0xf5, 0x32, 0x53, 0x8a, 0xf5, 0x02, 0xe0, 0xb6, 0x05, 0x31, 0x65, 0xba, 0x08,
	0xa9, 0x72, 0x71, 0x80, 0x8c, 0x8e, 0x07, 0x29, 0x63, 0x02, 0x2b, 0xe3, 0x06, 0x2b, 0x73, 0x01,
	0xdb, 0x17, 0xc4, 0x06, 0xb1, 0xd7, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0xd3, 0x1e, 0xc6, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0xe2, 0x4b, 0xce, 0xcf, 0xd5, 0x43, 0xf8,
	0xd2, 0x49, 0x18, 0xd9, 0x1d, 0x61, 0x86, 0x01, 0x20, 0xe7, 0x05, 0x30, 0x46, 0x99, 0xa6, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0x16, 0x17, 0xa7, 0x96, 0xf8, 0x26,
	0xe6, 0x95, 0xe4, 0xa4, 0xea, 0xe7, 0xe6, 0xa7, 0x94, 0xe6, 0xa4, 0x16, 0xeb, 0x63, 0x0f, 0xd5,
	0x45, 0x4c, 0xcc, 0x01, 0x11, 0x11, 0xab, 0x98, 0xb8, 0x02, 0xe0, 0xe2, 0xa7, 0x90, 0x39, 0x8f,
	0x98, 0xc4, 0x10, 0x9c, 0x18, 0xf7, 0x00, 0x27, 0xdf, 0xd4, 0x92, 0x44, 0x90, 0xcb, 0x5f, 0x21,
	0xab, 0x4a, 0x62, 0x03, 0x07, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x35, 0xc1, 0x4f, 0xb9,
	0xbc, 0x01, 0x00, 0x00,
}

func (m *MesaProperty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MesaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MesaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DataID != nil {
		{
			size, err := m.DataID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMesaPropertyV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ID != nil {
		{
			size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMesaPropertyV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMesaPropertyV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovMesaPropertyV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MesaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != nil {
		l = m.ID.Size()
		n += 1 + l + sovMesaPropertyV1(uint64(l))
	}
	if m.DataID != nil {
		l = m.DataID.Size()
		n += 1 + l + sovMesaPropertyV1(uint64(l))
	}
	return n
}

func sovMesaPropertyV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMesaPropertyV1(x uint64) (n int) {
	return sovMesaPropertyV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MesaProperty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesaPropertyV1
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
			return fmt.Errorf("proto: MesaProperty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MesaProperty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesaPropertyV1
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
				return ErrInvalidLengthMesaPropertyV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesaPropertyV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ID == nil {
				m.ID = &base.PropertyID{}
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesaPropertyV1
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
				return ErrInvalidLengthMesaPropertyV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesaPropertyV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DataID == nil {
				m.DataID = &base.DataID{}
			}
			if err := m.DataID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesaPropertyV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMesaPropertyV1
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
func skipMesaPropertyV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMesaPropertyV1
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
					return 0, ErrIntOverflowMesaPropertyV1
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
					return 0, ErrIntOverflowMesaPropertyV1
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
				return 0, ErrInvalidLengthMesaPropertyV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMesaPropertyV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMesaPropertyV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMesaPropertyV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMesaPropertyV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMesaPropertyV1 = fmt.Errorf("proto: unexpected end of group")
)
