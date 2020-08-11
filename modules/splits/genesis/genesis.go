/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type genesisState struct {
	SplitList []mappables.Split
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}

func (genesisState genesisState) Validate() error {
	for _, split := range genesisState.SplitList {
		var _, Error = govalidator.ValidateStruct(split)
		if Error != nil {
			return errors.Wrap(constants.IncorrectMessage, Error.Error())
		}
		if split.GetSplit().LT(sdkTypes.ZeroDec()) {
			return constants.InsufficientBalance
		}
	}
	return nil
}

func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {
	for _, split := range genesisState.SplitList {
		mapper.Create(ctx, split)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	assetsID := base.NewID("")

	appendableAssetList := func(mappable traits.Mappable) bool {
		genesisState.SplitList = append(genesisState.SplitList, mappable.(mappables.Split))
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

func newGenesisState(splitList []mappables.Split) helpers.GenesisState {
	return genesisState{
		SplitList: splitList,
	}
}

var State = newGenesisState([]mappables.Split{})
