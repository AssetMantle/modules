// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	parameters2 "github.com/AssetMantle/modules/schema/parameters"
)

type parameters struct {
	parameterList  []parameters2.Parameter
	paramsSubspace params.Subspace
}

var _ helpers.Parameters = (*parameters)(nil)

func (parameters parameters) String() string {
	parameterList := make([]string, len(parameters.parameterList))
	for i, parameter := range parameters.parameterList {
		parameterList[i] = parameter.String()
	}

	return strings.Join(parameterList, "\n")
}
func (parameters parameters) Validate() error {
	for _, parameter := range parameters.parameterList {
		if err := parameter.Validate(); err != nil {
			return err
		}
	}

	return nil
}
func (parameters parameters) Equal(compareParameters helpers.Parameters) bool {
	for _, compareParameter := range compareParameters.GetList() {
		if !compareParameter.Equal(parameters.Get(compareParameter.GetID())) {
			return false
		}
	}

	return true
}
func (parameters parameters) Get(id ids.ID) parameters2.Parameter {
	for _, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(id) == 0 {
			return parameter
		}
	}

	return nil
}
func (parameters parameters) GetList() []parameters2.Parameter {
	return parameters.parameterList
}
func (parameters parameters) Fetch(context sdkTypes.Context, id ids.ID) helpers.Parameters {
	var Data data.Data

	parameters.paramsSubspace.Get(context, id.Bytes(), &Data)

	for i, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(id) == 0 {
			parameters.parameterList[i] = parameter.Mutate(Data)
		}
	}

	return parameters
}
func (parameters parameters) Mutate(context sdkTypes.Context, newParameter parameters2.Parameter) helpers.Parameters {
	for i, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(newParameter.GetID()) == 0 {
			parameters.parameterList[i] = newParameter
			parameters.paramsSubspace.Set(context, parameter.GetID().Bytes(), parameter.GetData())

			break
		}
	}

	return parameters
}
func (parameters parameters) ParamSetPairs() params.ParamSetPairs {
	paramSetPairList := make([]params.ParamSetPair, len(parameters.parameterList))

	for i, parameter := range parameters.parameterList {
		paramSetPairList[i] = params.NewParamSetPair(parameter.GetID().Bytes(), parameter.GetData(), parameter.GetValidator())
	}

	return paramSetPairList
}
func (parameters parameters) GetKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(parameters)
}
func (parameters parameters) Initialize(paramsSubspace params.Subspace) helpers.Parameters {
	parameters.paramsSubspace = paramsSubspace
	return parameters
}

func NewParameters(parameterList ...parameters2.Parameter) helpers.Parameters {
	return parameters{
		parameterList: parameterList,
	}
}
