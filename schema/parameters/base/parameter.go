package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
)

var _ helpers.Parameter = (*Parameter)(nil)

func (m *Parameter) ValidateBasic() error {
	return m.MetaProperty.ValidateBasic()
}
func (m *Parameter) GetMetaProperty() properties.MetaProperty {
	return m.MetaProperty
}
func (m *Parameter) Mutate(data data.Data) helpers.Parameter {
	if m.MetaProperty.GetData().GetTypeID().Compare(data.GetTypeID()) == 0 {
		m.MetaProperty = base.NewMetaProperty(m.MetaProperty.GetKey(), data).(*base.MetaProperty)
	}
	return m
}

func ParametersToInterface(parameters []*Parameter) []helpers.Parameter {
	returnParameters := make([]helpers.Parameter, len(parameters))

	for i, parameter := range parameters {
		returnParameters[i] = parameter
	}

	return returnParameters
}

func ParametersFromInterface(parameters []helpers.Parameter) []*Parameter {
	returnParameters := make([]*Parameter, len(parameters))

	for i, parameter := range parameters {
		returnParameters[i] = parameter.(*Parameter)
	}

	return returnParameters
}

func NewParameter(metaProperty properties.MetaProperty) *Parameter {
	return &Parameter{
		MetaProperty: metaProperty.(*base.MetaProperty),
	}
}
