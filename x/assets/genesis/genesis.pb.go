// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/genesis/genesis.proto

package genesis

import (
	fmt "fmt"
	record "github.com/AssetMantle/modules/x/assets/record"
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
	return fileDescriptor_be3ad5ba5ab94325, []int{0}
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
	proto.RegisterType((*Genesis)(nil), "assetmantle.modules.assets.genesis.Genesis")
}

func init() { proto.RegisterFile("assets/genesis/genesis.proto", fileDescriptor_be3ad5ba5ab94325) }

var fileDescriptor_be3ad5ba5ab94325 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x73, 0x79, 0xcb, 0x5b, 0x48, 0xd5, 0xa1, 0x38, 0x94, 0x22, 0x69, 0x29, 0x28, 0x15,
	0xe4, 0x0e, 0xea, 0x76, 0x5b, 0xea, 0xd0, 0xc5, 0x40, 0xa8, 0x9b, 0x04, 0xe4, 0x9a, 0x1e, 0x69,
	0xa0, 0xd7, 0x2b, 0xb9, 0x2b, 0xf8, 0x11, 0x1c, 0xfd, 0x08, 0xe2, 0xa6, 0x7e, 0x11, 0x71, 0xea,
	0xe8, 0x28, 0xe9, 0xe6, 0xa7, 0x90, 0xdc, 0x3d, 0x95, 0x66, 0xb1, 0xd3, 0x93, 0xe4, 0x7e, 0xbf,
	0xe7, 0xff, 0x5c, 0x1e, 0xef, 0x84, 0x29, 0xc5, 0xb5, 0x22, 0x29, 0x5f, 0x70, 0x95, 0xfd, 0x56,
	0xbc, 0xcc, 0xa5, 0x96, 0xcd, 0x9e, 0x39, 0x15, 0x6c, 0xa1, 0xe7, 0x1c, 0x0b, 0x39, 0x5d, 0xcd,
	0xb9, 0xc2, 0xd6, 0xc0, 0x40, 0xb6, 0x8f, 0x53, 0x99, 0x4a, 0x83, 0x93, 0xf2, 0xc9, 0x9a, 0xed,
	0x36, 0xf4, 0xcd, 0x79, 0x22, 0xf3, 0x29, 0x14, 0x38, 0xeb, 0xcc, 0x33, 0xa5, 0x15, 0x99, 0x30,
	0xc5, 0xc9, 0x92, 0xe5, 0x4c, 0x70, 0xcd, 0xf3, 0xbb, 0xf2, 0xa3, 0x05, 0x7a, 0x2f, 0xc8, 0xab,
	0x8f, 0x6c, 0xfb, 0xe6, 0x95, 0x57, 0xb7, 0xb2, 0x6a, 0xa1, 0xee, 0xbf, 0x7e, 0x63, 0x70, 0x8e,
	0xff, 0x18, 0x0a, 0x72, 0xc6, 0xa6, 0x8c, 0xb7, 0x66, 0xf3, 0xc6, 0x3b, 0xaa, 0x06, 0xb5, 0xdc,
	0x2e, 0xea, 0x37, 0x06, 0x17, 0x95, 0x5e, 0x2a, 0x99, 0x71, 0xc1, 0xb0, 0x99, 0x0e, 0x97, 0xd3,
	0xe1, 0x68, 0x2b, 0x5d, 0x67, 0x4a, 0x8f, 0x0f, 0x97, 0xbb, 0xaf, 0xb4, 0xf6, 0xf0, 0xd4, 0x71,
	0x86, 0x6f, 0xee, 0x7b, 0xe1, 0xa3, 0x75, 0xe1, 0xa3, 0xaf, 0xc2, 0x47, 0x8f, 0x1b, 0xdf, 0x59,
	0x6f, 0x7c, 0xe7, 0x73, 0xe3, 0x3b, 0xde, 0x59, 0x22, 0x05, 0xde, 0xff, 0x07, 0x87, 0x07, 0x70,
	0xd7, 0xa8, 0xbc, 0x7c, 0x84, 0x6e, 0x49, 0x9a, 0xe9, 0xd9, 0x6a, 0x82, 0x13, 0x29, 0x48, 0x50,
	0xa2, 0xa1, 0xd1, 0x09, 0xe8, 0xe4, 0x9e, 0x54, 0x97, 0xf6, 0xec, 0xd6, 0x82, 0x30, 0x18, 0xbd,
	0xba, 0xbd, 0x60, 0x27, 0x2c, 0x84, 0xb0, 0xc0, 0x86, 0x41, 0xc6, 0x47, 0x05, 0x8a, 0x01, 0x8a,
	0x2d, 0x14, 0x03, 0x54, 0xb8, 0x78, 0x3f, 0x14, 0x8f, 0xa2, 0x61, 0xc8, 0x35, 0x9b, 0x32, 0xcd,
	0xbe, 0xdd, 0xd3, 0x1d, 0x81, 0x52, 0x30, 0x28, 0xb5, 0x0a, 0xa5, 0xe0, 0x4c, 0xfe, 0x9b, 0x05,
	0x5f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x74, 0x2c, 0x2e, 0x77, 0x02, 0x00, 0x00,
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
