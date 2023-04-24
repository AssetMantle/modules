// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/transactions/mutate/transaction_response.proto

package mutate

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c96c1e603dfcb44, []int{0}
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

func init() {
	proto.RegisterType((*TransactionResponse)(nil), "assetmantle.modules.identities.transactions.mutate.TransactionResponse")
}

func init() {
	proto.RegisterFile("x/identities/internal/transactions/mutate/transaction_response.proto", fileDescriptor_6c96c1e603dfcb44)
}

var fileDescriptor_6c96c1e603dfcb44 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd1, 0xb1, 0x4a, 0xc4, 0x30,
	0x1c, 0xc7, 0xf1, 0xb6, 0xa2, 0x43, 0x47, 0x45, 0x70, 0xca, 0xe0, 0x03, 0x24, 0x70, 0xc2, 0x0d,
	0x19, 0x84, 0x16, 0x41, 0x6f, 0x08, 0x14, 0xe9, 0x24, 0x05, 0xc9, 0xb5, 0x7f, 0x34, 0xd0, 0x26,
	0x47, 0xf3, 0x2f, 0xf8, 0x18, 0x3e, 0xc3, 0x8d, 0x3e, 0x89, 0x38, 0xdd, 0xe8, 0x28, 0xed, 0xe6,
	0x53, 0xc8, 0xb5, 0x85, 0x44, 0x70, 0xf0, 0xd6, 0xc0, 0xf7, 0xf3, 0xff, 0xd1, 0xc6, 0x37, 0x2f,
	0x4c, 0x55, 0xa0, 0x51, 0xa1, 0x02, 0xcb, 0x94, 0x46, 0x68, 0xb5, 0xac, 0x19, 0xb6, 0x52, 0x5b,
	0x59, 0xa2, 0x32, 0xda, 0xb2, 0xa6, 0x43, 0x89, 0xe0, 0xbf, 0x3d, 0xb6, 0x60, 0x37, 0x46, 0x5b,
	0xa0, 0x9b, 0xd6, 0xa0, 0x39, 0x5d, 0x48, 0x6b, 0x01, 0x1b, 0xa9, 0xb1, 0x06, 0xda, 0x98, 0xaa,
	0xab, 0xc1, 0x52, 0xe7, 0x52, 0x9f, 0xa3, 0x13, 0x77, 0x79, 0x1e, 0x9f, 0xe5, 0xee, 0xf9, 0x7e,
	0x06, 0xd3, 0xed, 0xd1, 0x7b, 0x4f, 0xc2, 0x5d, 0x4f, 0xc2, 0xaf, 0x9e, 0x84, 0xaf, 0x03, 0x09,
	0x76, 0x03, 0x09, 0x3e, 0x07, 0x12, 0xc4, 0xcb, 0xd2, 0x34, 0xf4, 0xf0, 0x4b, 0xe9, 0xc5, 0x1f,
	0x77, 0xb2, 0xfd, 0xee, 0x2c, 0x7c, 0xb8, 0x7b, 0x52, 0xf8, 0xdc, 0xad, 0x69, 0x69, 0x1a, 0x96,
	0xec, 0x69, 0x31, 0xd2, 0x6c, 0xa6, 0xd9, 0xbf, 0x3f, 0xcf, 0x36, 0x3a, 0x4e, 0xc4, 0x2a, 0x17,
	0x6f, 0xd1, 0x22, 0xf1, 0x26, 0x8a, 0x79, 0xe2, 0xca, 0x4d, 0xcc, 0xfd, 0x89, 0x62, 0x8c, 0x3f,
	0x7e, 0x45, 0xc5, 0x1c, 0x15, 0x2e, 0x2a, 0xfc, 0xa8, 0x98, 0xa2, 0x3e, 0xba, 0x3e, 0x3c, 0x2a,
	0x6e, 0xb3, 0x54, 0x00, 0xca, 0x4a, 0xa2, 0xfc, 0x8e, 0x96, 0x1e, 0xc0, 0xf9, 0x2c, 0x70, 0xee,
	0x08, 0xce, 0x7d, 0x83, 0xf3, 0x09, 0x59, 0x9f, 0x8c, 0xbf, 0xfd, 0xea, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x37, 0x72, 0xea, 0xc9, 0x3e, 0x02, 0x00, 0x00,
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
	return len(dAtA) - i, nil
}

func encodeVarintTransactionResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransactionResponse(v)
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
	return n
}

func sovTransactionResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransactionResponse(x uint64) (n int) {
	return sovTransactionResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransactionResponse
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
		default:
			iNdEx = preIndex
			skippy, err := skipTransactionResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransactionResponse
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
func skipTransactionResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransactionResponse
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
					return 0, ErrIntOverflowTransactionResponse
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
					return 0, ErrIntOverflowTransactionResponse
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
				return 0, ErrInvalidLengthTransactionResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransactionResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransactionResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransactionResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransactionResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransactionResponse = fmt.Errorf("proto: unexpected end of group")
)