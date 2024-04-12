// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"context"
	"reflect"

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

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, AuxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	if err := AuxiliaryRequest.Validate(); err != nil {
		return nil, err
	}

	auxiliaryRequest, ok := AuxiliaryRequest.(auxiliaryRequest)
	if !ok {
		return nil, errorConstants.InvalidRequest.Wrapf("invalid request type: %s", reflect.TypeOf(AuxiliaryRequest).String())
	}

	if !auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.TransferEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
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

	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
