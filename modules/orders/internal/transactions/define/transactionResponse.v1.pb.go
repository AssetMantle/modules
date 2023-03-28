// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/orders/internal/transactions/define/transactionResponse.v1.proto

package define

import (
	fmt "fmt"
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

type TransactionResponse struct {
	ClassificationId string `protobuf:"bytes,1,opt,name=classification_id,json=classificationId,proto3" json:"classification_id,omitempty"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f3bd6d150e50798, []int{0}
}
func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetClassificationId() string {
	if m != nil {
		return m.ClassificationId
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionResponse)(nil), "orders.transactions.define.TransactionResponse")
}

func init() {
	proto.RegisterFile("modules/orders/internal/transactions/define/transactionResponse.v1.proto", fileDescriptor_6f3bd6d150e50798)
}

var fileDescriptor_6f3bd6d150e50798 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0xc8, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0x4f, 0x49, 0x4d, 0xcb, 0xcc, 0x4b, 0x45, 0x16, 0x0b, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0xd5, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x82, 0x98, 0xa0, 0x87,
	0xac, 0x51, 0x0f, 0xa2, 0x51, 0xc9, 0x89, 0x4b, 0x38, 0x04, 0x53, 0xaf, 0x90, 0x36, 0x97, 0x60,
	0x72, 0x4e, 0x62, 0x71, 0x71, 0x66, 0x5a, 0x66, 0x72, 0x22, 0x48, 0x26, 0x3e, 0x33, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x48, 0x00, 0x55, 0xc2, 0x33, 0xc5, 0x69, 0x15, 0xd3, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0xc9, 0x25, 0xe7, 0xe7, 0xea, 0xe1, 0xb6, 0xde,
	0x49, 0x0a, 0x8b, 0xe5, 0x61, 0x86, 0x01, 0x20, 0x67, 0x07, 0x30, 0x46, 0x79, 0xa5, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0x16, 0x17, 0xa7, 0x96, 0xf8, 0x26, 0xe6,
	0x95, 0xe4, 0xa4, 0xea, 0xc3, 0x42, 0x86, 0x84, 0x10, 0x5a, 0xc4, 0xc4, 0xec, 0x1f, 0xe2, 0xb2,
	0x8a, 0x49, 0xca, 0x1f, 0xe2, 0x98, 0x10, 0x64, 0xc7, 0xb8, 0x80, 0x95, 0x9c, 0x82, 0x49, 0xc6,
	0x20, 0x4b, 0xc6, 0x40, 0x24, 0x1f, 0x31, 0xa9, 0xe1, 0x96, 0x8c, 0x71, 0x0f, 0x70, 0xf2, 0x4d,
	0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x7c, 0xc5, 0x24, 0x03, 0x51, 0x68, 0x65, 0x85, 0xac, 0xd2,
	0xca, 0x0a, 0xa2, 0x34, 0x89, 0x0d, 0x1c, 0x27, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0x51, 0xd8, 0x5e, 0xdf, 0x01, 0x00, 0x00,
}

func (m *TransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassificationId) > 0 {
		i -= len(m.ClassificationId)
		copy(dAtA[i:], m.ClassificationId)
		i = encodeVarintTransactionResponseV1(dAtA, i, uint64(len(m.ClassificationId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransactionResponseV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransactionResponseV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassificationId)
	if l > 0 {
		n += 1 + l + sovTransactionResponseV1(uint64(l))
	}
	return n
}

func sovTransactionResponseV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransactionResponseV1(x uint64) (n int) {
	return sovTransactionResponseV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransactionResponseV1
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
			return fmt.Errorf("proto: TransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionResponseV1
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
				return ErrInvalidLengthTransactionResponseV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionResponseV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassificationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransactionResponseV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransactionResponseV1
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
func skipTransactionResponseV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransactionResponseV1
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
					return 0, ErrIntOverflowTransactionResponseV1
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
					return 0, ErrIntOverflowTransactionResponseV1
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
				return 0, ErrInvalidLengthTransactionResponseV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransactionResponseV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransactionResponseV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransactionResponseV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransactionResponseV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransactionResponseV1 = fmt.Errorf("proto: unexpected end of group")
)
