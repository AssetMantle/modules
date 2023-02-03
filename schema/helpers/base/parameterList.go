// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"golang.org/x/net/context"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
)

type parameterList struct {
	validatableParameters []helpers.ValidatableParameter
	paramsSubspace        paramsTypes.Subspace
}

var _ helpers.ParameterList = (*parameterList)(nil)

func (parameterList parameterList) Get() []helpers.Parameter {
	parameters := make([]helpers.Parameter, len(parameterList.validatableParameters))
	for i, validatableParameter := range parameterList.validatableParameters {
		parameters[i] = validatableParameter.GetParameter()
	}
	return parameters
}
func (parameterList parameterList) Fetch(context context.Context) helpers.ParameterList {
	for i, validatableParameter := range parameterList.validatableParameters {
		var anyData data.AnyData
		parameterList.paramsSubspace.Get(sdkTypes.UnwrapSDKContext(context), validatableParameter.GetParameter().GetMetaProperty().GetID().Bytes(), &anyData)
		parameterList.validatableParameters[i] = validatableParameter.Mutate(anyData)
	}

	return parameterList
}
func (parameterList parameterList) Set(context context.Context, parameters ...helpers.Parameter) {
	for _, parameter := range parameters {
		parameterList.paramsSubspace.Set(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), parameter.GetMetaProperty().GetData())
	}
}
func (parameterList parameterList) ParamSetPairs() paramsTypes.ParamSetPairs {
	paramSetPairList := make([]paramsTypes.ParamSetPair, len(parameterList.validatableParameters))

	for i, validatableParameter := range parameterList.validatableParameters {
		paramSetPairList[i] = paramsTypes.NewParamSetPair(validatableParameter.GetParameter().GetMetaProperty().GetID().Bytes(), validatableParameter.GetParameter().GetMetaProperty().GetData(), validatableParameter.GetValidator())
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

func NewParameterList(validatableParameters ...helpers.ValidatableParameter) helpers.ParameterList {
	return parameterList{
		validatableParameters: validatableParameters,
	}
}
