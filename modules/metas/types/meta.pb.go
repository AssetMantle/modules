// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistenceSDK/modules/metas/meta.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	base "github.com/persistenceOne/persistenceSDK/schema/types/base"
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

type MetaID struct {
	TypeID base.ID `protobuf:"bytes,1,opt,name=type_i_d,json=typeID,proto3" json:"type_i_d"`
	HashID base.ID `protobuf:"bytes,2,opt,name=hash_i_d,json=hashID,proto3" json:"hash_i_d"`
}

func (m *MetaID) Reset()      { *m = MetaID{} }
func (*MetaID) ProtoMessage() {}
func (*MetaID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eafa3781de421bd, []int{0}
}
func (m *MetaID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetaID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetaID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetaID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaID.Merge(m, src)
}
func (m *MetaID) XXX_Size() int {
	return m.Size()
}
func (m *MetaID) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaID.DiscardUnknown(m)
}

var xxx_messageInfo_MetaID proto.InternalMessageInfo

func (m *MetaID) GetTypeID() base.ID {
	if m != nil {
		return m.TypeID
	}
	return base.ID{}
}

func (m *MetaID) GetHashID() base.ID {
	if m != nil {
		return m.HashID
	}
	return base.ID{}
}

type Meta struct {
	ID   MetaID    `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3" json:"i_d"`
	Data types.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eafa3781de421bd, []int{1}
}
func (m *Meta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return m.Size()
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MetaID)(nil), "persistenceSDK.modules.metas.MetaID")
	proto.RegisterType((*Meta)(nil), "persistenceSDK.modules.metas.Meta")
}

func init() {
	proto.RegisterFile("persistenceSDK/modules/metas/meta.proto", fileDescriptor_1eafa3781de421bd)
}

var fileDescriptor_1eafa3781de421bd = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x0d, 0x76, 0xf1, 0xd6, 0xcf, 0xcd, 0x4f, 0x29, 0xcd,
	0x49, 0x2d, 0xd6, 0xcf, 0x4d, 0x2d, 0x49, 0x84, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x32, 0xa8, 0x0a, 0xf5, 0xa0, 0x0a, 0xf5, 0xc0, 0x0a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1,
	0x0a, 0xf5, 0x41, 0x2c, 0x88, 0x1e, 0x29, 0x4d, 0x34, 0xc3, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13,
	0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x93, 0x12, 0x8b, 0x53, 0xf5, 0x33, 0x53, 0xa0, 0x4a,
	0x25, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0xc4, 0xbc,
	0x4a, 0x98, 0x54, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x71, 0x3c, 0xc4, 0x78, 0x08, 0x07, 0x22, 0xa5,
	0x34, 0x8b, 0x91, 0x8b, 0xcd, 0x37, 0xb5, 0x24, 0xd1, 0xd3, 0x45, 0xc8, 0x85, 0x8b, 0x03, 0x64,
	0x6e, 0x7c, 0x66, 0x7c, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x8a, 0x1e, 0x9a, 0x93,
	0x21, 0xd6, 0xeb, 0x81, 0xad, 0xd7, 0x03, 0x59, 0xaf, 0xe7, 0xe9, 0xe2, 0xc4, 0x72, 0xe2, 0x9e,
	0x3c, 0x43, 0x10, 0x1b, 0x48, 0x10, 0x62, 0x4a, 0x46, 0x62, 0x71, 0x06, 0xd8, 0x14, 0x26, 0xd2,
	0x4d, 0x01, 0xe9, 0xf5, 0x74, 0xb1, 0x62, 0x99, 0xb1, 0x40, 0x9e, 0x41, 0x69, 0x37, 0x23, 0x17,
	0x0b, 0xc8, 0x71, 0x42, 0xd6, 0x5c, 0xcc, 0x78, 0x5c, 0x85, 0x12, 0x90, 0x7a, 0x10, 0xdf, 0x40,
	0xcd, 0x63, 0xca, 0x74, 0x11, 0xd2, 0xe3, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x84, 0xba, 0x46, 0x44,
	0x0f, 0x12, 0x4e, 0x7a, 0xb0, 0x70, 0xd2, 0x73, 0xcc, 0xab, 0x84, 0xaa, 0x06, 0xab, 0xb3, 0x72,
	0xeb, 0x58, 0x20, 0xcf, 0x70, 0x69, 0x8b, 0xae, 0x5d, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x3e, 0x92, 0x8d, 0xfe, 0x79, 0xa9, 0xfa, 0xd8, 0x63, 0x25, 0x37, 0xb1, 0xa0,
	0x20, 0x31, 0x09, 0xe4, 0x12, 0x90, 0x1b, 0x9c, 0xc2, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1,
	0x58, 0x8e, 0x21, 0xca, 0x86, 0x68, 0x93, 0x51, 0x13, 0x13, 0x38, 0xc4, 0x92, 0xd8, 0xc0, 0x2e,
	0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x09, 0xc1, 0x17, 0xcd, 0x79, 0x02, 0x00, 0x00,
}

func (m *MetaID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetaID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetaID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HashID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMeta(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.TypeID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMeta(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Meta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Meta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Meta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMeta(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMeta(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMeta(dAtA []byte, offset int, v uint64) int {
	offset -= sovMeta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetaID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TypeID.Size()
	n += 1 + l + sovMeta(uint64(l))
	l = m.HashID.Size()
	n += 1 + l + sovMeta(uint64(l))
	return n
}

func (m *Meta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovMeta(uint64(l))
	l = m.Data.Size()
	n += 1 + l + sovMeta(uint64(l))
	return n
}

func sovMeta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMeta(x uint64) (n int) {
	return sovMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetaID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: MetaID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetaID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TypeID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HashID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMeta
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
func (m *Meta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: Meta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Meta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
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
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
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
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMeta
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
func skipMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
				return 0, ErrInvalidLengthMeta
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMeta
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMeta
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMeta        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeta          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMeta = fmt.Errorf("proto: unexpected end of group")
)
