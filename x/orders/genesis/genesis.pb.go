// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/genesis/genesis.proto

package genesis

import (
	fmt "fmt"
	record "github.com/AssetMantle/modules/x/orders/record"
	base "github.com/AssetMantle/schema/go/lists/base"
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
	Records       []*record.Record    `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	ParameterList *base.ParameterList `protobuf:"bytes,2,opt,name=parameter_list,json=parameterList,proto3" json:"parameter_list,omitempty"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e82ddf984d76b2, []int{0}
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
	proto.RegisterType((*Genesis)(nil), "assetmantle.modules.orders.genesis.Genesis")
}

func init() { proto.RegisterFile("orders/genesis/genesis.proto", fileDescriptor_12e82ddf984d76b2) }

var fileDescriptor_12e82ddf984d76b2 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0x4b, 0xfb, 0x40,
	0x18, 0xc7, 0x73, 0xf9, 0x95, 0x5f, 0x21, 0x55, 0x87, 0xe2, 0x50, 0x8a, 0xa4, 0xa5, 0xa0, 0x54,
	0x90, 0x3b, 0xa8, 0xdb, 0x6d, 0xad, 0x43, 0x17, 0x43, 0x43, 0xdd, 0x24, 0x20, 0xd7, 0xe4, 0x48,
	0x03, 0x4d, 0xaf, 0xdc, 0x5d, 0xc1, 0x97, 0xe0, 0xe8, 0x4b, 0x10, 0x37, 0xf5, 0x8d, 0x88, 0x53,
	0x47, 0x47, 0x49, 0x37, 0x5f, 0x85, 0x24, 0xf7, 0x54, 0x92, 0xc5, 0x4e, 0x4f, 0xfe, 0x7c, 0x3e,
	0xcf, 0xf7, 0xb9, 0x7b, 0x9c, 0x13, 0x21, 0x23, 0x2e, 0x15, 0x89, 0xf9, 0x92, 0xab, 0xe4, 0xb7,
	0xe2, 0x95, 0x14, 0x5a, 0x34, 0x7b, 0x4c, 0x29, 0xae, 0x53, 0xb6, 0xd4, 0x0b, 0x8e, 0x53, 0x11,
	0xad, 0x17, 0x5c, 0x61, 0x63, 0x60, 0x20, 0xdb, 0xc7, 0xb1, 0x88, 0x45, 0x81, 0x93, 0xfc, 0xc9,
	0x98, 0xed, 0x36, 0xf4, 0x95, 0x3c, 0x14, 0x32, 0x82, 0x02, 0xff, 0x3a, 0x8b, 0x44, 0x69, 0x45,
	0x66, 0x4c, 0x71, 0xb2, 0x62, 0x92, 0xa5, 0x5c, 0x73, 0x79, 0x97, 0x7f, 0x34, 0x40, 0xef, 0x05,
	0x39, 0xf5, 0xb1, 0x69, 0xdf, 0xbc, 0x72, 0xea, 0x46, 0x56, 0x2d, 0xd4, 0xfd, 0xd7, 0x6f, 0x0c,
	0xce, 0xf1, 0x1f, 0x43, 0x41, 0xce, 0xb4, 0x28, 0xd3, 0x9d, 0xd9, 0xbc, 0x71, 0x8e, 0xaa, 0x41,
	0x2d, 0xbb, 0x8b, 0xfa, 0x8d, 0xc1, 0x45, 0xa5, 0x97, 0x0a, 0xe7, 0x3c, 0x65, 0xb8, 0x98, 0x0e,
	0xe7, 0xd3, 0x61, 0x7f, 0x27, 0x5d, 0x27, 0x4a, 0x4f, 0x0f, 0x57, 0xe5, 0x57, 0x5a, 0x7b, 0x78,
	0xea, 0x58, 0xa3, 0x37, 0xfb, 0x3d, 0x73, 0xd1, 0x26, 0x73, 0xd1, 0x57, 0xe6, 0xa2, 0xc7, 0xad,
	0x6b, 0x6d, 0xb6, 0xae, 0xf5, 0xb9, 0x75, 0x2d, 0xe7, 0x2c, 0x14, 0x29, 0xde, 0x7f, 0x83, 0xa3,
	0x03, 0x38, 0xab, 0x9f, 0x1f, 0xde, 0x47, 0xb7, 0x24, 0x4e, 0xf4, 0x7c, 0x3d, 0xc3, 0xa1, 0x48,
	0xc9, 0x30, 0xd7, 0xbd, 0x42, 0x27, 0xa0, 0x93, 0x7b, 0x52, 0x5d, 0xda, 0xb3, 0x5d, 0x1b, 0x7a,
	0x93, 0xf1, 0xab, 0xdd, 0x1b, 0x96, 0xc2, 0x3c, 0x08, 0x9b, 0x98, 0x30, 0xc8, 0xf8, 0xa8, 0x40,
	0x01, 0x40, 0x81, 0x81, 0x02, 0x80, 0x32, 0x1b, 0xef, 0x87, 0x82, 0xb1, 0x3f, 0xf2, 0xb8, 0x66,
	0x11, 0xd3, 0xec, 0xdb, 0x3e, 0x2d, 0x09, 0x94, 0x82, 0x41, 0xa9, 0x51, 0x28, 0x05, 0x67, 0xf6,
	0xbf, 0x58, 0xf0, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xe1, 0x80, 0xbe, 0x77, 0x02,
	0x00, 0x00,
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
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Records) > 0 {
		for _, e := range m.Records {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
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
			m.Records = append(m.Records, &record.Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
