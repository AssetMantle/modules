// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"context"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/utilities"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	splits := auxiliaryKeeper.mapper.NewCollection(context)

	circulatingSupply := utilities.GetTotalSupply(splits, auxiliaryRequest.OwnableID)
	if !circulatingSupply.Equal(auxiliaryRequest.Value) {
		return nil, errorConstants.InvalidRequest.Wrapf("circulating supply %d doesn't match asset's supply %d", circulatingSupply, auxiliaryRequest.Value)
	}

	splitID := baseIDs.NewSplitID(auxiliaryRequest.OwnableID, auxiliaryRequest.OwnerID)
	Mappable := splits.Fetch(key.NewKey(splitID)).GetMappable(key.NewKey(splitID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("split with ID %s not found", splitID.AsString())
	}
	split := mappable.GetSplit(Mappable)

	if !split.GetValue().Equal(auxiliaryRequest.Value) {
		return nil, errorConstants.InvalidRequest.Wrapf("owned value %d doesn't match asset's circulating supply %d", split.GetValue(), auxiliaryRequest.Value)
	}

	splits.Remove(mappable.NewMappable(split))

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
