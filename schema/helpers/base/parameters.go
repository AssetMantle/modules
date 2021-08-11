/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type parameters struct {
	parameterList  []types.Parameter
	paramsSubspace paramTypes.Subspace
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
		if Error := parameter.Validate(); Error != nil {
			return Error
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
func (parameters parameters) Get(id types.ID) types.Parameter {
	for _, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(id) == 0 {
			return parameter
		}
	}

	return nil
}
func (parameters parameters) GetList() []types.Parameter {
	return parameters.parameterList
}
func (parameters parameters) Fetch(context sdkTypes.Context, id types.ID) helpers.Parameters {
	var data types.Data

	parameters.paramsSubspace.Get(context, id.Bytes(), &data)

	for i, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(id) == 0 {
			parameters.parameterList[i] = parameter.Mutate(data)
		}
	}

	return parameters
}
func (parameters parameters) Mutate(context sdkTypes.Context, newParameter types.Parameter) helpers.Parameters {
	for i, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(newParameter.GetID()) == 0 {
			parameters.parameterList[i] = newParameter
			parameters.paramsSubspace.Set(context, parameter.GetID().Bytes(), parameter.GetData())

			break
		}
	}

	return parameters
}
func (parameters parameters) ParamSetPairs() paramTypes.ParamSetPairs {
	paramSetPairList := make([]paramTypes.ParamSetPair, len(parameters.parameterList))

	for i, parameter := range parameters.parameterList {
		paramSetPairList[i] = paramTypes.NewParamSetPair(parameter.GetID().Bytes(), parameter.GetData(), parameter.GetValidator())
	}

	return paramSetPairList
}
func (parameters parameters) GetKeyTable() paramTypes.KeyTable {
	return paramTypes.NewKeyTable().RegisterParamSet(parameters)
}
func (parameters parameters) Initialize(paramsSubspace paramTypes.Subspace) helpers.Parameters {
	parameters.paramsSubspace = paramsSubspace
	return parameters
}

func NewParameters(parameterList ...types.Parameter) helpers.Parameters {
	return parameters{
		parameterList: parameterList,
	}
}
