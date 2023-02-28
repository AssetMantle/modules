package base

import (
	"github.com/AssetMantle/modules/schema/helpers"
)

var _ helpers.ParameterList = (*ParameterList)(nil)

func (parameterList *ParameterList) Get() []helpers.Parameter {
	parameters := make([]helpers.Parameter, len(parameterList.Parameters))
	for i, parameter := range parameterList.Parameters {
		parameters[i] = parameter
	}
	return parameters
}

func NewParameterList(parameters ...helpers.Parameter) helpers.ParameterList {
	Parameters := make([]*Parameter, len(parameters))
	for i, parameter := range parameters {
		Parameters[i] = parameter.(*Parameter)
	}

	return &ParameterList{
		Parameters: Parameters,
	}
}
