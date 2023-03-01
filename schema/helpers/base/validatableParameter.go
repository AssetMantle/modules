package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
)

type validatableParameter struct {
	parameter helpers.Parameter
	validator func(i interface{}) error
}

var _ helpers.ValidatableParameter = (*validatableParameter)(nil)

func (validatableParameter validatableParameter) GetParameter() helpers.Parameter {
	return validatableParameter.parameter
}
func (validatableParameter validatableParameter) Mutate(data data.Data) helpers.ValidatableParameter {
	validatableParameter.parameter = validatableParameter.parameter.Mutate(data)
	return validatableParameter
}
func (validatableParameter validatableParameter) GetValidator() func(i interface{}) error {
	return validatableParameter.validator
}
func (validatableParameter validatableParameter) Validate() error {
	return validatableParameter.validator(validatableParameter.parameter)
}

func NewValidatableParameter(parameter helpers.Parameter, validator func(i interface{}) error) helpers.ValidatableParameter {
	return &validatableParameter{
		parameter: parameter,
		validator: validator,
	}
}
