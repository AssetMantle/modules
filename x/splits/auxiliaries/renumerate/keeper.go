// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"context"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/internal/utilities"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	splits := auxiliaryKeeper.mapper.NewCollection(context)

	switch totalSplitsValue := utilities.GetOwnableTotalSplitsValue(splits, auxiliaryRequest.OwnableID); {
	case totalSplitsValue.LT(auxiliaryRequest.Value):
		if _, err := utilities.AddSplits(splits, auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID, auxiliaryRequest.Value.Sub(totalSplitsValue)); err != nil {
			return nil, err
		}
	case totalSplitsValue.GT(auxiliaryRequest.Value):
		if _, err := utilities.SubtractSplits(splits, auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID, totalSplitsValue.Sub(auxiliaryRequest.Value)); err != nil {
			return nil, err
		}
	case totalSplitsValue.IsZero():
		return nil, errorConstants.EntityNotFound.Wrapf("no splits found for ownable %s", auxiliaryRequest.OwnableID.AsString())
	default:
		return newAuxiliaryResponse(), nil
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}