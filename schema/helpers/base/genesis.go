// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/asaskevich/govalidator"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	parametersSchema "github.com/AssetMantle/modules/schema/parameters"
)

type genesis struct {
	keyPrototype      func() helpers.Key
	mappablePrototype func() helpers.Mappable

	defaultMappableList  []helpers.Mappable
	defaultParameterList []parametersSchema.Parameter

	MappableList  []helpers.Mappable           `json:"mappableList"`
	ParameterList []parametersSchema.Parameter `json:"parameterList"`
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	return genesis.Initialize(genesis.defaultMappableList, genesis.defaultParameterList)
}
func (genesis genesis) Validate() error {
	if len(genesis.ParameterList) != len(genesis.defaultParameterList) {
		return constants.InvalidParameter
	}

	for _, parameter := range genesis.ParameterList {
		var isPresent bool
		for _, defaultParameter := range genesis.defaultParameterList {
			isPresent = false
			if defaultParameter.GetID().Compare(parameter.GetID()) == 0 {
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

	_, err := govalidator.ValidateStruct(genesis)

	return err
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
	var mappableList []helpers.Mappable

	appendMappableList := func(mappable helpers.Mappable) bool {
		mappableList = append(mappableList, mappable)
		return false
	}
	mapper.Iterate(context, genesis.keyPrototype(), appendMappableList)

	for _, defaultParameter := range genesis.defaultParameterList {
		parameters = parameters.Fetch(context, defaultParameter.GetID())
	}

	return genesis.Initialize(mappableList, parameters.GetList())
}
func (genesis genesis) Encode(jsonCodec sdkCodec.JSONCodec) []byte {
	bytes, err := jsonCodec.MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return bytes
}
func (genesis genesis) Decode(jsonCodec sdkCodec.JSONCodec, byte []byte) helpers.Genesis {
	newGenesis := genesis
	if err := jsonCodec.UnmarshalJSON(byte, &newGenesis); err != nil {
		panic(err)
	}

	return NewGenesis(genesis.keyPrototype, genesis.mappablePrototype, genesis.defaultMappableList, genesis.defaultParameterList).Initialize(newGenesis.MappableList, newGenesis.ParameterList)
}
func (genesis genesis) Initialize(mappableList []helpers.Mappable, parameterList []parametersSchema.Parameter) helpers.Genesis {
	if len(mappableList) == 0 {
		genesis.MappableList = genesis.defaultMappableList
	} else {
		genesis.MappableList = mappableList
	}

	if len(parameterList) == 0 {
		genesis.ParameterList = genesis.defaultParameterList
	} else {
		for _, defaultParameter := range genesis.defaultParameterList {
			for i, parameter := range parameterList {
				if defaultParameter.GetID().Compare(parameter.GetID()) == 0 {
					parameterList[i] = defaultParameter.Mutate(parameter.GetData())
				}
			}
		}
		genesis.ParameterList = parameterList
	}

	if err := genesis.Validate(); err != nil {
		panic(err)
	}

	return genesis
}

func (genesis genesis) GetParameterList() []parametersSchema.Parameter {
	return genesis.ParameterList
}
func (genesis genesis) GetMappableList() []helpers.Mappable {
	return genesis.MappableList
}

func NewGenesis(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable, defaultMappableList []helpers.Mappable, defaultParameterList []parametersSchema.Parameter) helpers.Genesis {
	return genesis{
		keyPrototype:         keyPrototype,
		mappablePrototype:    mappablePrototype,
		defaultMappableList:  defaultMappableList,
		defaultParameterList: defaultParameterList,
		MappableList:         []helpers.Mappable{},
		ParameterList:        []parametersSchema.Parameter{},
	}
}
