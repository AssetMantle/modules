// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/transactions/get/transaction_response.proto

package get

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
	return fileDescriptor_93ba844f41a7df86, []int{0}
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
	proto.RegisterType((*TransactionResponse)(nil), "assetmantle.modules.orders.transactions.get.TransactionResponse")
}

func init() {
	proto.RegisterFile("orders/transactions/get/transaction_response.proto", fileDescriptor_93ba844f41a7df86)
}

var fileDescriptor_93ba844f41a7df86 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0xdb, 0x80, 0x0e, 0x1d, 0x15, 0xc1, 0x29, 0x83, 0xeb, 0x41, 0x02, 0x3a, 0x28, 0xd9,
	0xda, 0x25, 0x53, 0xb9, 0x22, 0x9d, 0x24, 0x20, 0xb9, 0xf6, 0xa3, 0x1e, 0x5c, 0x9b, 0x23, 0xf9,
	0x0e, 0xfc, 0x19, 0xfe, 0x06, 0x47, 0x7f, 0x89, 0x38, 0xdd, 0xe8, 0x28, 0xed, 0xe6, 0xe2, 0x5f,
	0x90, 0xbb, 0x04, 0x8c, 0xa0, 0xc3, 0xad, 0x1f, 0xef, 0xfb, 0xbc, 0x0f, 0x49, 0x76, 0x69, 0x6c,
	0x0b, 0xd6, 0x71, 0xb4, 0x7a, 0x70, 0xba, 0xc1, 0xa5, 0x19, 0x1c, 0xef, 0x00, 0xe3, 0xc3, 0xbd,
	0x05, 0xb7, 0x36, 0x83, 0x03, 0xb6, 0xb6, 0x06, 0xcd, 0xc9, 0x4c, 0x3b, 0x07, 0xd8, 0xeb, 0x01,
	0x57, 0xc0, 0x7a, 0xd3, 0x6e, 0x56, 0xe0, 0x98, 0xe7, 0xb0, 0x98, 0xc3, 0x3a, 0xc0, 0x8b, 0xb3,
	0xec, 0xb4, 0xfe, 0xb9, 0xdd, 0x06, 0x52, 0xf1, 0x45, 0x5e, 0x47, 0x9a, 0x6e, 0x47, 0x9a, 0x7e,
	0x8c, 0x34, 0x7d, 0x9a, 0x68, 0xb2, 0x9d, 0x68, 0xf2, 0x3e, 0xd1, 0x24, 0xe3, 0x8d, 0xe9, 0xd9,
	0x01, 0x13, 0xc5, 0xf9, 0x1f, 0x03, 0xd5, 0xce, 0xb4, 0x4a, 0xef, 0x6e, 0xba, 0x25, 0x3e, 0x6c,
	0x16, 0xac, 0x31, 0x3d, 0xcf, 0x77, 0xcc, 0x72, 0xcf, 0xe4, 0x81, 0xc9, 0x1f, 0xf9, 0x3f, 0x0f,
	0xf0, 0x4c, 0x8e, 0xf2, 0x72, 0x5e, 0xcb, 0x17, 0x32, 0xcb, 0x23, 0x97, 0x32, 0xb8, 0xcc, 0xbd,
	0x4b, 0x1d, 0xbb, 0x48, 0xc0, 0xb7, 0x5f, 0x69, 0x15, 0xd2, 0xca, 0xa7, 0x55, 0x9c, 0x56, 0x12,
	0x70, 0x24, 0xd7, 0x07, 0xa4, 0x95, 0xac, 0x8a, 0x12, 0x50, 0xb7, 0x1a, 0xf5, 0x27, 0xe1, 0x51,
	0x53, 0x88, 0x50, 0x15, 0xc2, 0x77, 0x85, 0x88, 0xcb, 0x42, 0x48, 0xc0, 0xc5, 0xf1, 0xfe, 0xf3,
	0xae, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xc2, 0xff, 0xaa, 0xf2, 0x01, 0x00, 0x00,
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
