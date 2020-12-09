/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"encoding/json"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type parameters struct {
	parameterList  []types.Parameter
	paramsSubspace params.Subspace
}

var _ helpers.Parameters = (*parameters)(nil)

func (parameters parameters) String() string {
	var parameterList []string
	for _, parameter := range parameters.parameterList {
		parameterList = append(parameterList, parameter.String())
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

func (parameters parameters) Equal(Parameters helpers.Parameters) bool {
	Bytes, Error := json.Marshal(parameters)
	if Error != nil {
		panic(Error)
	}
	CompareBytes, Error := json.Marshal(Parameters)
	if Error != nil {
		panic(Error)
	}
	return bytes.Compare(Bytes, CompareBytes) == 0
}

func (parameters parameters) Fetch(context sdkTypes.Context, id types.ID) types.Parameter {
	var data types.Data
	parameters.paramsSubspace.Get(context, id.Bytes(), &data)
	var validator func(interface{}) error
	for _, parameter := range parameters.GetList() {
		if parameter.GetID().Equals(id) {
			validator = parameter.GetValidator()
		}
	}
	return base.NewParameter(id, data, validator)
}

func (parameters parameters) Get(id types.ID) types.Parameter {
	for _, parameter := range parameters.parameterList {
		if parameter.GetID().Equals(id) {
			return parameter
		}
	}
	return nil
}

func (parameters parameters) GetList() []types.Parameter {
	var parameterList []types.Parameter
	for _, parameter := range parameters.parameterList {
		parameterList = append(parameterList, parameter)
	}
	return parameterList
}

func (parameters parameters) Mutate(context sdkTypes.Context, Parameter types.Parameter) helpers.Parameters {
	for i, parameter := range parameters.parameterList {
		if parameter.GetID().Equals(Parameter.GetID()) {
			parameters.parameterList[i] = Parameter
			parameters.paramsSubspace.Set(context, parameter.GetID().Bytes(), parameter.GetData())
			break
		}
	}
	return parameters
}

func (parameters parameters) ParamSetPairs() params.ParamSetPairs {
	var paramSetPairList []params.ParamSetPair
	for _, parameter := range parameters.parameterList {
		paramSetPairList = append(paramSetPairList, params.NewParamSetPair(parameter.GetID().Bytes(), parameter.GetData(), parameter.GetValidator()))
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

func NewParameters(parameterList ...types.Parameter) helpers.Parameters {
	return parameters{
		parameterList: parameterList,
	}
}
