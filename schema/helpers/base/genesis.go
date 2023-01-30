// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"

	"github.com/asaskevich/govalidator"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

type genesis struct {
	keyPrototype func() helpers.Key

	genesisState helpers.GenesisState
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	genesis.genesisState = genesis.genesisState.Default()
	return genesis
}
func (genesis genesis) Validate() error {
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

	// TODO ***** define validation for mappable list
	_, err := govalidator.ValidateStruct(genesis)

	return err
}
func (genesis genesis) Import(context context.Context, mapper helpers.Mapper) {
	for _, mappable := range genesis.GetMappableList() {
		mapper.Create(context, mappable)
	}

	NewParameters(genesis.GetParameterList()...).Set(context)
}
func (genesis genesis) Export(context context.Context, mapper helpers.Mapper) helpers.Genesis {
	var mappableList []helpers.Mappable

	appendMappableList := func(mappable helpers.Mappable) bool {
		mappableList = append(mappableList, mappable)
		return false
	}
	mapper.Iterate(context, genesis.keyPrototype(), appendMappableList)

	return genesis.Initialize(mappableList, NewParameters(genesis.Default().GetParameterList()...).Fetch(context).Get())
}
func (genesis genesis) Encode(jsonCodec sdkCodec.JSONCodec) []byte {
	bytes, err := jsonCodec.MarshalJSON(genesis.genesisState)
	if err != nil {
		panic(err)
	}

	return bytes
}
func (genesis genesis) Decode(jsonCodec sdkCodec.JSONCodec, byte []byte) helpers.Genesis {
	newGenesisState := genesis.genesisState
	if err := jsonCodec.UnmarshalJSON(byte, newGenesisState); err != nil {
		panic(err)
	}

	return NewGenesis(genesis.keyPrototype, newGenesisState)
}
func (genesis genesis) Initialize(mappableList []helpers.Mappable, parameterList []helpers.Parameter) helpers.Genesis {
	genesis.genesisState = genesis.genesisState.Initialize(mappableList, parameterList)
	return genesis
}

func (genesis genesis) GetParameterList() []helpers.Parameter {
	return genesis.genesisState.GetParameters()
}
func (genesis genesis) GetMappableList() []helpers.Mappable {
	return genesis.genesisState.GetMappables()
}

func NewGenesis(keyPrototype func() helpers.Key, genesisState helpers.GenesisState) helpers.Genesis {
	return genesis{
		keyPrototype: keyPrototype,
		genesisState: genesisState,
	}
}
