// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/splits/internal/transactions/send/transactionResponse.proto

package send

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
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c928b0dfeea40e9, []int{0}
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
	proto.RegisterType((*TransactionResponse)(nil), "splits.transactions.send.TransactionResponse")
}

func init() {
	proto.RegisterFile("x/splits/internal/transactions/send/transactionResponse.proto", fileDescriptor_2c928b0dfeea40e9)
}

var fileDescriptor_2c928b0dfeea40e9 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xad, 0xd0, 0x2f, 0x2e,
	0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x29,
	0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x4e, 0xcd, 0x4b, 0x41,
	0x16, 0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x80, 0x68, 0xd6, 0x43, 0xd6, 0xa3, 0x07, 0xd2, 0xa3, 0x24, 0xca, 0x25, 0x1c, 0x82, 0xa9,
	0xcd, 0x69, 0x1a, 0xd3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0xc9, 0x24,
	0xe7, 0xe7, 0xea, 0xe1, 0x32, 0xcf, 0x49, 0x02, 0x8b, 0x69, 0x01, 0x20, 0x37, 0x04, 0x30, 0x46,
	0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0x16, 0x17, 0xa7,
	0x96, 0xf8, 0x26, 0xe6, 0x95, 0xe4, 0xa4, 0xea, 0xe7, 0xe6, 0xa7, 0x94, 0xe6, 0xa4, 0x16, 0xeb,
	0x13, 0xe1, 0xc7, 0x45, 0x4c, 0xcc, 0xc1, 0x21, 0xc1, 0xab, 0x98, 0x24, 0x82, 0x21, 0x0e, 0x08,
	0x41, 0x76, 0x40, 0x70, 0x6a, 0x5e, 0xca, 0x29, 0x98, 0x54, 0x0c, 0xb2, 0x54, 0x0c, 0x48, 0xea,
	0x11, 0x93, 0x0a, 0x2e, 0xa9, 0x18, 0xf7, 0x00, 0x27, 0xdf, 0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92,
	0xc4, 0x57, 0x4c, 0x52, 0x10, 0x65, 0x56, 0x56, 0xc8, 0xea, 0xac, 0xac, 0x40, 0x0a, 0x93, 0xd8,
	0xc0, 0x01, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x24, 0x2c, 0x51, 0x70, 0x91, 0x01, 0x00,
	0x00,
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
