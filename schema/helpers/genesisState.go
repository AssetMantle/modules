package helpers

import (
	"github.com/gogo/protobuf/proto"
)

type GenesisState interface {
	proto.Message
	Default() GenesisState
	Initialize(mappableList []Mappable, parameterList []Parameter) GenesisState
}
