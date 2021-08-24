/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type genesis struct {
	legacyAminoCodec *codec.LegacyAmino

	keyPrototype      func() helpers.Key
	mappablePrototype func() helpers.Mappable

	defaultMappableList  []helpers.Mappable
	defaultParameterList []types.Parameter

	MappableList  []helpers.Mappable `json:"mappableList"`
	ParameterList []types.Parameter  `json:"parameterList"`
}

func (genesis genesis) Encode(marshaler codec.JSONMarshaler) []byte {
	panic("implement me")
}

func (genesis genesis) Decode(marshaler codec.JSONMarshaler, bytes []byte) helpers.Genesis {
	panic("implement me")
}

func (genesis genesis) Size() int {
	panic("implement me")
}

func (genesis genesis) MarshalTo(data []byte) (int, error) {
	panic("implement me")
}

func (genesis genesis) Unmarshal(dAtA []byte) error {
	panic("implement me")
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	return genesis.Initialize(genesis.defaultMappableList, genesis.defaultParameterList)
}
func (genesis genesis) Validate() error {
	if len(genesis.ParameterList) != len(genesis.defaultParameterList) {
		return errors.InvalidParameter
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
func (genesis genesis) LegacyAminoEncode() []byte {
	bytes, Error := genesis.legacyAminoCodec.MarshalJSON(genesis)
	if Error != nil {
		panic(Error)
	}

	return bytes
}
func (genesis genesis) LegacyAminoDecode(byte []byte) helpers.Genesis {
	newGenesis := genesis
	if Error := genesis.legacyAminoCodec.UnmarshalJSON(byte, &newGenesis); Error != nil {
		panic(Error)
	}

	return NewGenesis(genesis.keyPrototype, genesis.mappablePrototype, genesis.defaultMappableList, genesis.defaultParameterList).Initialize(newGenesis.MappableList, newGenesis.ParameterList)
}
func (genesis genesis) Initialize(mappableList []helpers.Mappable, parameterList []types.Parameter) helpers.Genesis {
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

	if Error := genesis.Validate(); Error != nil {
		panic(Error)
	}

	return genesis
}

func (genesis genesis) GetParameterList() []types.Parameter {
	return genesis.ParameterList
}
func (genesis genesis) GetMappableList() []helpers.Mappable {
	return genesis.MappableList
}

func NewGenesis(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable, defaultMappableList []helpers.Mappable, defaultParameterList []types.Parameter) helpers.Genesis {
	LegacyAminoCodec := codec.NewLegacyAmino()
	keyPrototype().RegisterLegacyAminoCodec(LegacyAminoCodec)
	mappablePrototype().RegisterLegacyAminoCodec(LegacyAminoCodec)
	schema.RegisterLegacyAminoCodec(LegacyAminoCodec)
	LegacyAminoCodec.Seal()

	return genesis{
		legacyAminoCodec:     LegacyAminoCodec,
		keyPrototype:         keyPrototype,
		mappablePrototype:    mappablePrototype,
		defaultMappableList:  defaultMappableList,
		defaultParameterList: defaultParameterList,
		MappableList:         []helpers.Mappable{},
		ParameterList:        []types.Parameter{},
	}
}
