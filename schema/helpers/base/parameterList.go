// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"golang.org/x/net/context"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

type parameterList struct {
	parameters     []helpers.Parameter
	paramsSubspace paramsTypes.Subspace
}

var _ helpers.ParameterList = (*parameterList)(nil)

func (parameterList parameterList) Get() []helpers.Parameter {
	return parameterList.parameters
}
func (parameterList parameterList) Fetch(context context.Context) helpers.ParameterList {
	for i, parameter := range parameterList.parameters {
		var anyData data.AnyData
		parameterList.paramsSubspace.Get(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), &anyData)
		parameterList.parameters[i] = parameter.Mutate(anyData)
	}

	return parameterList
}
func (parameterList parameterList) Set(context context.Context, parameters ...helpers.Parameter) {
	parameterList.parameters = parameters
	for _, parameter := range parameterList.parameters {
		parameterList.paramsSubspace.Set(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), parameter.GetMetaProperty().GetData())
	}
}
func (parameterList parameterList) ParamSetPairs() paramsTypes.ParamSetPairs {
	paramSetPairList := make([]paramsTypes.ParamSetPair, len(parameterList.parameters))

	for i, parameter := range parameterList.parameters {
		validator := func(i interface{}) error {
			switch value := i.(type) {
			case helpers.Parameter:
				return value.Validate()
			default:
				return constants.InvalidParameter
			}
		}
		paramSetPairList[i] = paramsTypes.NewParamSetPair(parameter.GetMetaProperty().GetID().Bytes(), parameter.GetMetaProperty().GetData(), validator)
	}

	return paramSetPairList
}
func (parameterList parameterList) GetKeyTable() paramsTypes.KeyTable {
	return paramsTypes.NewKeyTable().RegisterParamSet(parameterList)
}
func (parameterList parameterList) Initialize(subspace paramsTypes.Subspace) helpers.ParameterList {
	parameterList.paramsSubspace = subspace
	return parameterList
}

func NewParameterList(parameters ...helpers.Parameter) helpers.ParameterList {
	return parameterList{
		parameters: parameters,
	}
}
