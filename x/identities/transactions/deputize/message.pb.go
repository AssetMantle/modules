// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/transactions/deputize/message.proto

package deputize

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
	base1 "github.com/AssetMantle/schema/go/lists/base"
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

type Message struct {
	From                 string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID               *base.IdentityID       `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ToID                 *base.IdentityID       `protobuf:"bytes,3,opt,name=to_i_d,json=toID,proto3" json:"to_i_d,omitempty"`
	ClassificationID     *base.ClassificationID `protobuf:"bytes,4,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
	MaintainedProperties *base1.PropertyList    `protobuf:"bytes,5,opt,name=maintained_properties,json=maintainedProperties,proto3" json:"maintained_properties,omitempty"`
	CanMintAsset         bool                   `protobuf:"varint,6,opt,name=can_mint_asset,json=canMintAsset,proto3" json:"can_mint_asset,omitempty"`
	CanBurnAsset         bool                   `protobuf:"varint,7,opt,name=can_burn_asset,json=canBurnAsset,proto3" json:"can_burn_asset,omitempty"`
	CanRenumerateAsset   bool                   `protobuf:"varint,8,opt,name=can_renumerate_asset,json=canRenumerateAsset,proto3" json:"can_renumerate_asset,omitempty"`
	CanAddMaintainer     bool                   `protobuf:"varint,9,opt,name=can_add_maintainer,json=canAddMaintainer,proto3" json:"can_add_maintainer,omitempty"`
	CanRemoveMaintainer  bool                   `protobuf:"varint,10,opt,name=can_remove_maintainer,json=canRemoveMaintainer,proto3" json:"can_remove_maintainer,omitempty"`
	CanMutateMaintainer  bool                   `protobuf:"varint,11,opt,name=can_mutate_maintainer,json=canMutateMaintainer,proto3" json:"can_mutate_maintainer,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_251909eb88314c37, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetFromID() *base.IdentityID {
	if m != nil {
		return m.FromID
	}
	return nil
}

func (m *Message) GetToID() *base.IdentityID {
	if m != nil {
		return m.ToID
	}
	return nil
}

func (m *Message) GetClassificationID() *base.ClassificationID {
	if m != nil {
		return m.ClassificationID
	}
	return nil
}

func (m *Message) GetMaintainedProperties() *base1.PropertyList {
	if m != nil {
		return m.MaintainedProperties
	}
	return nil
}

func (m *Message) GetCanMintAsset() bool {
	if m != nil {
		return m.CanMintAsset
	}
	return false
}

func (m *Message) GetCanBurnAsset() bool {
	if m != nil {
		return m.CanBurnAsset
	}
	return false
}

func (m *Message) GetCanRenumerateAsset() bool {
	if m != nil {
		return m.CanRenumerateAsset
	}
	return false
}

func (m *Message) GetCanAddMaintainer() bool {
	if m != nil {
		return m.CanAddMaintainer
	}
	return false
}

func (m *Message) GetCanRemoveMaintainer() bool {
	if m != nil {
		return m.CanRemoveMaintainer
	}
	return false
}

func (m *Message) GetCanMutateMaintainer() bool {
	if m != nil {
		return m.CanMutateMaintainer
	}
	return false
}

func init() {
	proto.RegisterType((*Message)(nil), "assetmantle.modules.identities.transactions.deputize.Message")
}

func init() {
	proto.RegisterFile("identities/transactions/deputize/message.proto", fileDescriptor_251909eb88314c37)
}

var fileDescriptor_251909eb88314c37 = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6b, 0xd4, 0x4c,
	0x18, 0xc7, 0x9b, 0xed, 0x76, 0xbb, 0x9d, 0x96, 0x97, 0x92, 0xb7, 0x85, 0xb0, 0x87, 0xb0, 0x88,
	0xe0, 0x82, 0x9a, 0x48, 0xed, 0xa1, 0x04, 0x04, 0xb3, 0x06, 0x24, 0x60, 0x60, 0x09, 0x3d, 0xe9,
	0x42, 0x9c, 0xcd, 0x4c, 0xdb, 0x81, 0x9d, 0x99, 0x25, 0x33, 0x11, 0xeb, 0xa7, 0xf0, 0xea, 0xd5,
	0xa3, 0x9f, 0x44, 0x3c, 0xf5, 0xe8, 0x51, 0x76, 0x6f, 0x1e, 0xfd, 0x04, 0x32, 0x93, 0xd9, 0x24,
	0xad, 0xa2, 0xd8, 0xd3, 0x2e, 0xcf, 0xf3, 0xff, 0xfd, 0xe6, 0xc9, 0x43, 0x26, 0xc0, 0x23, 0x08,
	0x33, 0x49, 0x24, 0xc1, 0xc2, 0x97, 0x05, 0x64, 0x02, 0xe6, 0x92, 0x70, 0x26, 0x7c, 0x84, 0x17,
	0xa5, 0x24, 0xef, 0xb0, 0x4f, 0xb1, 0x10, 0xf0, 0x1c, 0x7b, 0x8b, 0x82, 0x4b, 0x6e, 0x1f, 0x43,
	0x21, 0xb0, 0xa4, 0x90, 0xc9, 0x39, 0xf6, 0x28, 0x47, 0xe5, 0x1c, 0x8b, 0x96, 0xc3, 0x6b, 0x3b,
	0xbc, 0xb5, 0x63, 0x30, 0x24, 0x48, 0xf8, 0x33, 0x28, 0xb0, 0x9f, 0xcf, 0xa1, 0x10, 0xe4, 0x8c,
	0xe4, 0x50, 0x45, 0x32, 0x82, 0x2a, 0xef, 0x60, 0x50, 0x27, 0x8c, 0xec, 0xb2, 0xe9, 0xb9, 0x73,
	0x22, 0xa4, 0xe9, 0x2e, 0x0a, 0xbe, 0xc0, 0x85, 0xbc, 0xcc, 0x54, 0xad, 0xea, 0xdf, 0xf9, 0xd1,
	0x05, 0xdb, 0x49, 0x35, 0xa5, 0x6d, 0x83, 0xee, 0x59, 0xc1, 0xa9, 0x63, 0x0d, 0xad, 0xd1, 0x4e,
	0xaa, 0xff, 0xdb, 0x21, 0xe8, 0xab, 0xdf, 0x8c, 0x64, 0xc8, 0xe9, 0x0c, 0xad, 0xd1, 0xee, 0xd1,
	0x3d, 0xaf, 0xfd, 0x18, 0x22, 0xbf, 0xc0, 0x14, 0x7a, 0x04, 0x09, 0x4f, 0x9d, 0xe1, 0xc5, 0x66,
	0x82, 0x38, 0x4a, 0x7b, 0x0a, 0x8c, 0x23, 0xfb, 0x09, 0xe8, 0x49, 0xae, 0x05, 0x9b, 0xff, 0x26,
	0xe8, 0x4a, 0x1e, 0x47, 0xf6, 0x2b, 0x60, 0xdf, 0x7c, 0xf0, 0x0c, 0x39, 0x5d, 0xad, 0x7a, 0xf8,
	0x47, 0xd5, 0xb3, 0x6b, 0x58, 0x1c, 0xa5, 0xfb, 0xf9, 0x8d, 0x8a, 0xfd, 0x1a, 0x1c, 0x52, 0x48,
	0x98, 0x84, 0x84, 0x61, 0x94, 0x99, 0x05, 0x11, 0x2c, 0x9c, 0x2d, 0xed, 0xbf, 0xff, 0x3b, 0xbf,
	0xde, 0x68, 0x75, 0xc2, 0xc4, 0x6c, 0xf4, 0x05, 0x11, 0x32, 0x3d, 0x68, 0x4c, 0x93, 0x5a, 0x64,
	0xdf, 0x05, 0xff, 0xe5, 0x90, 0x65, 0x94, 0x30, 0x99, 0x69, 0x99, 0xd3, 0x1b, 0x5a, 0xa3, 0x7e,
	0xba, 0x97, 0x43, 0x96, 0x10, 0x26, 0x43, 0x55, 0x5b, 0xa7, 0x66, 0x65, 0xc1, 0x4c, 0x6a, 0xbb,
	0x4e, 0x8d, 0xcb, 0x82, 0x55, 0xa9, 0x47, 0xe0, 0x40, 0xa5, 0x0a, 0xcc, 0x4a, 0x8a, 0x0b, 0x28,
	0xb1, 0xc9, 0xf6, 0x75, 0xd6, 0xce, 0x21, 0x4b, 0xeb, 0x56, 0x45, 0x3c, 0x00, 0xaa, 0x9a, 0x41,
	0x84, 0xb2, 0x7a, 0xba, 0xc2, 0xd9, 0xd1, 0xf9, 0xfd, 0x1c, 0xb2, 0x10, 0xa1, 0xa4, 0xae, 0xdb,
	0x47, 0xe0, 0xb0, 0xf2, 0x53, 0xfe, 0x06, 0xb7, 0x01, 0xa0, 0x81, 0xff, 0xf5, 0x01, 0xaa, 0xf7,
	0x2b, 0x43, 0x4b, 0xa9, 0xe6, 0x69, 0x31, 0xbb, 0x35, 0x93, 0xe8, 0x5e, 0xc3, 0x8c, 0x3f, 0x6c,
	0x7e, 0x5e, 0xba, 0xd6, 0xd5, 0xd2, 0xb5, 0xbe, 0x2d, 0x5d, 0xeb, 0xfd, 0xca, 0xdd, 0xb8, 0x5a,
	0xb9, 0x1b, 0x5f, 0x57, 0xee, 0x06, 0x38, 0xc9, 0x39, 0xf5, 0x6e, 0x73, 0x4f, 0xc6, 0x7b, 0xe6,
	0x35, 0x9e, 0xa8, 0xf7, 0x7a, 0x62, 0xbd, 0x0c, 0xcf, 0x89, 0xbc, 0x28, 0x67, 0x5e, 0xce, 0xa9,
	0xaf, 0x97, 0x91, 0x68, 0xa1, 0x6f, 0x84, 0xfe, 0x5b, 0xff, 0x6f, 0xd7, 0xf7, 0x63, 0x67, 0x2b,
	0x4c, 0xe2, 0xd3, 0xe8, 0x53, 0xe7, 0x38, 0x6c, 0x4d, 0x94, 0x98, 0x89, 0xe2, 0x66, 0xa2, 0xd3,
	0xf6, 0x44, 0x91, 0xc1, 0xbf, 0x5c, 0xc3, 0xa6, 0x06, 0x9b, 0x36, 0xd8, 0xb4, 0x8d, 0x4d, 0xd7,
	0xd8, 0xb2, 0xf3, 0xf4, 0x36, 0xd8, 0xf4, 0xf9, 0x64, 0x9c, 0x60, 0x09, 0x11, 0x94, 0xf0, 0x7b,
	0xe7, 0xa4, 0xa5, 0x08, 0x02, 0xe3, 0x08, 0x82, 0x46, 0x12, 0x04, 0x6d, 0x4b, 0x10, 0xac, 0x35,
	0xb3, 0x9e, 0xfe, 0x2e, 0x3c, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x5d, 0x2a, 0x47, 0xdd,
	0x04, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x58
	}
	if m.CanRemoveMaintainer {
		i--
		if m.CanRemoveMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.CanAddMaintainer {
		i--
		if m.CanAddMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.CanRenumerateAsset {
		i--
		if m.CanRenumerateAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.CanBurnAsset {
		i--
		if m.CanBurnAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.CanMintAsset {
		i--
		if m.CanMintAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.MaintainedProperties != nil {
		{
			size, err := m.MaintainedProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ToID != nil {
		{
			size, err := m.ToID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FromID != nil {
		{
			size, err := m.FromID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.FromID != nil {
		l = m.FromID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ToID != nil {
		l = m.ToID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.MaintainedProperties != nil {
		l = m.MaintainedProperties.Size()
		n += 1 + l + sovMessage(uint64(l))
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

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
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
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FromID == nil {
				m.FromID = &base.IdentityID{}
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ToID == nil {
				m.ToID = &base.IdentityID{}
			}
			if err := m.ToID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClassificationID == nil {
				m.ClassificationID = &base.ClassificationID{}
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintainedProperties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaintainedProperties == nil {
				m.MaintainedProperties = &base1.PropertyList{}
			}
			if err := m.MaintainedProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanMintAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanBurnAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRenumerateAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanAddMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRemoveMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanMutateMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)