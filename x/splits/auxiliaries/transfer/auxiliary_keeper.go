// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"context"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/utilities"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
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

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
