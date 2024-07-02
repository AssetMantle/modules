// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package purge

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseIDs "github.com/AssetMantle/schema/ids/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/record"
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

	splits := auxiliaryKeeper.mapper.NewCollection(context)

	circulatingSupply := utilities.GetTotalSupply(splits, auxiliaryRequest.AssetID)
	if !circulatingSupply.Equal(auxiliaryRequest.Supply) {
		return nil, errorConstants.InvalidRequest.Wrapf("circulating supply %d doesn't match asset's supply %d", circulatingSupply, auxiliaryRequest.Supply)
	}

	splitID := baseIDs.NewSplitID(auxiliaryRequest.AssetID, auxiliaryRequest.OwnerID)
	Mappable := splits.Fetch(key.NewKey(splitID)).GetMappable(key.NewKey(splitID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("split with ID %s not found", splitID.AsString())
	}
	split := mappable.GetSplit(Mappable)

	if !split.GetValue().Equal(auxiliaryRequest.Supply) {
		return nil, errorConstants.InvalidRequest.Wrapf("owned value %d doesn't match asset's circulating supply %d", split.GetValue(), auxiliaryRequest.Supply)
	}

	splits.Remove(record.NewRecord(splitID, split))

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
