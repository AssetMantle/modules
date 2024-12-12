// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package bond

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/constants"
	"github.com/AssetMantle/modules/x/classifications/key"
	"github.com/AssetMantle/modules/x/classifications/mappable"
)

type auxiliaryKeeper struct {
	mapper        helpers.Mapper
	bankKeeper    bankKeeper.Keeper
	stakingKeeper *stakingKeeper.Keeper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, AuxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest, ok := AuxiliaryRequest.(auxiliaryRequest)
	if !ok {
		return nil, errorConstants.InvalidRequest.Wrapf("invalid request type %T", AuxiliaryRequest)
	}

	if err := auxiliaryRequest.Validate(); err != nil {
		return nil, err
	}

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.classificationID))

	Mappable := classifications.GetMappable(key.NewKey(auxiliaryRequest.classificationID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("classification with ID %s not found", auxiliaryRequest.classificationID.AsString())
	}
	classification := mappable.GetClassification(Mappable)

	if auxiliaryRequest.bondAmount.LT(classification.GetBondAmount()) {
		return nil, errorConstants.InvalidRequest.Wrapf("bond amount %s is less than minimum bond amount %s", auxiliaryRequest.bondAmount.String(), classification.GetBondAmount().String())
	}

	if err := auxiliaryKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.accAddress, constants.ModuleName, sdkTypes.NewCoins(sdkTypes.NewCoin(auxiliaryKeeper.stakingKeeper.BondDenom(sdkTypes.UnwrapSDKContext(context)), auxiliaryRequest.bondAmount))); err != nil {
		return nil, err
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bankKeeper.Keeper:
			auxiliaryKeeper.bankKeeper = value
		case *stakingKeeper.Keeper:
			auxiliaryKeeper.stakingKeeper = value
		}
	}
	helpers.PanicOnUninitializedKeeperFields(auxiliaryKeeper)
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
