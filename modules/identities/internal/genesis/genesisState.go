package genesis

import (
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/parameters/base"
)

var _ helpers.GenesisState = (*GenesisState)(nil)

func (m *GenesisState) Default() helpers.GenesisState {
	return PrototypeGenesisState()
}
func (m *GenesisState) GetMappables() []helpers.Mappable {
	return mappable.MappablesToInterface(m.Mappables)
}
func (m *GenesisState) GetParameters() []helpers.Parameter {
	return base.ParametersToInterfaces(m.Parameters)
}
func (m *GenesisState) Initialize(mappableList []helpers.Mappable, parameterList []helpers.Parameter) helpers.GenesisState {
	if len(mappableList) == 0 {
		m.Mappables = mappable.MappablesFromInterface(m.Default().GetMappables())
	} else {
		m.Mappables = mappable.MappablesFromInterface(mappableList)
	}

	if len(parameterList) == 0 {
		m.Parameters = base.ParametersFromInterfaces(m.Default().GetParameters())
	} else {
		for _, defaultParameter := range m.Default().GetParameters() {
			for i, parameter := range parameterList {
				if defaultParameter.GetID().Compare(parameter.GetID()) == 0 {
					parameterList[i] = defaultParameter.Mutate(parameter.GetData())
				}
			}
		}
		m.Parameters = base.ParametersFromInterfaces(parameterList)
	}

	return m
}

func PrototypeGenesisState() helpers.GenesisState {
	return &GenesisState{
		Mappables:  []*mappable.Mappable{},
		Parameters: base.ParametersFromInterfaces(parameters.Prototype().GetList()),
	}
}
