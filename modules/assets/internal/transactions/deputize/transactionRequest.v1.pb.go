// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/deputize/transactionRequest.v1.proto

package deputize

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
	FromId               string `protobuf:"bytes,1,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	ToId                 string `protobuf:"bytes,2,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	ClassificationId     string `protobuf:"bytes,3,opt,name=classification_id,json=classificationId,proto3" json:"classification_id,omitempty"`
	MaintainedProperties string `protobuf:"bytes,4,opt,name=maintained_properties,json=maintainedProperties,proto3" json:"maintained_properties,omitempty"`
	CanMintAsset         bool   `protobuf:"varint,5,opt,name=can_mint_asset,json=canMintAsset,proto3" json:"can_mint_asset,omitempty"`
	CanBurnAsset         bool   `protobuf:"varint,6,opt,name=can_burn_asset,json=canBurnAsset,proto3" json:"can_burn_asset,omitempty"`
	CanRenumerateAsset   bool   `protobuf:"varint,7,opt,name=can_renumerate_asset,json=canRenumerateAsset,proto3" json:"can_renumerate_asset,omitempty"`
	CanAddMaintainer     bool   `protobuf:"varint,8,opt,name=can_add_maintainer,json=canAddMaintainer,proto3" json:"can_add_maintainer,omitempty"`
	CanRemoveMaintainer  bool   `protobuf:"varint,9,opt,name=can_remove_maintainer,json=canRemoveMaintainer,proto3" json:"can_remove_maintainer,omitempty"`
	CanMutateMaintainer  bool   `protobuf:"varint,10,opt,name=can_mutate_maintainer,json=canMutateMaintainer,proto3" json:"can_mutate_maintainer,omitempty"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba7326f9ab7bb597, []int{0}
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
	proto.RegisterType((*TransactionRequest)(nil), "deputize.TransactionRequest")
}

func init() {
	proto.RegisterFile("modules/assets/internal/transactions/deputize/transactionRequest.v1.proto", fileDescriptor_ba7326f9ab7bb597)
}

var fileDescriptor_ba7326f9ab7bb597 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0xb3, 0x49, 0x9a, 0xa6, 0x43, 0x91, 0x3a, 0x4d, 0x31, 0x7a, 0x58, 0x8b, 0x78, 0x28,
	0x28, 0x59, 0x6b, 0x6f, 0xde, 0x1a, 0x0a, 0x12, 0x70, 0x21, 0x04, 0x91, 0x22, 0xc2, 0xf2, 0x65,
	0x66, 0x1a, 0x07, 0x32, 0x33, 0x71, 0xe6, 0x9b, 0x1e, 0x7c, 0x02, 0x8f, 0x3e, 0x82, 0x78, 0xf4,
	0x21, 0x3c, 0x8b, 0x17, 0x7b, 0xf4, 0x28, 0xc9, 0xcd, 0xa7, 0x90, 0x99, 0x64, 0x37, 0x91, 0x9c,
	0x7a, 0x9b, 0xfd, 0xff, 0x7f, 0xbf, 0xd9, 0xd9, 0x9d, 0x8f, 0x0c, 0x94, 0xe1, 0x7e, 0x2a, 0x5c,
	0x06, 0xce, 0x09, 0x74, 0x99, 0xd4, 0x28, 0xac, 0x86, 0x69, 0x86, 0x16, 0xb4, 0x03, 0x86, 0xd2,
	0x68, 0x97, 0x71, 0x31, 0xf3, 0x28, 0x3f, 0x8a, 0xcd, 0x74, 0x24, 0x3e, 0x78, 0xe1, 0xb0, 0x77,
	0x7d, 0xda, 0x9b, 0x59, 0x83, 0x86, 0xb6, 0x4b, 0xea, 0x41, 0x67, 0x62, 0x26, 0x26, 0x86, 0x59,
	0x58, 0x2d, 0xfb, 0x47, 0xdf, 0x1b, 0x84, 0xbe, 0xde, 0xf2, 0xe9, 0x3d, 0xb2, 0x7b, 0x65, 0x8d,
	0x2a, 0x24, 0xef, 0x26, 0xc7, 0xc9, 0xc9, 0xde, 0xa8, 0x15, 0x1e, 0x07, 0x9c, 0x1e, 0x92, 0x1d,
	0x34, 0x21, 0xae, 0xc7, 0xb8, 0x89, 0x66, 0xc0, 0xe9, 0x13, 0x72, 0x97, 0x4d, 0xc1, 0x39, 0x79,
	0x25, 0x19, 0x84, 0x6d, 0x02, 0xd0, 0x88, 0xc0, 0xc1, 0xff, 0xc5, 0x80, 0xd3, 0x33, 0x72, 0xa4,
	0x40, 0x6a, 0x04, 0xa9, 0x05, 0x2f, 0x66, 0xd6, 0xcc, 0x84, 0x45, 0x29, 0x5c, 0xb7, 0x19, 0x85,
	0xce, 0xba, 0x1c, 0x56, 0x1d, 0x7d, 0x4c, 0xee, 0x30, 0xd0, 0x85, 0x92, 0x1a, 0x8b, 0xf8, 0x53,
	0xba, 0x3b, 0xc7, 0xc9, 0x49, 0x7b, 0xb4, 0xcf, 0x40, 0xe7, 0x52, 0xe3, 0x79, 0xc8, 0x4a, 0x6a,
	0xec, 0xad, 0x5e, 0x51, 0xad, 0x8a, 0xea, 0x7b, 0xab, 0x97, 0xd4, 0x33, 0xd2, 0x09, 0x94, 0x15,
	0xda, 0x2b, 0x61, 0x01, 0xc5, 0x8a, 0xdd, 0x8d, 0x2c, 0x65, 0xa0, 0x47, 0x55, 0xb5, 0x34, 0x9e,
	0x92, 0x90, 0x16, 0xc0, 0x79, 0x51, 0x9d, 0xce, 0x76, 0xdb, 0x91, 0x3f, 0x60, 0xa0, 0xcf, 0x39,
	0xcf, 0xab, 0x9c, 0x3e, 0x27, 0x47, 0xcb, 0xfd, 0x95, 0xb9, 0x16, 0x9b, 0xc2, 0x5e, 0x14, 0x0e,
	0xe3, 0x0b, 0x42, 0xb7, 0xed, 0x28, 0x8f, 0xe1, 0x3c, 0x1b, 0x0e, 0xa9, 0x9c, 0x3c, 0x76, 0x6b,
	0xe7, 0x45, 0xf3, 0xd3, 0x97, 0x87, 0xb5, 0xfe, 0xaf, 0xe4, 0xc7, 0x3c, 0x4d, 0x6e, 0xe6, 0x69,
	0xf2, 0x67, 0x9e, 0x26, 0x9f, 0x17, 0x69, 0xed, 0x66, 0x91, 0xd6, 0x7e, 0x2f, 0xd2, 0x1a, 0xd9,
	0x67, 0x46, 0xf5, 0xca, 0xfb, 0xef, 0xdf, 0xdf, 0xbe, 0xe6, 0x37, 0xa7, 0xc3, 0x30, 0x04, 0xc3,
	0xe4, 0xed, 0xab, 0x89, 0xc4, 0xf7, 0x7e, 0xdc, 0x63, 0x46, 0x65, 0xf1, 0x9b, 0x73, 0xd0, 0x38,
	0x15, 0x59, 0x39, 0x88, 0xb7, 0x1a, 0xc8, 0xaf, 0xf5, 0xc6, 0xc5, 0xe5, 0xe5, 0xb7, 0x7a, 0xfb,
	0x62, 0x15, 0xfc, 0x5c, 0x2f, 0xe7, 0xf5, 0x4e, 0xb9, 0x7c, 0xf7, 0x72, 0xd8, 0xcf, 0x05, 0x02,
	0x07, 0x84, 0xbf, 0x6b, 0x62, 0xdc, 0x8a, 0x93, 0x79, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x40,
	0x6e, 0x50, 0x6f, 0x06, 0x03, 0x00, 0x00,
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
	if m.CanMutateMaintainer {
		i--
		if m.CanMutateMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.CanRemoveMaintainer {
		i--
		if m.CanRemoveMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.CanAddMaintainer {
		i--
		if m.CanAddMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.CanRenumerateAsset {
		i--
		if m.CanRenumerateAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.CanBurnAsset {
		i--
		if m.CanBurnAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.CanMintAsset {
		i--
		if m.CanMintAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.MaintainedProperties) > 0 {
		i -= len(m.MaintainedProperties)
		copy(dAtA[i:], m.MaintainedProperties)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.MaintainedProperties)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClassificationId) > 0 {
		i -= len(m.ClassificationId)
		copy(dAtA[i:], m.ClassificationId)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.ClassificationId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ToId) > 0 {
		i -= len(m.ToId)
		copy(dAtA[i:], m.ToId)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.ToId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromId) > 0 {
		i -= len(m.FromId)
		copy(dAtA[i:], m.FromId)
		i = encodeVarintTransactionRequestV1(dAtA, i, uint64(len(m.FromId)))
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
	l = len(m.FromId)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.ToId)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.ClassificationId)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	l = len(m.MaintainedProperties)
	if l > 0 {
		n += 1 + l + sovTransactionRequestV1(uint64(l))
	}
	if m.CanMintAsset {
		n += 2
	}
	if m.CanBurnAsset {
		n += 2
	}
	if m.CanRenumerateAsset {
		n += 2
	}
	if m.CanAddMaintainer {
		n += 2
	}
	if m.CanRemoveMaintainer {
		n += 2
	}
	if m.CanMutateMaintainer {
		n += 2
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
				return fmt.Errorf("proto: wrong wireType = %d for field FromId", wireType)
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
			m.FromId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToId", wireType)
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
			m.ToId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationId", wireType)
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
			m.ClassificationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintainedProperties", wireType)
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
			m.MaintainedProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanMintAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanMintAsset = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanBurnAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanBurnAsset = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRenumerateAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanRenumerateAsset = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanAddMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanAddMaintainer = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRemoveMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanRemoveMaintainer = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanMutateMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionRequestV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanMutateMaintainer = bool(v != 0)
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
