// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/parameter.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	_ "github.com/regen-network/cosmos-proto"
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

type Parameter struct {
	ID   github_com_persistenceOne_persistenceSDK_schema_types.ID   `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"i_d"`
	Data github_com_persistenceOne_persistenceSDK_schema_types.Data `protobuf:"bytes,2,opt,name=data,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Data" json:"data"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c55cbae88e4b524a, []int{0}
}
func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return m.Size()
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Parameter)(nil), "persistence_sdk.schema.types.base.Parameter")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/parameter.proto", fileDescriptor_c55cbae88e4b524a)
}

var fileDescriptor_c55cbae88e4b524a = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2c, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x48,
	0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x44,
	0xd3, 0xa2, 0x07, 0xd1, 0xa2, 0x07, 0xd6, 0xa2, 0x07, 0xd2, 0x22, 0x25, 0x92, 0x9e, 0x9f, 0x9e,
	0x0f, 0x56, 0xad, 0x0f, 0x62, 0x41, 0x34, 0x4a, 0x49, 0x26, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xc7,
	0x43, 0x24, 0x20, 0x1c, 0xa8, 0x94, 0x16, 0x61, 0x67, 0x64, 0xa6, 0x40, 0xd4, 0x2a, 0xcd, 0x65,
	0xe2, 0xe2, 0x0c, 0x80, 0xb9, 0x49, 0xa8, 0x82, 0x8b, 0x39, 0x33, 0x3e, 0x45, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x55, 0x8f, 0xa0, 0xdb, 0xf4, 0x3c, 0x5d, 0x9c, 0x1c, 0x4e, 0xdc, 0x93,
	0x67, 0xb8, 0x75, 0x4f, 0xde, 0x22, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x1f, 0x49, 0xa3, 0x7f, 0x5e, 0x2a, 0x32, 0x37, 0xd8, 0xc5, 0x1b, 0xc5, 0x39, 0x7a, 0x9e, 0x2e,
	0x41, 0x4c, 0x99, 0x2e, 0x42, 0x53, 0x19, 0xb9, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x98, 0x14,
	0x18, 0x35, 0x38, 0x9d, 0x1a, 0x18, 0xa1, 0xa6, 0x5a, 0x91, 0x67, 0xaa, 0x4b, 0x62, 0x49, 0xe2,
	0xa9, 0x2d, 0xba, 0x14, 0xe8, 0x0e, 0x02, 0x3b, 0xc7, 0x29, 0xe4, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xc8, 0x33, 0x1b, 0x1c, 0xfc, 0x49, 0x6c, 0xe0, 0xc0, 0x37, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x79, 0x2f, 0xf0, 0x50, 0x31, 0x02, 0x00, 0x00,
}

func (m *Parameter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Parameter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Parameter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Data.Size()
		i -= size
		if _, err := m.Data.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParameter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParameter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParameter(dAtA []byte, offset int, v uint64) int {
	offset -= sovParameter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Parameter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovParameter(uint64(l))
	l = m.Data.Size()
	n += 1 + l + sovParameter(uint64(l))
	return n
}

func sovParameter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParameter(x uint64) (n int) {
	return sovParameter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Parameter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParameter
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
			return fmt.Errorf("proto: Parameter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Parameter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParameter
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
				return ErrInvalidLengthParameter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParameter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParameter
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
				return ErrInvalidLengthParameter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParameter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParameter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParameter
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
func skipParameter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParameter
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
					return 0, ErrIntOverflowParameter
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
					return 0, ErrIntOverflowParameter
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
				return 0, ErrInvalidLengthParameter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParameter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParameter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParameter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParameter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParameter = fmt.Errorf("proto: unexpected end of group")
)
