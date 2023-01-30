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
	"github.com/AssetMantle/modules/schema/properties/base"
)

type parameters struct {
	parameterList  []helpers.Parameter
	paramsSubspace paramsTypes.Subspace
}

var _ helpers.Parameters = (*parameters)(nil)

func (parameters parameters) Get() []helpers.Parameter {
	return parameters.parameterList
}
func (parameters parameters) Validate() error {

	if len(genesis.GetParameterList()) != len(genesis.Default().GetParameterList()) {
		return constants.InvalidParameter
	}

	for _, parameter := range genesis.GetParameterList() {
		var isPresent bool
		for _, defaultParameter := range genesis.Default().GetParameterList() {
			isPresent = false
			if defaultParameter.GetMetaProperty().Compare(parameter.GetMetaProperty()) == 0 {
				isPresent = true
				break
			}
		}

		if !isPresent {
			return constants.InvalidParameter
		}

		if err := parameter.Validate(); err != nil {
			return err
		}
	}

	for _, parameter := range parameters.parameterList {
		if err := parameter.Validate(); err != nil {
			return err
		}
	}

	return nil
}
func (parameters parameters) Fetch(context context.Context) helpers.Parameters {
	for i, parameter := range parameters.parameterList {
		var anyData data.AnyData
		parameters.paramsSubspace.Get(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), &anyData)
		parameters.parameterList[i] = parameter.Mutate(base.NewMetaProperty(parameter.GetMetaProperty().GetID().GetKey(), anyData))
	}

	return parameters
}
func (parameters parameters) Set(context context.Context) {
	for _, parameter := range parameters.parameterList {
		parameters.paramsSubspace.Set(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), parameter.GetMetaProperty().GetData())
	}
}
func (parameters parameters) ParamSetPairs() paramsTypes.ParamSetPairs {
	paramSetPairList := make([]paramsTypes.ParamSetPair, len(parameters.parameterList))

	for i, parameter := range parameters.parameterList {
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
func (parameters parameters) GetKeyTable() paramsTypes.KeyTable {
	return paramsTypes.NewKeyTable().RegisterParamSet(parameters)
}
func (parameters parameters) Initialize(subspace paramsTypes.Subspace) helpers.Parameters {
	parameters.paramsSubspace = subspace
	return parameters
}

func NewParameters(parameterList ...helpers.Parameter) helpers.Parameters {
	return parameters{
		parameterList: parameterList,
	}
}
