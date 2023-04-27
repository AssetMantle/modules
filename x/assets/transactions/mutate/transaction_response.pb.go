// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/transactions/mutate/transaction_response.proto

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
	return fileDescriptor_44d38d2bba1859b6, []int{0}
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
	proto.RegisterType((*TransactionResponse)(nil), "assetmantle.modules.assets.transactions.mutate.TransactionResponse")
}

func init() {
	proto.RegisterFile("assets/transactions/mutate/transaction_response.proto", fileDescriptor_44d38d2bba1859b6)
}

var fileDescriptor_44d38d2bba1859b6 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0xc7, 0xdb, 0x8a, 0x0e, 0x1d, 0x15, 0xc1, 0x29, 0x83, 0x0f, 0x90, 0xc0, 0x89, 0x4b, 0x9c,
	0xd2, 0xc5, 0x29, 0x50, 0xa4, 0x93, 0x04, 0x24, 0xd7, 0x0b, 0x7a, 0x70, 0x69, 0x8e, 0x7e, 0x5f,
	0xc1, 0xd5, 0x37, 0xf0, 0x19, 0x1c, 0x7d, 0x12, 0x71, 0xba, 0xd1, 0x51, 0xda, 0xcd, 0xa7, 0x90,
	0x4b, 0x02, 0x46, 0xd0, 0xa1, 0xeb, 0x1f, 0x7e, 0xbf, 0xef, 0x47, 0x52, 0x5e, 0x6a, 0x00, 0x83,
	0xc0, 0xb0, 0xd7, 0x1d, 0xe8, 0x16, 0xd7, 0xae, 0x03, 0x66, 0x07, 0xd4, 0x68, 0xd2, 0xed, 0xae,
	0x37, 0xb0, 0x75, 0x1d, 0x18, 0xba, 0xed, 0x1d, 0xba, 0x63, 0xea, 0x31, 0xab, 0x3b, 0xdc, 0x18,
	0x6a, 0xdd, 0x6a, 0xd8, 0x18, 0x08, 0x1b, 0xd0, 0x54, 0x45, 0x83, 0xea, 0xfc, 0xb4, 0x3c, 0x69,
	0x7e, 0xe6, 0x9b, 0x28, 0xab, 0x9e, 0x0e, 0xde, 0x46, 0x92, 0xef, 0x46, 0x92, 0x7f, 0x8e, 0x24,
	0x7f, 0x9e, 0x48, 0xb6, 0x9b, 0x48, 0xf6, 0x31, 0x91, 0xac, 0x5c, 0xb4, 0xce, 0xce, 0xbc, 0x52,
	0x9d, 0xfd, 0x71, 0xa3, 0xde, 0xf7, 0xd6, 0xf9, 0xed, 0xd5, 0xfd, 0x1a, 0x1f, 0x86, 0x25, 0x6d,
	0x9d, 0x65, 0x62, 0xaf, 0x90, 0x5e, 0xcb, 0xa2, 0x96, 0x3d, 0xb2, 0xff, 0x5f, 0xe2, 0xa5, 0x38,
	0x14, 0x52, 0x34, 0xf2, 0xb5, 0xa0, 0x22, 0x29, 0x92, 0xb1, 0x48, 0x84, 0xa2, 0x26, 0x2d, 0x92,
	0x1e, 0x7c, 0xff, 0x05, 0xa8, 0x08, 0xa8, 0x00, 0xa8, 0x14, 0x50, 0x01, 0x18, 0x0b, 0x3e, 0x0f,
	0x50, 0xd7, 0x75, 0x25, 0x0d, 0xea, 0x95, 0x46, 0xfd, 0x55, 0x2c, 0x12, 0x98, 0xf3, 0x48, 0xf3,
	0xa0, 0x04, 0xce, 0x53, 0x9e, 0xf3, 0x20, 0x58, 0x1e, 0xf9, 0x1f, 0xbd, 0xf8, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x29, 0xe6, 0xce, 0x8b, 0x0a, 0x02, 0x00, 0x00,
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