/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type genesis struct {
	codec                *codec.Codec
	MappableList         []traits.Mappable `json:"mappableList"`
	ParameterList        []types.Parameter `json:"parameterList"`
	DefaultMappableList  []traits.Mappable `json:"defaultMappableList"`
	DefaultParameterList []types.Parameter `json:"defaultParameterList"`
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	return genesis.Initialize(genesis.DefaultMappableList, genesis.DefaultParameterList)
}

func (genesis genesis) Validate() error {
	if len(genesis.ParameterList) != len(genesis.DefaultParameterList) {
		return errors.InvalidParameter
	}
	for _, parameter := range genesis.ParameterList {
		var isPresent bool
		for _, defaultParameter := range genesis.DefaultParameterList {
			isPresent = false
			if defaultParameter.GetID().Equal(parameter.GetID()) {
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
	_, Error := govalidator.ValidateStruct(genesis)
	return Error
}

func (genesis genesis) Import(context sdkTypes.Context, mapper helpers.Mapper, parameters helpers.Parameters) {
	for _, mappable := range genesis.MappableList {
		mapper.Create(context, mappable)
	}
	for _, parameter := range genesis.ParameterList {
		parameters.Mutate(context, parameter)
	}
}

func (genesis genesis) Export(context sdkTypes.Context, mapper helpers.Mapper, parameters helpers.Parameters) helpers.Genesis {
	var mappableList []traits.Mappable
	appendMappableList := func(mappable traits.Mappable) bool {
		mappableList = append(mappableList, mappable)
		return false
	}
	mapper.Iterate(context, base.NewID(""), appendMappableList)
	var parameterList []types.Parameter
	for _, defaultParameter := range genesis.DefaultParameterList {
		parameterList = append(parameterList, parameters.Fetch(context, defaultParameter.GetID()))
	}
	return genesis.Initialize(mappableList, parameterList)
}

func (genesis genesis) Marshall() []byte {
	genesis.codec.
		bytes, Error := json.Marshal(genesis)
	if Error != nil {
		panic(Error)
	}
	return bytes
}

func (genesis genesis) Unmarshall(byte []byte) helpers.Genesis {
	if Error := json.Unmarshal(byte, &genesis); Error != nil {
		panic(Error)
	}
	return genesis
}

func (genesis genesis) Initialize(mappableList []traits.Mappable, parameterList []types.Parameter) helpers.Genesis {
	genesis.MappableList = mappableList
	genesis.ParameterList = parameterList
	return genesis
}

func GenesisPrototype(codec *codec.Codec, defaultMappableList []traits.Mappable, defaultParameterList []types.Parameter) helpers.Genesis {
	return genesis{
		codec:                codec,
		MappableList:         []traits.Mappable{},
		ParameterList:        []types.Parameter{},
		DefaultMappableList:  defaultMappableList,
		DefaultParameterList: defaultParameterList,
	}
}
