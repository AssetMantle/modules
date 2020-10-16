/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type genesis struct {
	MappableList         []helpers.Mappable `json:"mappableList"`
	ParameterList        []types.Parameter  `json:"parameterList"`
	defaultMappableList  []helpers.Mappable
	defaultParameterList []types.Parameter
}

var _ helpers.Genesis = (*genesis)(nil)

func (Genesis genesis) Default() helpers.Genesis {
	return Genesis.Initialize(Genesis.defaultMappableList, Genesis.defaultParameterList)
}

func (Genesis genesis) Validate() error {
	if len(Genesis.ParameterList) != len(Genesis.defaultParameterList) {
		return errors.InvalidParameter
	}
	for _, parameter := range Genesis.ParameterList {
		var isPresent bool
		for _, defaultParameter := range Genesis.defaultParameterList {
			isPresent = false
			if defaultParameter.GetID().Equals(parameter.GetID()) {
				isPresent = true
				break
			}
		}
		if isPresent != true {
			return errors.InvalidParameter
		}
		if Error := parameter.Validate(); Error != nil {
			return Error
		}
	}
	_, Error := govalidator.ValidateStruct(Genesis)
	return Error
}

func (Genesis genesis) Import(context sdkTypes.Context, mapper helpers.Mapper, parameters helpers.Parameters) {
	for _, mappable := range Genesis.MappableList {
		mapper.Create(context, mappable)
	}
	for _, parameter := range Genesis.ParameterList {
		parameters.Mutate(context, parameter)
	}
}

func (Genesis genesis) Export(context sdkTypes.Context, mapper helpers.Mapper, parameters helpers.Parameters) helpers.Genesis {
	var mappableList []helpers.Mappable
	appendMappableList := func(mappable helpers.Mappable) bool {
		mappableList = append(mappableList, mappable)
		return false
	}
	mapper.Iterate(context, nil, appendMappableList)
	var parameterList []types.Parameter
	for _, defaultParameter := range Genesis.defaultParameterList {
		parameterList = append(parameterList, parameters.Fetch(context, defaultParameter.GetID()))
	}
	return Genesis.Initialize(mappableList, parameterList)
}

func (Genesis genesis) Encode() []byte {
	bytes, Error := Genesis.codec.MarshalJSON(Genesis)
	if Error != nil {
		panic(Error)
	}
	return bytes
}

func (Genesis genesis) Decode(byte []byte) helpers.Genesis {
	var genesis genesis
	if Error := Genesis.codec.UnmarshalJSON(byte, &genesis); Error != nil {
		panic(Error)
	}
	return NewGenesis(Genesis.defaultMappableList, Genesis.defaultParameterList).Initialize(genesis.MappableList, genesis.ParameterList)
}

func (Genesis genesis) Initialize(mappableList []helpers.Mappable, parameterList []types.Parameter) helpers.Genesis {
	for _, defaultParameter := range Genesis.defaultParameterList {
		for i, parameter := range parameterList {
			if defaultParameter.GetID().Equals(parameter.GetID()) {
				parameterList[i] = defaultParameter.Mutate(parameter.GetData())
			}
		}
	}
	Genesis.MappableList = mappableList
	Genesis.ParameterList = parameterList
	if Error := Genesis.Validate(); Error != nil {
		panic(Error)
	}
	return Genesis
}

func NewGenesis(defaultMappableList []helpers.Mappable, defaultParameterList []types.Parameter) helpers.Genesis {
	return genesis{
		MappableList:         []helpers.Mappable{},
		ParameterList:        []types.Parameter{},
		defaultMappableList:  defaultMappableList,
		defaultParameterList: defaultParameterList,
	}
}
