// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/unprovision/transaction_response.proto

package unprovision

import (
	fmt "fmt"
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

type TransactionResponse struct {
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3ae37edba41ed3e, []int{0}
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
	proto.RegisterType((*TransactionResponse)(nil), "AssetMantle.modules.x.identities.transactions.unprovision.TransactionResponse")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/unprovision/transaction_response.proto", fileDescriptor_f3ae37edba41ed3e)
}

var fileDescriptor_f3ae37edba41ed3e = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0x07, 0xf0, 0x6b, 0x85, 0x1b, 0x3a, 0x2a, 0x82, 0x53, 0x06, 0x1f, 0x20, 0x19, 0x9c, 0x0c,
	0x28, 0x5c, 0x51, 0xe4, 0x86, 0x42, 0x91, 0x1e, 0x14, 0x09, 0x48, 0xae, 0x0d, 0x1a, 0xb8, 0x26,
	0xa5, 0xf9, 0x2a, 0xf7, 0x18, 0x3e, 0x83, 0xa3, 0xaf, 0xe1, 0x22, 0x4e, 0x37, 0x3a, 0x4a, 0xbb,
	0xf9, 0x14, 0xa2, 0x04, 0xf2, 0x1d, 0x38, 0x75, 0xfe, 0xf8, 0xff, 0xbe, 0x7f, 0x3e, 0x92, 0x14,
	0x0b, 0xe7, 0x14, 0x64, 0xd2, 0xc0, 0x46, 0xb1, 0xc6, 0xd6, 0xfd, 0x46, 0x39, 0xb6, 0x65, 0xba,
	0x56, 0x06, 0x34, 0x68, 0xe5, 0x18, 0x74, 0xd2, 0x38, 0x59, 0x81, 0xb6, 0xc6, 0xb1, 0xde, 0xb4,
	0x9d, 0x7d, 0xd2, 0x4e, 0x5b, 0x83, 0x07, 0xf7, 0x9d, 0x72, 0xad, 0x35, 0x4e, 0xd1, 0xb6, 0xb3,
	0x60, 0x0f, 0xcf, 0x91, 0x4a, 0xbd, 0x4a, 0xb7, 0x34, 0xa8, 0x14, 0xab, 0x14, 0xa9, 0xa7, 0xc7,
	0xc9, 0x51, 0x11, 0x66, 0xb7, 0xde, 0x4d, 0xdf, 0x0e, 0xde, 0x07, 0x12, 0xed, 0x06, 0x12, 0x7d,
	0x0d, 0x24, 0x7a, 0x1e, 0xc9, 0x6c, 0x37, 0x92, 0xd9, 0xe7, 0x48, 0x66, 0xc9, 0x45, 0x65, 0x1b,
	0x3a, 0x79, 0x61, 0x7a, 0xf2, 0xcf, 0xba, 0xfc, 0xf7, 0x15, 0x79, 0x74, 0x77, 0xf5, 0xa0, 0xe1,
	0xb1, 0x5f, 0xd3, 0xca, 0x36, 0x6c, 0xf2, 0xa1, 0x5e, 0xe2, 0xf9, 0x22, 0x2b, 0x97, 0xc5, 0xea,
	0x35, 0xde, 0xbb, 0x4b, 0xe6, 0x6b, 0x96, 0x74, 0x19, 0x6a, 0x16, 0xb8, 0xe6, 0x2a, 0x20, 0x1f,
	0x7b, 0x59, 0xe1, 0xb3, 0xa2, 0x14, 0x21, 0x2b, 0x70, 0x56, 0xa0, 0xec, 0x10, 0x5f, 0x4f, 0xce,
	0x8a, 0x9b, 0x3c, 0xcd, 0x14, 0xc8, 0x5a, 0x82, 0xfc, 0x8e, 0x2f, 0x91, 0xc3, 0xb9, 0x87, 0x38,
	0x2f, 0x39, 0x0f, 0x14, 0xe7, 0xd8, 0xe2, 0x1c, 0x61, 0xeb, 0xf9, 0xdf, 0xf7, 0x38, 0xfb, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xd3, 0x1c, 0xee, 0x76, 0x02, 0x00, 0x00,
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
