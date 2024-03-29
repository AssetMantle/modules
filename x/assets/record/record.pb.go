// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/record/record.proto

package record

import (
	fmt "fmt"
	key "github.com/AssetMantle/modules/x/assets/key"
	mappable "github.com/AssetMantle/modules/x/assets/mappable"
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

type Record struct {
	Key      *key.Key           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Mappable *mappable.Mappable `protobuf:"bytes,2,opt,name=mappable,proto3" json:"mappable,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_468fd57baf5368b0, []int{0}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Record)(nil), "assetmantle.modules.assets.record.Record")
}

func init() { proto.RegisterFile("assets/record/record.proto", fileDescriptor_468fd57baf5368b0) }

var fileDescriptor_468fd57baf5368b0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2c, 0x2e, 0x4e,
	0x2d, 0x29, 0xd6, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0x81, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x8a, 0x60, 0xb9, 0xdc, 0xc4, 0xbc, 0x92, 0x9c, 0x54, 0xbd, 0xdc, 0xfc, 0x94, 0xd2,
	0x9c, 0xd4, 0x62, 0x3d, 0x88, 0x7a, 0x3d, 0x88, 0x42, 0x29, 0x11, 0xa8, 0xf6, 0xec, 0xd4, 0x4a,
	0x10, 0x86, 0x68, 0x94, 0x92, 0x83, 0x8a, 0xe6, 0x26, 0x16, 0x14, 0x24, 0x26, 0xe5, 0xa4, 0xc2,
	0x19, 0x50, 0x79, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0x4d,
	0x60, 0xe4, 0x62, 0x0b, 0x02, 0x1b, 0x2b, 0x64, 0xca, 0xc5, 0x9c, 0x9d, 0x5a, 0x29, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xac, 0x87, 0xc7, 0x1d, 0x20, 0x4b, 0xbd, 0x53, 0x2b, 0x83, 0x40,
	0xea, 0x85, 0x3c, 0xb9, 0x38, 0x60, 0x36, 0x49, 0x30, 0x81, 0xf5, 0xea, 0xe2, 0xd3, 0x0b, 0x77,
	0x95, 0x2f, 0x94, 0x11, 0x04, 0xd7, 0x6e, 0xc5, 0xd2, 0xb1, 0x40, 0x9e, 0xc1, 0x69, 0x09, 0xd3,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0xa9, 0x26, 0xe7, 0xe7, 0xea, 0x11,
	0x0c, 0x20, 0x27, 0x6e, 0x88, 0x8f, 0x02, 0x40, 0x3e, 0x0c, 0x60, 0x8c, 0xd2, 0x4b, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x04, 0x29, 0xf4, 0x05, 0x6b, 0xd6, 0x87,
	0x6a, 0xd6, 0xaf, 0xd0, 0x47, 0x89, 0x8f, 0x45, 0x4c, 0x2c, 0x8e, 0xbe, 0x8e, 0x41, 0xab, 0x98,
	0x14, 0x1d, 0x91, 0x6c, 0xf2, 0x85, 0xda, 0xe4, 0x08, 0xb1, 0x09, 0x62, 0xc3, 0x29, 0x14, 0x35,
	0x31, 0x50, 0x35, 0x31, 0x10, 0x35, 0x31, 0x10, 0x35, 0x8f, 0x98, 0x74, 0x09, 0xaa, 0x89, 0x71,
	0x0f, 0x70, 0xf2, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x7c, 0xc5, 0xa4, 0x82, 0xa4, 0xde,
	0xca, 0x0a, 0xaa, 0xc1, 0xca, 0x0a, 0xa2, 0xc3, 0xca, 0x0a, 0xa2, 0x25, 0x89, 0x0d, 0x1c, 0x81,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x57, 0x54, 0x68, 0x4d, 0x02, 0x00, 0x00,
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mappable != nil {
		{
			size, err := m.Mappable.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	if m.Mappable != nil {
		l = m.Mappable.Size()
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}

func sovRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecord(x uint64) (n int) {
	return sovRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &key.Key{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mappable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Mappable == nil {
				m.Mappable = &mappable.Mappable{}
			}
			if err := m.Mappable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func skipRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
				return 0, ErrInvalidLengthRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecord = fmt.Errorf("proto: unexpected end of group")
)
