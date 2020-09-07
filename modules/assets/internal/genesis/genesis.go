/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type genesisState struct {
	AssetList     []mappables.InterNFT `json:"assetList"`
	ParameterList []types.Parameter    `json:"parameterList"`
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}

func (genesisState genesisState) Validate() error {
	_, Error := govalidator.ValidateStruct(genesisState)
	return Error
}

func (genesisState genesisState) Initialize(context sdkTypes.Context, Mapper helpers.Mapper) {
	assets := mapper.NewAssets(context, Mapper)
	for _, asset := range genesisState.AssetList {
		assets = assets.Add(asset)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, Mapper helpers.Mapper) helpers.GenesisState {
	assets := mapper.NewAssets(context, Mapper).Fetch(base.NewID(""))
	//TODO add parameters
	return NewGenesisState(assets.GetList(), nil)
}

func (genesisState genesisState) Marshall() []byte {
	return packageCodec.MustMarshalJSON(genesisState)
}

func (genesisState genesisState) Unmarshall(byte []byte) helpers.GenesisState {
	if Error := packageCodec.UnmarshalJSON(byte, &genesisState); Error != nil {
		return nil
	}
	return genesisState
}

func NewGenesisState(assetList []mappables.InterNFT, parameterList ...types.Parameter) helpers.GenesisState {
	return genesisState{
		AssetList:     assetList,
		ParameterList: parameterList,
	}
}

var State = NewGenesisState(nil, nil)
