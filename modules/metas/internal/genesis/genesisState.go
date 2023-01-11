package genesis

import (
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/modules/metas/internal/parameters"
	"github.com/AssetMantle/modules/schema/helpers"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
)

var _ helpers.GenesisState = (*GenesisState)(nil)

func (genesisState *GenesisState) Default() helpers.GenesisState {
	return PrototypeGenesisState()
}
func (genesisState *GenesisState) GetMappables() []helpers.Mappable {
	return mappable.MappablesToInterface(genesisState.Mappables)
}
func (genesisState *GenesisState) GetParameters() []helpers.Parameter {
	return baseParameters.ParametersToInterfaces(genesisState.Parameters)
}
func (genesisState *GenesisState) Initialize(mappableList []helpers.Mappable, parameterList []helpers.Parameter) helpers.GenesisState {
	if len(mappableList) == 0 {
		genesisState.Mappables = mappable.MappablesFromInterface(genesisState.Default().GetMappables())
	} else {
		genesisState.Mappables = mappable.MappablesFromInterface(mappableList)
	}

	if len(parameterList) == 0 {
		genesisState.Parameters = baseParameters.ParametersFromInterfaces(genesisState.Default().GetParameters())
	} else {
		for _, defaultParameter := range genesisState.Default().GetParameters() {
			for i, parameter := range parameterList {
				if defaultParameter.GetID().Compare(parameter.GetID()) == 0 {
					parameterList[i] = defaultParameter.Mutate(parameter.GetData())
				}
			}
		}
		genesisState.Parameters = baseParameters.ParametersFromInterfaces(parameterList)
	}

	return genesisState
}

func PrototypeGenesisState() helpers.GenesisState {
	return &GenesisState{
		Mappables:  []*mappable.Mappable{},
		Parameters: baseParameters.ParametersFromInterfaces(parameters.Prototype().GetList()),
	}
}
