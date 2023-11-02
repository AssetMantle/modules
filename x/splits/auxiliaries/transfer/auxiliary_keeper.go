// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"context"

	"github.com/AssetMantle/schema/go/data"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	propertyConstants "github.com/AssetMantle/schema/go/properties/constants"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/utilities"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	if !auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.TransferEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("transfer is not enabled")
	}

	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	splits := auxiliaryKeeper.mapper.NewCollection(context)

	if _, err := utilities.SubtractSplits(splits, auxiliaryRequest.FromID, auxiliaryRequest.AssetID, auxiliaryRequest.Value); err != nil {
		return newAuxiliaryResponse(), err
	}

	if _, err := utilities.AddSplits(splits, auxiliaryRequest.ToID, auxiliaryRequest.AssetID, auxiliaryRequest.Value); err != nil {
		return newAuxiliaryResponse(), err
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper, auxiliaryKeeper.parameterManager = mapper, parameterManager

	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
