// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/orders/transactions/put/message.proto

package put

import (
	fmt "fmt"
	base "github.com/AssetMantle/schema/ids/base"
	base1 "github.com/AssetMantle/schema/types/base"
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

type Message struct {
	From         string           `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID       *base.IdentityID `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	MakerAssetID *base.AssetID    `protobuf:"bytes,3,opt,name=maker_asset_i_d,json=makerAssetID,proto3" json:"maker_asset_i_d,omitempty"`
	TakerAssetID *base.AssetID    `protobuf:"bytes,4,opt,name=taker_asset_i_d,json=takerAssetID,proto3" json:"taker_asset_i_d,omitempty"`
	MakerSplit   string           `protobuf:"bytes,5,opt,name=maker_split,json=makerSplit,proto3" json:"maker_split,omitempty"`
	TakerSplit   string           `protobuf:"bytes,6,opt,name=taker_split,json=takerSplit,proto3" json:"taker_split,omitempty"`
	ExpiryHeight *base1.Height    `protobuf:"bytes,7,opt,name=expiry_height,json=expiryHeight,proto3" json:"expiry_height,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c424bb34b9c88e68, []int{0}
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

func (m *Message) GetMakerAssetID() *base.AssetID {
	if m != nil {
		return m.MakerAssetID
	}
	return nil
}

func (m *Message) GetTakerAssetID() *base.AssetID {
	if m != nil {
		return m.TakerAssetID
	}
	return nil
}

func (m *Message) GetMakerSplit() string {
	if m != nil {
		return m.MakerSplit
	}
	return ""
}

func (m *Message) GetTakerSplit() string {
	if m != nil {
		return m.TakerSplit
	}
	return ""
}

func (m *Message) GetExpiryHeight() *base1.Height {
	if m != nil {
		return m.ExpiryHeight
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "AssetMantle.modules.x.orders.transactions.put.Message")
}

func init() {
	proto.RegisterFile("AssetMantle/modules/x/orders/transactions/put/message.proto", fileDescriptor_c424bb34b9c88e68)
}

var fileDescriptor_c424bb34b9c88e68 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x9b, 0xec, 0xda, 0xd5, 0xd9, 0x8a, 0x90, 0x53, 0xd8, 0x43, 0x5c, 0x44, 0x71, 0x11,
	0x76, 0x06, 0xd7, 0x8b, 0x8e, 0xa7, 0x96, 0x82, 0x56, 0x09, 0x1b, 0xea, 0x1e, 0x8a, 0x04, 0xc2,
	0xb4, 0x33, 0xb6, 0x83, 0x4d, 0x13, 0x32, 0x5f, 0xa0, 0x7d, 0x03, 0x8f, 0x3e, 0x83, 0x47, 0x9f,
	0x44, 0x3c, 0xed, 0xd1, 0xa3, 0xb4, 0x37, 0x2f, 0xbe, 0x82, 0x64, 0x66, 0x60, 0x67, 0x65, 0x59,
	0xc8, 0x29, 0x99, 0x99, 0xdf, 0xff, 0x9f, 0xff, 0x97, 0xef, 0x1b, 0xf4, 0xba, 0xaf, 0x94, 0x80,
	0x98, 0xad, 0x60, 0x29, 0x48, 0x5e, 0xf0, 0x7a, 0x29, 0x14, 0x59, 0x93, 0xa2, 0xe2, 0xa2, 0x52,
	0x04, 0x2a, 0xb6, 0x52, 0x6c, 0x06, 0xb2, 0x58, 0x29, 0x52, 0xd6, 0x40, 0x72, 0xa1, 0x14, 0x9b,
	0x0b, 0x5c, 0x56, 0x05, 0x14, 0xc1, 0xa9, 0x23, 0xc6, 0x56, 0x8c, 0xd7, 0xd8, 0x88, 0xb1, 0x2b,
	0xc6, 0x65, 0x0d, 0x47, 0xcf, 0xdc, 0x6f, 0xa9, 0xd9, 0x42, 0xe4, 0x8c, 0x48, 0xae, 0xc8, 0x94,
	0x29, 0x41, 0x58, 0x73, 0x96, 0x49, 0x6e, 0xac, 0x8f, 0x4e, 0x6f, 0x63, 0x25, 0x17, 0x2b, 0x90,
	0xb0, 0xb9, 0xc2, 0x6f, 0xb2, 0x86, 0x4d, 0x29, 0xac, 0x60, 0x21, 0xe4, 0x7c, 0x01, 0x86, 0x7d,
	0xf4, 0x65, 0x0f, 0x1d, 0xc4, 0xa6, 0x8e, 0x20, 0x40, 0xfb, 0x9f, 0xaa, 0x22, 0x0f, 0xbd, 0x63,
	0xef, 0xe4, 0xde, 0x58, 0xbf, 0x07, 0x7d, 0x74, 0xb7, 0x79, 0x66, 0x32, 0xe3, 0xa1, 0x7f, 0xec,
	0x9d, 0x1c, 0x9e, 0x3d, 0xc5, 0x6e, 0xa1, 0xc6, 0x1e, 0x4b, 0xae, 0x70, 0x63, 0x8e, 0x47, 0x36,
	0xcd, 0x68, 0x38, 0xee, 0x36, 0xc2, 0xd1, 0x30, 0x78, 0x8f, 0x1e, 0xe4, 0xec, 0xb3, 0xa8, 0x32,
	0x5b, 0x55, 0xc6, 0xc3, 0x3d, 0xed, 0xf4, 0xf8, 0x56, 0x27, 0x7d, 0x36, 0x1a, 0x8e, 0x7b, 0x5a,
	0x6c, 0x57, 0x8d, 0x19, 0xfc, 0x67, 0xb6, 0xdf, 0xc6, 0x0c, 0x5c, 0xb3, 0x87, 0xe8, 0xd0, 0x24,
	0x53, 0xe5, 0x52, 0x42, 0x78, 0x47, 0xd7, 0x8d, 0xf4, 0xd6, 0x87, 0x66, 0xa7, 0x01, 0xc0, 0x01,
	0xba, 0x06, 0x80, 0x2b, 0xe0, 0x1d, 0xba, 0x2f, 0xd6, 0xa5, 0xac, 0x36, 0x99, 0xf9, 0xab, 0xe1,
	0x81, 0x0e, 0xf3, 0xe4, 0xa6, 0x30, 0xba, 0x05, 0x26, 0xce, 0x5b, 0x0d, 0x8f, 0x7b, 0x46, 0x6b,
	0x56, 0x83, 0xbf, 0xfe, 0x8f, 0x6d, 0xe4, 0x5d, 0x6e, 0x23, 0xef, 0xf7, 0x36, 0xf2, 0xbe, 0xee,
	0xa2, 0xce, 0xe5, 0x2e, 0xea, 0xfc, 0xda, 0x45, 0x1d, 0xf4, 0x7c, 0x56, 0xe4, 0xb8, 0xd5, 0x7c,
	0x0d, 0x7a, 0xb6, 0xab, 0x49, 0xd3, 0xe6, 0xc4, 0xfb, 0xf8, 0x72, 0x2e, 0x61, 0x51, 0x4f, 0xf1,
	0xac, 0xc8, 0x49, 0xab, 0x31, 0xff, 0xe6, 0x77, 0xfb, 0xf1, 0xe4, 0xfc, 0x22, 0xf9, 0xee, 0x5f,
	0x1b, 0xf1, 0xd8, 0x46, 0x98, 0xe0, 0x73, 0x13, 0xe1, 0xc2, 0x8d, 0x90, 0xd4, 0xf0, 0xf3, 0x1a,
	0x9f, 0x5a, 0x3e, 0x9d, 0xa4, 0x86, 0x4f, 0x5d, 0x3e, 0x4d, 0x6a, 0xd8, 0xfa, 0xaf, 0x5a, 0xf1,
	0xe9, 0x9b, 0x64, 0x10, 0x0b, 0x60, 0x9c, 0x01, 0xfb, 0xe3, 0x9f, 0x39, 0x5a, 0x4a, 0xad, 0x98,
	0xd2, 0x09, 0xa5, 0x46, 0x4e, 0xa9, 0xab, 0xa7, 0x34, 0xa9, 0x61, 0xda, 0xd5, 0x77, 0xe0, 0xc5,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x21, 0xd2, 0x75, 0xf8, 0x03, 0x00, 0x00,
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
	if m.ExpiryHeight != nil {
		{
			size, err := m.ExpiryHeight.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TakerSplit) > 0 {
		i -= len(m.TakerSplit)
		copy(dAtA[i:], m.TakerSplit)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.TakerSplit)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MakerSplit) > 0 {
		i -= len(m.MakerSplit)
		copy(dAtA[i:], m.MakerSplit)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.MakerSplit)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TakerAssetID != nil {
		{
			size, err := m.TakerAssetID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.MakerAssetID != nil {
		{
			size, err := m.MakerAssetID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.MakerAssetID != nil {
		l = m.MakerAssetID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.TakerAssetID != nil {
		l = m.TakerAssetID.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.MakerSplit)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.TakerSplit)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ExpiryHeight != nil {
		l = m.ExpiryHeight.Size()
		n += 1 + l + sovMessage(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field MakerAssetID", wireType)
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
			if m.MakerAssetID == nil {
				m.MakerAssetID = &base.AssetID{}
			}
			if err := m.MakerAssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerAssetID", wireType)
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
			if m.TakerAssetID == nil {
				m.TakerAssetID = &base.AssetID{}
			}
			if err := m.TakerAssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerSplit", wireType)
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
			m.MakerSplit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerSplit", wireType)
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
			m.TakerSplit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryHeight", wireType)
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
			if m.ExpiryHeight == nil {
				m.ExpiryHeight = &base1.Height{}
			}
			if err := m.ExpiryHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
