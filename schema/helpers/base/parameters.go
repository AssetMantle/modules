// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"golang.org/x/net/context"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

type parameters struct {
	parameterList  []helpers.Parameter
	paramsSubspace paramsTypes.Subspace
}

var _ helpers.Parameters = (*parameters)(nil)

func (parameters parameters) String() string {
	parameterList := make([]string, len(parameters.parameterList))
	for i, parameter := range parameters.parameterList {
		parameterList[i] = parameter.AsString()
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
func (parameters parameters) Get(id ids.ID) helpers.Parameter {
	for _, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(id) == 0 {
			return parameter
		}
	}

	return nil
}
func (parameters parameters) GetList() []helpers.Parameter {
	return parameters.parameterList
}
func (parameters parameters) Fetch(context context.Context, id ids.ID) helpers.Parameters {
	var Data data.AnyData

	parameters.paramsSubspace.Get(sdkTypes.UnwrapSDKContext(context), id.Bytes(), &Data)

	for i, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(id) == 0 {
			parameters.parameterList[i] = parameter.Mutate(Data)
		}
	}

	return parameters
}
func (parameters parameters) Mutate(context context.Context, newParameter helpers.Parameter) helpers.Parameters {
	for i, parameter := range parameters.parameterList {
		if parameter.GetID().Compare(newParameter.GetID()) == 0 {
			parameters.parameterList[i] = newParameter
			parameters.paramsSubspace.Set(sdkTypes.UnwrapSDKContext(context), parameter.GetID().Bytes(), parameter.GetData())

			break
		}
	}

	return parameters
}
func (parameters parameters) ParamSetPairs() paramsTypes.ParamSetPairs {
	paramSetPairList := make([]paramsTypes.ParamSetPair, len(parameters.parameterList))

	for i, parameter := range parameters.parameterList {
		paramSetPairList[i] = paramsTypes.NewParamSetPair(parameter.GetID().Bytes(), parameter.GetData(), parameter.GetValidator())
	}

	return paramSetPairList
}
func (parameters parameters) GetKeyTable() paramsTypes.KeyTable {
	return paramsTypes.NewKeyTable().RegisterParamSet(parameters)
}
func (parameters parameters) Initialize(paramsSubspace paramsTypes.Subspace) helpers.Parameters {
	parameters.paramsSubspace = paramsSubspace
	return parameters
}

func NewParameters(parameterList ...helpers.Parameter) helpers.Parameters {
	return parameters{
		parameterList: parameterList,
	}
}
