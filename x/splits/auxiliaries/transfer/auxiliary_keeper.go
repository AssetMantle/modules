// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/splits/utilities"
	"github.com/AssetMantle/schema/data"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
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

	if !auxiliaryKeeper.parameterManager.Fetch(context).Get().GetParameter(propertyConstants.TransferEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("transfer is not enabled")
	}

	splits := auxiliaryKeeper.mapper.NewCollection(context)

	if _, err := utilities.SubtractSplits(splits, auxiliaryRequest.FromID, auxiliaryRequest.AssetID, auxiliaryRequest.Value); err != nil {
		return nil, err
	}

	if _, err := utilities.AddSplits(splits, auxiliaryRequest.ToID, auxiliaryRequest.AssetID, auxiliaryRequest.Value); err != nil {
		return nil, err
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper, auxiliaryKeeper.parameterManager = mapper, parameterManager

	helpers.PanicOnUninitializedKeeperFields(auxiliaryKeeper)
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
