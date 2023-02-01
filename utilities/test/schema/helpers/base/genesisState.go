package base

import (
	"github.com/AssetMantle/modules/schema/helpers"
	parametersSchema "github.com/AssetMantle/modules/schema/parameters"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
)

var _ helpers.Genesis = (*GenesisState)(nil)

func (genesisState *GenesisState) Default() helpers.Genesis {
	return PrototypeGenesisState()
}
func (genesisState *GenesisState) GetMappables() []helpers.Mappable {
	return MappablesToInterface(genesisState.Mappables)
}
func (genesisState *GenesisState) GetParameters() []parametersSchema.Parameter {
	return baseParameters.ParametersToInterfaces(genesisState.Parameters)
}
func (genesisState *GenesisState) Initialize(mappableList []helpers.Mappable, parameterList []parametersSchema.Parameter) helpers.Genesis {
	if len(mappableList) == 0 {
		genesisState.Mappables = MappablesFromInterface(genesisState.Default().GetMappables())
	} else {
		genesisState.Mappables = MappablesFromInterface(mappableList)
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

func PrototypeGenesisState() helpers.Genesis {
	return &GenesisState{
		Mappables:  []*TestMappable{},
		Parameters: nil,
	}
}
