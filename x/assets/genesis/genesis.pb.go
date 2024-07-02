// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/assets/genesis/genesis.proto

package genesis

import (
	fmt "fmt"
	record "github.com/AssetMantle/modules/x/assets/record"
	base "github.com/AssetMantle/schema/lists/base"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Genesis struct {
	Records       []*record.Record    `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	ParameterList *base.ParameterList `protobuf:"bytes,2,opt,name=parameter_list,json=parameterList,proto3" json:"parameter_list,omitempty"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a6c3849ff34e961, []int{0}
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
	proto.RegisterType((*Genesis)(nil), "AssetMantle.modules.x.assets.genesis.Genesis")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/assets/genesis/genesis.proto", fileDescriptor_6a6c3849ff34e961)
}

var fileDescriptor_6a6c3849ff34e961 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xb3, 0xf9, 0xff, 0xb5, 0x90, 0xaa, 0x87, 0xe2, 0xa1, 0xf4, 0xb0, 0x16, 0xf1, 0x10,
	0x50, 0x76, 0x4b, 0xbd, 0xed, 0xad, 0x05, 0xe9, 0xc5, 0x40, 0xa8, 0x97, 0x20, 0x01, 0xd9, 0xa6,
	0x4b, 0x1a, 0x68, 0xba, 0x25, 0xbb, 0x85, 0x3e, 0x82, 0x47, 0x1f, 0x41, 0x3c, 0xea, 0x03, 0xf8,
	0x0a, 0xe2, 0xa9, 0x47, 0x8f, 0x92, 0xde, 0x7c, 0x0a, 0x49, 0x76, 0x2a, 0xcd, 0xa5, 0xf4, 0x34,
	0xc9, 0xf0, 0xfd, 0xe6, 0xfb, 0x76, 0xc6, 0xe9, 0xf6, 0x94, 0x12, 0xda, 0xe3, 0x33, 0x3d, 0x15,
	0x34, 0x95, 0xe3, 0xc5, 0x54, 0x28, 0xba, 0xa4, 0xbc, 0xe8, 0x2a, 0x1a, 0x8b, 0x99, 0x50, 0xc9,
	0x5f, 0x25, 0xf3, 0x4c, 0x6a, 0xd9, 0xb8, 0xd8, 0x62, 0x08, 0x30, 0x64, 0x49, 0x0c, 0x43, 0x40,
	0xdb, 0x3a, 0x8d, 0x65, 0x2c, 0x4b, 0x80, 0x16, 0x5f, 0x86, 0x6d, 0x75, 0x76, 0xfa, 0x65, 0x22,
	0x92, 0xd9, 0x18, 0x0a, 0x10, 0x95, 0x84, 0x2a, 0x9a, 0x88, 0x94, 0xd3, 0x69, 0xa2, 0xb4, 0xa2,
	0x23, 0xae, 0x04, 0x9d, 0xf3, 0x8c, 0xa7, 0x42, 0x8b, 0xec, 0xa1, 0x68, 0x1a, 0xe6, 0xfc, 0x0d,
	0x39, 0xb5, 0x81, 0xc9, 0xd1, 0xb8, 0x71, 0x6a, 0x66, 0x9e, 0x6a, 0xa2, 0xf6, 0x3f, 0xb7, 0xde,
	0xbd, 0x24, 0x3b, 0xf3, 0x83, 0xf9, 0xb0, 0x2c, 0xc3, 0x0d, 0xdb, 0xb8, 0x73, 0x4e, 0xaa, 0x56,
	0x4d, 0xbb, 0x8d, 0xdc, 0x7a, 0xf7, 0xaa, 0x32, 0xcd, 0xe4, 0x23, 0x65, 0x3e, 0x52, 0xe4, 0x23,
	0xfe, 0x06, 0xba, 0x4d, 0x94, 0x1e, 0x1e, 0xcf, 0xb7, 0x7f, 0xd9, 0xff, 0xc7, 0xe7, 0x33, 0xab,
	0xff, 0x6e, 0x7f, 0xe4, 0x18, 0xad, 0x72, 0x8c, 0xbe, 0x73, 0x8c, 0x9e, 0xd6, 0xd8, 0x5a, 0xad,
	0xb1, 0xf5, 0xb5, 0xc6, 0x96, 0xe3, 0x46, 0x32, 0x25, 0xfb, 0xac, 0xbb, 0x7f, 0x04, 0xef, 0xf5,
	0x8b, 0x05, 0xf8, 0xe8, 0x9e, 0xc6, 0x89, 0x9e, 0x2c, 0x46, 0x24, 0x92, 0x29, 0xdd, 0xe7, 0xc6,
	0x2f, 0xf6, 0x41, 0xcf, 0x0b, 0x7a, 0x83, 0x57, 0xbb, 0x72, 0x5e, 0x0f, 0xfc, 0x02, 0x93, 0x42,
	0x11, 0xb0, 0xf9, 0xac, 0xc8, 0x42, 0x90, 0x85, 0x41, 0x68, 0x64, 0x21, 0xc8, 0x72, 0xbb, 0xb3,
	0x8f, 0x2c, 0x1c, 0xf8, 0x7d, 0x4f, 0x68, 0x3e, 0xe6, 0x9a, 0xff, 0xd8, 0xee, 0x16, 0xc2, 0x18,
	0x30, 0x8c, 0x05, 0x8c, 0x19, 0x8a, 0x31, 0xc0, 0x46, 0x87, 0xe5, 0xb9, 0xaf, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xf1, 0x5c, 0x8a, 0xe6, 0xc6, 0x02, 0x00, 0x00,
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
