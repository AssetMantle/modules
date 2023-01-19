// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"

	"github.com/asaskevich/govalidator"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	parametersSchema "github.com/AssetMantle/modules/schema/parameters"
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

	// TODO ***** define validation for mappable list
	_, err := govalidator.ValidateStruct(genesis)

	return err
}
func (genesis genesis) Import(context context.Context, mapper helpers.Mapper, parameters helpers.Parameters) {
	for _, mappable := range genesis.GetMappableList() {
		mapper.Create(context, mappable)
	}

	for _, parameter := range genesis.GetParameterList() {
		parameters.Mutate(context, parameter)
	}
}
func (genesis genesis) Export(context context.Context, mapper helpers.Mapper, parameters helpers.Parameters) helpers.Genesis {
	var mappableList []helpers.Mappable

	appendMappableList := func(mappable helpers.Mappable) bool {
		mappableList = append(mappableList, mappable)
		return false
	}
	mapper.Iterate(context, genesis.keyPrototype(), appendMappableList)

	for _, defaultParameter := range genesis.Default().GetParameterList() {
		parameters = parameters.Fetch(context, defaultParameter.GetID())
	}

	return genesis.Initialize(mappableList, parameters.GetList())
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
func (genesis genesis) Initialize(mappableList []helpers.Mappable, parameterList []parametersSchema.Parameter) helpers.Genesis {
	genesis.genesisState = genesis.genesisState.Initialize(mappableList, parameterList)
	return genesis
}

func (genesis genesis) GetParameterList() []parametersSchema.Parameter {
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
