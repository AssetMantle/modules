package base

import (
	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/parameters"

	"github.com/AssetMantle/modules/helpers"
)

type validatableParameter struct {
	parameter parameters.Parameter
	validator func(i interface{}) error
}

var _ helpers.ValidatableParameter = (*validatableParameter)(nil)

func (validatableParameter validatableParameter) GetParameter() parameters.Parameter {
	return validatableParameter.parameter
}
func (validatableParameter validatableParameter) Mutate(data data.Data) helpers.ValidatableParameter {
	if data != nil {
		validatableParameter.parameter = validatableParameter.parameter.Mutate(data)
	}
	return validatableParameter
}
func (validatableParameter validatableParameter) GetValidator() func(i interface{}) error {
	return validatableParameter.validator
}
func (validatableParameter validatableParameter) Validate() error {
	if validatableParameter.validator == nil {
		return nil
	}
	return validatableParameter.validator(validatableParameter.parameter)
}

func NewValidatableParameter(parameter parameters.Parameter, validator func(i interface{}) error) helpers.ValidatableParameter {
	return &validatableParameter{
		parameter: parameter,
		validator: validator,
	}
}
