/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"gopkg.in/validator.v2"
)

type genesisState struct {
	AssetList []mappables.InterNFT
}

//MapList := map[genesisState.AssetList]bool{}

var _ helpers.GenesisState = (*genesisState)(nil)


func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}


func (genesisState genesisState) Validate() error {
	for _, asset := range genesisState.AssetList {
		if errs := validator.Validate(asset); errs != nil {
			return errs
		}
	}
	return nil
}
func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, asset := range genesisState.AssetList {
		mapper.Create(ctx, asset)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	assetsID := base.NewID("")

	appendableAssetList := func(mappable traits.Mappable) bool {
		genesisState.AssetList = append(genesisState.AssetList, mappable.(mappables.InterNFT))
		return false
	}
	mapper.Iterate(context, assetsID, appendableAssetList)
	return genesisState
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

func newGenesisState(assetList []mappables.InterNFT) helpers.GenesisState {
	return genesisState{
		AssetList: assetList,
	}
}

var State = newGenesisState([]mappables.InterNFT{})
