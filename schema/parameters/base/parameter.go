// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	parametersSchema "github.com/AssetMantle/modules/schema/parameters"
)

var _ parametersSchema.Parameter = (*Parameter)(nil)

func (parameter *Parameter) GetValidator() func(interface{}) error {
	// TODO correct
	return func(_ interface{}) error {
		return nil
	}
}

func (parameter *Parameter) Equal(compareParameter parametersSchema.Parameter) bool {
	if compareParameter != nil && parameter.ID.Compare(compareParameter.GetID()) == 0 && parameter.Data.GetType().Compare(compareParameter.GetData().GetType()) == 0 && parameter.Data.Compare(compareParameter.GetData()) == 0 {
		return true
	}

	return false
}
func (parameter *Parameter) Validate() error {
	// TODO ****** validate parameter
	return nil
}
func (parameter *Parameter) GetID() ids.ID {
	return parameter.ID
}
func (parameter *Parameter) GetData() data.AnyData {
	return parameter.Data
}
func (parameter *Parameter) Mutate(data data.Data) parametersSchema.Parameter {
	// TODO ****** data type check
	parameter.Data = data.ToAnyData().(*baseData.AnyData)
	return parameter
}

func NewParameter(id ids.StringID, data data.Data, validator func(interface{}) error) parametersSchema.Parameter {
	return &Parameter{
		ID:   id.(*baseIds.StringID),
		Data: data.ToAnyData().(*baseData.AnyData),
	}
}

func ParametersFromInterfaces(parameters []parametersSchema.Parameter) []*Parameter {
	Parameters := make([]*Parameter, len(parameters))
	for index, parameter := range parameters {
		Parameters[index] = parameter.(*Parameter)
	}
	return Parameters
}

func ParametersToInterfaces(parameters []*Parameter) []parametersSchema.Parameter {
	Parameters := make([]parametersSchema.Parameter, len(parameters))
	for index, parameter := range parameters {
		Parameters[index] = parameter
	}
	return Parameters
}
