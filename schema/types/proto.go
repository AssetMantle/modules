package types

import "github.com/gogo/protobuf/proto"

type Proto interface {
	proto.Message
	Size() int
	MarshalTo(dAtA []byte) (int, error)
	Unmarshal(dAtA []byte) error
}
