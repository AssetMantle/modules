// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/splits/utilities"
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

	if _, err := utilities.AddSplits(auxiliaryKeeper.mapper.NewCollection(context), auxiliaryRequest.OwnerID, auxiliaryRequest.AssetID, auxiliaryRequest.Value); err != nil {
		return nil, err
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
