// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/splits/utilities"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
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

	if auxiliaryRequest.Supply.LTE(sdkTypes.ZeroInt()) {
		return nil, errorConstants.IncorrectFormat.Wrapf("value is less than or equal to 0 for asset: %s", auxiliaryRequest.AssetID.AsString())
	}

	splits := auxiliaryKeeper.mapper.NewCollection(context)

	switch totalSplitsValue := utilities.GetTotalSupply(splits, auxiliaryRequest.AssetID); {
	case totalSplitsValue.LT(auxiliaryRequest.Supply):
		if _, err := utilities.AddSplits(splits, auxiliaryRequest.OwnerID, auxiliaryRequest.AssetID, auxiliaryRequest.Supply.Sub(totalSplitsValue)); err != nil {
			return nil, err
		}
	case totalSplitsValue.GT(auxiliaryRequest.Supply):
		if _, err := utilities.SubtractSplits(splits, auxiliaryRequest.OwnerID, auxiliaryRequest.AssetID, totalSplitsValue.Sub(auxiliaryRequest.Supply)); err != nil {
			return nil, err
		}
	case totalSplitsValue.IsZero():
		return nil, errorConstants.EntityNotFound.Wrapf("no splits found for assetID %s", auxiliaryRequest.AssetID.AsString())
	default:
		return newAuxiliaryResponse(), nil
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper

	helpers.PanicOnUninitializedKeeperFields(auxiliaryKeeper)
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
