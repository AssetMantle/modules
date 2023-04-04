// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/splits/internal/genesis/genesis.proto

package genesis

import (
	fmt "fmt"
	mappable "github.com/AssetMantle/modules/modules/splits/internal/mappable"
	base "github.com/AssetMantle/schema/x/parameters/base"
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

type Genesis struct {
	Mappables     []*mappable.Mappable `protobuf:"bytes,1,rep,name=mappables,proto3" json:"mappables,omitempty"`
	ParameterList *base.ParameterList  `protobuf:"bytes,2,opt,name=parameterList,proto3" json:"parameterList,omitempty"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0819016d2585af3, []int{0}
}
func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return m.Size()
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Genesis)(nil), "splits.Genesis")
}

func init() {
	proto.RegisterFile("modules/splits/internal/genesis/genesis.proto", fileDescriptor_a0819016d2585af3)
}

var fileDescriptor_a0819016d2585af3 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcd, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x2f, 0x2e, 0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0x86, 0xd1, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x65, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x21, 0x7d,
	0x10, 0x0b, 0x22, 0x2b, 0xa5, 0x8f, 0xcb, 0xb0, 0xdc, 0xc4, 0x82, 0x82, 0xc4, 0xa4, 0x9c, 0x54,
	0x38, 0x03, 0xaa, 0x41, 0xab, 0x38, 0x39, 0x23, 0x35, 0x37, 0x51, 0xbf, 0x20, 0xb1, 0x28, 0x31,
	0x37, 0xb5, 0x24, 0xb5, 0xa8, 0x58, 0x3f, 0x29, 0xb1, 0x38, 0x15, 0xc1, 0xf7, 0xc9, 0x2c, 0x2e,
	0x81, 0xa8, 0x55, 0x6a, 0x60, 0xe4, 0x62, 0x77, 0x87, 0x38, 0x46, 0x48, 0x8f, 0x8b, 0x13, 0x66,
	0x52, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x80, 0x1e, 0xc4, 0x52, 0x3d, 0x5f, 0xa8,
	0x44, 0x10, 0x42, 0x89, 0x90, 0x3d, 0x17, 0x2f, 0x8a, 0x91, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0x92, 0x7a, 0x08, 0x8b, 0xf5, 0x02, 0x90, 0x15, 0x04, 0xa1, 0xaa, 0xb7, 0x62, 0xe9, 0x58,
	0x20, 0xcf, 0xe0, 0xb4, 0x8a, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xb8,
	0xb8, 0x92, 0xf3, 0x73, 0xa1, 0x2e, 0x70, 0xe2, 0x81, 0x3a, 0x33, 0x00, 0xe4, 0xee, 0x00, 0xc6,
	0x28, 0xbb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xc7, 0xe2, 0xe2,
	0xd4, 0x12, 0xdf, 0xc4, 0xbc, 0x12, 0x50, 0x60, 0x40, 0x43, 0x8b, 0x40, 0x14, 0x2c, 0x62, 0x62,
	0x0e, 0x8e, 0x88, 0x58, 0xc5, 0xc4, 0x16, 0x0c, 0x96, 0x3f, 0x05, 0x63, 0x3c, 0x62, 0x12, 0x82,
	0x30, 0x62, 0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x5f, 0xc1, 0x64,
	0x93, 0xd8, 0xc0, 0xc1, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x54, 0xef, 0x1f, 0xe2,
	0x01, 0x00, 0x00,
}

func (m *Genesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Genesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Genesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ParameterList != nil {
		{
			size, err := m.ParameterList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Mappables) > 0 {
		for iNdEx := len(m.Mappables) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Mappables[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Genesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Mappables) > 0 {
		for _, e := range m.Mappables {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ParameterList != nil {
		l = m.ParameterList.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Genesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Genesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Genesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mappables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mappables = append(m.Mappables, &mappable.Mappable{})
			if err := m.Mappables[len(m.Mappables)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParameterList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParameterList == nil {
				m.ParameterList = &base.ParameterList{}
			}
			if err := m.ParameterList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
