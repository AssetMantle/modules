/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
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

func (parameters parameters) GetList() []types.Parameter {
	return parameters.parameterList
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

func (parameters parameters) ParamSetPairs() params.ParamSetPairs {
	var paramSetPairList []params.ParamSetPair
	for _, parameter := range parameters.parameterList {
		paramSetPairList = append(paramSetPairList, params.NewParamSetPair([]byte(parameter.GetKey()), parameter.GetData(), parameter.GetValidator()))
	}
	return paramSetPairList
}

func (parameters parameters) GetKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(parameters)
}

func (parameters *parameters) Initialize(paramsSubspace params.Subspace) {
	parameters.paramsSubspace = paramsSubspace
}

func NewParameters(parameterList ...types.Parameter) helpers.Parameters {
	return &parameters{
		parameterList: parameterList,
	}
}
