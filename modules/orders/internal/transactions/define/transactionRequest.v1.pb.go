// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/orders/internal/transactions/define/transactionRequest.v1.proto

package define

import (
	fmt "fmt"
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

type TransactionRequest struct {
	From                    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID                  string `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ImmutableMetaProperties string `protobuf:"bytes,3,opt,name=immutable_meta_properties,json=immutableMetaProperties,proto3" json:"immutable_meta_properties,omitempty"`
	ImmutableProperties     string `protobuf:"bytes,4,opt,name=immutable_properties,json=immutableProperties,proto3" json:"immutable_properties,omitempty"`
	MutableMetaProperties   string `protobuf:"bytes,5,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3" json:"mutable_meta_properties,omitempty"`
	MutableProperties       string `protobuf:"bytes,6,opt,name=mutable_properties,json=mutableProperties,proto3" json:"mutable_properties,omitempty"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe0b4db7aca1565, []int{0}
}
func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TransactionRequest)(nil), "orders.transactions.define.TransactionRequest")
}

func init() {
	proto.RegisterFile("modules/orders/internal/transactions/define/transactionRequest.v1.proto", fileDescriptor_ffe0b4db7aca1565)
}

var fileDescriptor_ffe0b4db7aca1565 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0x93, 0xb4, 0xb7, 0xdc, 0x3b, 0xbb, 0x3b, 0x56, 0x9a, 0x16, 0x89, 0xe2, 0x42, 0xdc,
	0x98, 0x50, 0x04, 0x17, 0xd9, 0x59, 0x0a, 0x45, 0xa1, 0x34, 0x94, 0xe2, 0x42, 0x0a, 0x61, 0xda,
	0xfc, 0xad, 0x81, 0x24, 0x53, 0x67, 0x26, 0x3e, 0x83, 0x4b, 0x9f, 0x40, 0xc4, 0x9d, 0x3e, 0x89,
	0xb8, 0xea, 0xd2, 0xa5, 0xa4, 0x3b, 0x9f, 0x42, 0x32, 0xa1, 0x75, 0xa0, 0xed, 0xc2, 0x55, 0x26,
	0x73, 0xce, 0x37, 0xe7, 0xe7, 0xf0, 0xa3, 0x4e, 0x4c, 0x83, 0x34, 0x02, 0xee, 0x50, 0x16, 0x00,
	0xe3, 0x4e, 0x98, 0x08, 0x60, 0x09, 0x89, 0x1c, 0xc1, 0x48, 0xc2, 0xc9, 0x58, 0x84, 0x34, 0xe1,
	0x4e, 0x00, 0x93, 0x30, 0x01, 0xf5, 0xae, 0x0f, 0xb7, 0x29, 0x70, 0x61, 0xdf, 0x35, 0xed, 0x19,
	0xa3, 0x82, 0xe2, 0x46, 0xf1, 0x80, 0xad, 0x72, 0x76, 0xc1, 0x35, 0xaa, 0x53, 0x3a, 0xa5, 0xd2,
	0xe6, 0xe4, 0xa7, 0x82, 0x38, 0x7c, 0x34, 0x10, 0x1e, 0xac, 0xbd, 0x88, 0x31, 0x2a, 0x4f, 0x18,
	0x8d, 0x4d, 0xfd, 0x40, 0x3f, 0xfe, 0xd7, 0x97, 0x67, 0x6c, 0xa2, 0xbf, 0xf9, 0xd7, 0x0f, 0xfd,
	0xc0, 0x34, 0xe4, 0x7d, 0x25, 0xff, 0xbf, 0x68, 0x63, 0x17, 0xd5, 0xc3, 0x38, 0x4e, 0x05, 0x19,
	0x45, 0xe0, 0xc7, 0x20, 0x88, 0x3f, 0x63, 0x74, 0x06, 0x4c, 0x84, 0xc0, 0xcd, 0x92, 0xb4, 0xd6,
	0x56, 0x86, 0x2e, 0x08, 0xe2, 0xad, 0x64, 0xdc, 0x44, 0xd5, 0x1f, 0x56, 0xc1, 0xca, 0x12, 0xdb,
	0x59, 0x69, 0x0a, 0x72, 0x86, 0x6a, 0xdb, 0xc2, 0xfe, 0x48, 0x6a, 0x77, 0x73, 0xd4, 0x09, 0xc2,
	0x1b, 0x82, 0x2a, 0x12, 0xf9, 0xbf, 0x16, 0xe3, 0x96, 0xef, 0x9f, 0xf6, 0xb5, 0xd6, 0x8b, 0xf1,
	0x96, 0x59, 0xfa, 0x3c, 0xb3, 0xf4, 0xcf, 0xcc, 0xd2, 0x1f, 0x16, 0x96, 0x36, 0x5f, 0x58, 0xda,
	0xc7, 0xc2, 0xd2, 0x90, 0x35, 0xa6, 0xb1, 0xbd, 0xbd, 0xf1, 0x56, 0x7d, 0xbd, 0xd8, 0xab, 0xa6,
	0x97, 0xd7, 0xee, 0xe9, 0xd7, 0x97, 0xd3, 0x50, 0xdc, 0xa4, 0x23, 0x7b, 0x4c, 0x63, 0xe7, 0x9c,
	0x73, 0x10, 0x5d, 0x92, 0x88, 0x08, 0x9c, 0xe5, 0x2a, 0xfc, 0x62, 0x25, 0x9e, 0x8d, 0x52, 0x6f,
	0xd0, 0x7e, 0x35, 0x1a, 0xbd, 0x62, 0x96, 0x81, 0x3a, 0x4b, 0x5b, 0x5a, 0xde, 0x97, 0xe2, 0x50,
	0x15, 0x87, 0x85, 0x98, 0x19, 0x47, 0xdb, 0xc5, 0x61, 0xc7, 0x6b, 0xe5, 0x3d, 0x06, 0x44, 0x90,
	0x2f, 0x63, 0xaf, 0x30, 0xba, 0xae, 0xea, 0x74, 0xdd, 0xc2, 0x3a, 0xaa, 0xc8, 0x9d, 0x3a, 0xfd,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x75, 0x1a, 0x99, 0xda, 0xd0, 0x02, 0x00, 0x00,
}

