package helpers

import (
	"github.com/gogo/protobuf/proto"
)

type ParameterList interface {
	proto.Message
	Get() []Parameter
}
