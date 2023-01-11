package helpers

import (
	"github.com/gogo/protobuf/proto"
)

type GenesisState interface {
	proto.Message
	Default() GenesisState
	GetMappables() []Mappable
	GetParameters() []Parameter
	Initialize(mappableList []Mappable, parameterList []Parameter) GenesisState
}