func (m *TransactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MutableProperties) > 0 {
		i -= len(m.MutableProperties)
		copy(dAtA[i:], m.MutableProperties)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.MutableProperties)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MutableMetaProperties) > 0 {
		i -= len(m.MutableMetaProperties)
		copy(dAtA[i:], m.MutableMetaProperties)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.MutableMetaProperties)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ImmutableProperties) > 0 {
		i -= len(m.ImmutableProperties)
		copy(dAtA[i:], m.ImmutableProperties)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.ImmutableProperties)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ImmutableMetaProperties) > 0 {
		i -= len(m.ImmutableMetaProperties)
		copy(dAtA[i:], m.ImmutableMetaProperties)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.ImmutableMetaProperties)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FromID) > 0 {
		i -= len(m.FromID)
		copy(dAtA[i:], m.FromID)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.FromID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransactionRequestV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransactionRequestV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.FromID)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.ImmutableMetaProperties)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.ImmutableProperties)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.MutableMetaProperties)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.MutableProperties)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	return n
}

func sovTransactionRequestV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransactionRequestV1(x uint64) (n int) {
	return sovTransactionRequestV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransactionRequestV1
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
			return fmt.Errorf("proto: TransactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableMetaProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImmutableMetaProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImmutableProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MutableMetaProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
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
				return ErrInvalidLengthTransactionRequestV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MutableProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransactionRequestV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransactionRequestV1
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
func skipTransactionRequestV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransactionRequestV1
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
					return 0, ErrIntOverflowTransactionRequestV1
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
					return 0, ErrIntOverflowTransactionRequestV1
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
				return 0, ErrInvalidLengthTransactionRequestV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransactionRequestV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransactionRequestV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransactionRequestV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransactionRequestV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransactionRequestV1 = fmt.Errorf("proto: unexpected end of group")
)