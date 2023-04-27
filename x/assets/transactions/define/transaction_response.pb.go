// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/transactions/define/transaction_response.proto

package define

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
	ClassificationID string `protobuf:"bytes,1,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0661ce7a7dc35f, []int{0}
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

func (m *TransactionResponse) GetClassificationID() string {
	if m != nil {
		return m.ClassificationID
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionResponse)(nil), "assetmantle.modules.assets.transactions.define.TransactionResponse")
}

func init() {
	proto.RegisterFile("assets/transactions/define/transaction_response.proto", fileDescriptor_1c0661ce7a7dc35f)
}

var fileDescriptor_1c0661ce7a7dc35f = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xf4, 0x30,
	0x1c, 0x87, 0x9b, 0xbe, 0xbc, 0x82, 0x9d, 0xa4, 0x2e, 0x37, 0x05, 0x71, 0x72, 0x90, 0x04, 0x4e,
	0x5c, 0xe2, 0xd4, 0x5a, 0x10, 0x87, 0x42, 0x39, 0x3a, 0x49, 0xe1, 0xc8, 0xb5, 0x39, 0x0d, 0xb4,
	0xcd, 0xd1, 0x7f, 0x0e, 0x5c, 0xfd, 0x06, 0x7e, 0x06, 0x47, 0x3f, 0x89, 0x38, 0xdd, 0xe8, 0x28,
	0xed, 0xe6, 0xa7, 0x90, 0x6b, 0x02, 0xe6, 0x40, 0x87, 0x5b, 0x7f, 0xe5, 0x79, 0xfa, 0x24, 0x09,
	0x2e, 0x39, 0x80, 0xd0, 0x40, 0x75, 0xc7, 0x5b, 0xe0, 0xa5, 0x96, 0xaa, 0x05, 0x5a, 0x89, 0xa5,
	0x6c, 0x85, 0xbb, 0xcd, 0x3b, 0x01, 0x2b, 0xd5, 0x82, 0x20, 0xab, 0x4e, 0x69, 0x15, 0x92, 0x11,
	0x6b, 0x78, 0xab, 0x6b, 0x41, 0x1a, 0x55, 0xad, 0x6b, 0x01, 0x66, 0x03, 0xe2, 0xaa, 0x88, 0x51,
	0x9d, 0x5e, 0x07, 0xc7, 0xf9, 0xcf, 0x3c, 0xb3, 0xb2, 0xf0, 0x3c, 0x08, 0xcb, 0x9a, 0x03, 0xc8,
	0xa5, 0x2c, 0xf9, 0xf8, 0x1f, 0x39, 0xaf, 0x26, 0xe8, 0x04, 0x9d, 0x1d, 0xce, 0x8e, 0x76, 0xbf,
	0xdc, 0x26, 0xf1, 0xd3, 0xbf, 0xb7, 0x1e, 0xa3, 0x4d, 0x8f, 0xd1, 0x67, 0x8f, 0xd1, 0xf3, 0x80,
	0xbd, 0xcd, 0x80, 0xbd, 0x8f, 0x01, 0x7b, 0xc1, 0xb4, 0x54, 0xcd, 0x9e, 0x4d, 0xf1, 0xe4, 0x97,
	0xa2, 0x6c, 0x7b, 0xba, 0x0c, 0xdd, 0x5d, 0xdd, 0x4b, 0xfd, 0xb0, 0x5e, 0x90, 0x52, 0x35, 0x34,
	0xda, 0x2a, 0xd2, 0x51, 0x4b, 0xad, 0x96, 0x3e, 0xd2, 0xbf, 0xef, 0xed, 0xc5, 0xff, 0x1f, 0xa5,
	0x51, 0x9e, 0xbc, 0xfa, 0x24, 0x72, 0x8a, 0x52, 0x5b, 0x14, 0x99, 0xa2, 0xdc, 0x2d, 0x4a, 0x46,
	0xf0, 0x7d, 0x07, 0x28, 0x2c, 0x50, 0x18, 0xa0, 0x70, 0x81, 0xc2, 0x00, 0xbd, 0xcf, 0xf6, 0x03,
	0x8a, 0x9b, 0x2c, 0x4e, 0x85, 0xe6, 0x15, 0xd7, 0xfc, 0xcb, 0x9f, 0x3a, 0x30, 0x63, 0x96, 0x66,
	0x46, 0x09, 0x8c, 0xb9, 0x3c, 0x63, 0x46, 0xb0, 0x38, 0x18, 0xdf, 0xff, 0xe2, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0xd5, 0x9b, 0x4d, 0x71, 0x38, 0x02, 0x00, 0x00,
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
	if len(m.ClassificationID) > 0 {
		i -= len(m.ClassificationID)
		copy(dAtA[i:], m.ClassificationID)
		i = encodeVarintTransactionResponse(dAtA, i, uint64(len(m.ClassificationID)))
		i--
		dAtA[i] = 0xa
	}
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
	l = len(m.ClassificationID)
	if l > 0 {
		n += 1 + l + sovTransactionResponse(uint64(l))
	}
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionResponse
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
				return ErrInvalidLengthTransactionResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassificationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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