// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/classifications/internal/mappable/mappable.v1.proto

package mappable

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/documents/base"
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

type Mappable struct {
	Classification *base.Document `protobuf:"bytes,1,opt,name=classification,proto3" json:"classification,omitempty"`
}

func (m *Mappable) Reset()         { *m = Mappable{} }
func (m *Mappable) String() string { return proto.CompactTextString(m) }
func (*Mappable) ProtoMessage()    {}
func (*Mappable) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d85bfcae5a28ba4, []int{0}
}
func (m *Mappable) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mappable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mappable.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mappable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mappable.Merge(m, src)
}
func (m *Mappable) XXX_Size() int {
	return m.Size()
}
func (m *Mappable) XXX_DiscardUnknown() {
	xxx_messageInfo_Mappable.DiscardUnknown(m)
}

var xxx_messageInfo_Mappable proto.InternalMessageInfo

func (m *Mappable) GetClassification() *base.Document {
	if m != nil {
		return m.Classification
	}
	return nil
}

func init() {
	proto.RegisterType((*Mappable)(nil), "classifications.Mappable")
}

func init() {
	proto.RegisterFile("modules/classifications/internal/mappable/mappable.v1.proto", fileDescriptor_5d85bfcae5a28ba4)
}

var fileDescriptor_5d85bfcae5a28ba4 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xce, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0xce, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9,
	0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x4d, 0x2c,
	0x28, 0x48, 0x4c, 0xca, 0x49, 0x85, 0x33, 0xf4, 0xca, 0x0c, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xf8, 0xd1, 0x34, 0x49, 0xa9, 0x17, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0xea, 0xa7, 0xe4, 0x27,
	0x97, 0xe6, 0xa6, 0xe6, 0x95, 0x14, 0xeb, 0x27, 0x25, 0x16, 0xa7, 0xc2, 0xb9, 0x70, 0x9d, 0x4a,
	0xee, 0x5c, 0x1c, 0xbe, 0x50, 0xe3, 0x84, 0xac, 0xb9, 0xf8, 0x50, 0xcd, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x36, 0x12, 0xd6, 0x83, 0x1b, 0xa3, 0xe7, 0x02, 0x65, 0x05, 0xa1, 0x29, 0x75, 0x7a,
	0xc2, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x5c, 0xc2, 0xc9, 0xf9, 0xb9,
	0x7a, 0x68, 0x2e, 0x74, 0xe2, 0x87, 0x59, 0x1b, 0x66, 0x18, 0x00, 0x72, 0x49, 0x00, 0x63, 0x94,
	0x47, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x63, 0x71, 0x71, 0x6a,
	0x89, 0x6f, 0x62, 0x5e, 0x09, 0xc8, 0xbf, 0xd0, 0x90, 0x21, 0x3a, 0x84, 0x16, 0x31, 0x31, 0x3b,
	0x47, 0x44, 0xac, 0x62, 0xe2, 0x77, 0x46, 0x55, 0x79, 0x0a, 0x43, 0xe4, 0x11, 0x93, 0x34, 0x9a,
	0x48, 0x8c, 0x7b, 0x80, 0x93, 0x6f, 0x6a, 0x49, 0x62, 0x4a, 0x62, 0x49, 0xe2, 0x2b, 0x0c, 0xf5,
	0x49, 0x6c, 0xe0, 0x60, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x67, 0x0b, 0xfa, 0x31, 0xaf,
	0x01, 0x00, 0x00,
}

func (m *Mappable) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mappable) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mappable) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Classification != nil {
		{
			size, err := m.Classification.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMappableV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMappableV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovMappableV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mappable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Classification != nil {
		l = m.Classification.Size()
		n += 1 + l + sovMappableV1(uint64(l))
	}
	return n
}

func sovMappableV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMappableV1(x uint64) (n int) {
	return sovMappableV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mappable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMappableV1
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
			return fmt.Errorf("proto: Mappable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mappable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Classification", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMappableV1
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
				return ErrInvalidLengthMappableV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMappableV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Classification == nil {
				m.Classification = &base.Document{}
			}
			if err := m.Classification.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMappableV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMappableV1
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
func skipMappableV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMappableV1
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
					return 0, ErrIntOverflowMappableV1
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
					return 0, ErrIntOverflowMappableV1
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
				return 0, ErrInvalidLengthMappableV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMappableV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMappableV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMappableV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMappableV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMappableV1 = fmt.Errorf("proto: unexpected end of group")
)
