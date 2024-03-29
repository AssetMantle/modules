// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: maintainers/genesis/genesis.proto

package genesis

import (
	fmt "fmt"
	record "github.com/AssetMantle/modules/x/maintainers/record"
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
	return fileDescriptor_ccc39df958fd39c6, []int{0}
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
	proto.RegisterType((*Genesis)(nil), "assetmantle.modules.maintainers.genesis.Genesis")
}

func init() { proto.RegisterFile("maintainers/genesis/genesis.proto", fileDescriptor_ccc39df958fd39c6) }

var fileDescriptor_ccc39df958fd39c6 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x73, 0xb1, 0x58, 0x48, 0xd5, 0xa1, 0x38, 0x94, 0x0e, 0x69, 0x75, 0x69, 0x51, 0xb9,
	0x83, 0xd6, 0xe9, 0xb6, 0x76, 0xa9, 0x83, 0x81, 0x50, 0x37, 0x09, 0xc8, 0x35, 0xfd, 0x48, 0x03,
	0xbd, 0x5c, 0xc9, 0x5d, 0xc1, 0x47, 0x70, 0xf4, 0x11, 0xc4, 0xd1, 0xc1, 0xe7, 0x50, 0xa7, 0x8e,
	0x8e, 0x92, 0x6e, 0x3e, 0x85, 0x24, 0xb9, 0xca, 0x15, 0x04, 0x3b, 0x7d, 0x49, 0xf8, 0xff, 0xfe,
	0xbf, 0x2f, 0x97, 0x38, 0x27, 0x9c, 0xc5, 0x89, 0x62, 0x71, 0x02, 0xa9, 0x24, 0x11, 0x24, 0x20,
	0xe3, 0xdf, 0x89, 0x17, 0xa9, 0x50, 0xa2, 0xde, 0x61, 0x52, 0x82, 0xe2, 0x2c, 0x51, 0x73, 0xc0,
	0x5c, 0x4c, 0x97, 0x73, 0x90, 0xd8, 0xc0, 0xb0, 0x8e, 0x37, 0x8f, 0x23, 0x11, 0x89, 0x82, 0x21,
	0xf9, 0x55, 0x89, 0x37, 0x5b, 0xa6, 0x21, 0x85, 0x50, 0xa4, 0x53, 0x3d, 0x36, 0x81, 0x79, 0x2c,
	0x95, 0x24, 0x13, 0x26, 0x81, 0x2c, 0x58, 0xca, 0x38, 0x28, 0x48, 0xef, 0xf2, 0x87, 0x65, 0xe0,
	0xf4, 0x15, 0x39, 0xd5, 0x51, 0xe9, 0xa8, 0x5f, 0x39, 0xd5, 0x12, 0x96, 0x0d, 0xd4, 0xde, 0xeb,
	0xd6, 0x7a, 0x18, 0xff, 0xb7, 0x9e, 0x96, 0x8d, 0x8b, 0x31, 0xde, 0xe0, 0xf5, 0x1b, 0xe7, 0x68,
	0xdb, 0xd6, 0xb0, 0xdb, 0xa8, 0x5b, 0xeb, 0x5d, 0x6c, 0x15, 0xca, 0x70, 0x06, 0x9c, 0xe1, 0x62,
	0x45, 0x9c, 0xaf, 0x88, 0xfd, 0x0d, 0x74, 0x1d, 0x4b, 0x35, 0x3e, 0x5c, 0x98, 0xb7, 0xb4, 0xf2,
	0xf0, 0xd4, 0xb2, 0x86, 0xef, 0xf6, 0x5b, 0xe6, 0xa2, 0x55, 0xe6, 0xa2, 0xaf, 0xcc, 0x45, 0x8f,
	0x6b, 0xd7, 0x5a, 0xad, 0x5d, 0xeb, 0x73, 0xed, 0x5a, 0xce, 0x79, 0x28, 0x38, 0xde, 0xf1, 0x40,
	0x87, 0x07, 0xfa, 0xad, 0xfd, 0xfc, 0x18, 0x7c, 0x74, 0x7b, 0x19, 0xc5, 0x6a, 0xb6, 0x9c, 0xe0,
	0x50, 0x70, 0x32, 0xc8, 0x3b, 0xbc, 0xa2, 0x83, 0xe8, 0x0e, 0x72, 0x4f, 0xfe, 0xf8, 0x9a, 0xcf,
	0x76, 0x65, 0xe0, 0x79, 0xa3, 0x17, 0xbb, 0x33, 0x30, 0xb4, 0x9e, 0xd6, 0x7a, 0x86, 0x56, 0xdb,
	0x3e, 0xb6, 0x92, 0x81, 0x4e, 0x06, 0x46, 0x32, 0xd0, 0xc9, 0xcc, 0xee, 0xef, 0x98, 0x0c, 0x46,
	0xfe, 0xd0, 0x03, 0xc5, 0xa6, 0x4c, 0xb1, 0x6f, 0xfb, 0xcc, 0xa0, 0x28, 0xd5, 0x18, 0xa5, 0x06,
	0x47, 0xa9, 0x06, 0x27, 0xfb, 0xc5, 0x3f, 0xd0, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x88, 0x07,
	0x76, 0x67, 0xa9, 0x02, 0x00, 0x00,
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
