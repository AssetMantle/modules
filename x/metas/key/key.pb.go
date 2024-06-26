// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/metas/key/key.proto

package key

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/ids/base"
	proto "github.com/cosmos/gogoproto/proto"
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
	DataID *base.DataID `protobuf:"bytes,1,opt,name=data_i_d,json=dataID,proto3" json:"data_i_d,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bb0e14a278ed474, []int{0}
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

func (m *Key) GetDataID() *base.DataID {
	if m != nil {
		return m.DataID
	}
	return nil
}

func init() {
	proto.RegisterType((*Key)(nil), "AssetMantle.modules.x.metas.key.Key")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/metas/key/key.proto", fileDescriptor_7bb0e14a278ed474)
}

var fileDescriptor_7bb0e14a278ed474 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x74, 0x2c, 0x2e, 0x4e,
	0x2d, 0xf1, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0x2d, 0xd6,
	0xaf, 0xd0, 0xcf, 0x4d, 0x2d, 0x49, 0x2c, 0xd6, 0xcf, 0x4e, 0xad, 0x04, 0x61, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x79, 0x24, 0xa5, 0x7a, 0x50, 0xa5, 0x7a, 0x15, 0x7a, 0x60, 0xa5, 0x7a,
	0xd9, 0xa9, 0x95, 0x52, 0x28, 0x66, 0x15, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0xea, 0x67, 0xa6, 0x14,
	0xeb, 0x27, 0x25, 0x16, 0xa7, 0xea, 0xa7, 0x24, 0x96, 0x24, 0xc6, 0x67, 0xa6, 0x40, 0xcc, 0x52,
	0x72, 0xe1, 0x62, 0xf6, 0x4e, 0xad, 0x14, 0xb2, 0xe5, 0xe2, 0x80, 0x88, 0xc7, 0xa7, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x29, 0xeb, 0x21, 0xdb, 0x02, 0x31, 0x44, 0x2f, 0x33, 0xa5, 0x58,
	0x0f, 0x64, 0x88, 0x9e, 0x4b, 0x62, 0x49, 0xa2, 0xa7, 0x4b, 0x10, 0x5b, 0x0a, 0x98, 0x76, 0x9a,
	0xca, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x5c, 0xca, 0xc9, 0xf9, 0xb9,
	0x7a, 0x04, 0x1c, 0xec, 0xc4, 0xe1, 0x9d, 0x5a, 0x19, 0x00, 0x72, 0x4f, 0x00, 0x63, 0x94, 0x56,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x81, 0x30, 0x59, 0xc4, 0xc4,
	0xea, 0xe8, 0x1b, 0xe1, 0xeb, 0xbd, 0x8a, 0x09, 0x25, 0x48, 0x7c, 0xa1, 0x36, 0x44, 0xe8, 0xf9,
	0x82, 0x6d, 0xf0, 0x4e, 0xad, 0x3c, 0x85, 0xa2, 0x22, 0x06, 0xaa, 0x22, 0x26, 0x22, 0x06, 0xac,
	0x22, 0xc6, 0x3b, 0xb5, 0xf2, 0x11, 0x93, 0x36, 0x01, 0x15, 0x31, 0xee, 0x01, 0x4e, 0x20, 0x0e,
	0xc8, 0xcf, 0xaf, 0x98, 0x94, 0x91, 0x54, 0x5b, 0x59, 0x41, 0x95, 0x5b, 0x59, 0x45, 0x58, 0x59,
	0x81, 0x35, 0x58, 0x59, 0x79, 0xa7, 0x56, 0x26, 0xb1, 0x81, 0x03, 0xd9, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xd7, 0x2c, 0x26, 0x35, 0xdd, 0x01, 0x00, 0x00,
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
	if m.DataID != nil {
		{
			size, err := m.DataID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.DataID != nil {
		l = m.DataID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field DataID", wireType)
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
			if m.DataID == nil {
				m.DataID = &base.DataID{}
			}
			if err := m.DataID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
