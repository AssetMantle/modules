package helpers

import (
	"github.com/gogo/protobuf/proto"

	"github.com/AssetMantle/modules/schema/parameters"
)

type GenesisState interface {
	proto.Message
	Default() GenesisState
	GetMappables() []Mappable
	GetParameters() []parameters.Parameter
	Initialize(mappableList []Mappable, parameterList []parameters.Parameter) GenesisState
}
