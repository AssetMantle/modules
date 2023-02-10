// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package bond

import (
	"context"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper        helpers.Mapper
	bankKeeper    bankKeeper.Keeper
	stakingKeeper stakingKeeper.Keeper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.classificationID))

	Mappable := classifications.Get(key.NewKey(auxiliaryRequest.classificationID))
	if Mappable == nil {
		return newAuxiliaryResponse(errorConstants.EntityNotFound)
	}
	classification := mappable.GetClassification(Mappable)

	bondCoin, _ := sdkTypes.NewDecCoinFromDec(auxiliaryKeeper.stakingKeeper.BondDenom(sdkTypes.UnwrapSDKContext(context)), classification.GetBondAmount()).TruncateDecimal()
	if err := auxiliaryKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.accAddress, auxiliaryRequest.moduleName, sdkTypes.NewCoins(bondCoin)); err != nil {
		return newAuxiliaryResponse(err)
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterList, auxiliaries []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bankKeeper.Keeper:
			auxiliaryKeeper.bankKeeper = value
			return auxiliaryKeeper
		case stakingKeeper.Keeper:
			auxiliaryKeeper.stakingKeeper = value
			return auxiliaryKeeper
		}
	}

	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
