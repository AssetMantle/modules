// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/parameters"
)

type validatableParameter struct {
	parameter parameters.Parameter
	validator func(parameters.Parameter) error
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
func (validatableParameter validatableParameter) Validate() error {
	if validatableParameter.validator == nil {
		return fmt.Errorf("validator is not set for parameter  " + validatableParameter.parameter.GetMetaProperty().GetID().AsString())
	}

	return validatableParameter.validator(validatableParameter.parameter)
}

func NewValidatableParameter(parameter parameters.Parameter, validator func(parameters.Parameter) error) helpers.ValidatableParameter {
	return &validatableParameter{
		parameter: parameter,
		validator: validator,
	}
}
