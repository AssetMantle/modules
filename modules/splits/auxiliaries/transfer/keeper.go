// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"context"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	if auxiliaryRequest.Value.LTE(sdkTypes.ZeroDec()) {
		return nil, errorConstants.NotAuthorized
	}

	splits := auxiliaryKeeper.mapper.NewCollection(context)

	fromSplitID := baseIDs.NewSplitID(auxiliaryRequest.FromID, auxiliaryRequest.OwnableID)
	Mappable := splits.Fetch(key.NewKey(fromSplitID)).Get(key.NewKey(fromSplitID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound
	}
	fromSplit := mappable.GetSplit(Mappable)

	switch fromSplit = fromSplit.Send(auxiliaryRequest.Value); {
	case fromSplit.GetValue().LT(sdkTypes.ZeroDec()):
		return nil, errorConstants.NotAuthorized
	case fromSplit.GetValue().Equal(sdkTypes.ZeroDec()):
		splits.Remove(mappable.NewMappable(fromSplit))
	default:
		splits.Mutate(mappable.NewMappable(fromSplit))
	}

	toSplitID := baseIDs.NewSplitID(auxiliaryRequest.ToID, auxiliaryRequest.OwnableID)

	if Mappable := splits.Fetch(key.NewKey(toSplitID)).Get(key.NewKey(toSplitID)); Mappable == nil {
		splits.Add(mappable.NewMappable(base.NewSplit(auxiliaryRequest.ToID, auxiliaryRequest.OwnableID, auxiliaryRequest.Value)))
	} else {
		splits.Mutate(mappable.NewMappable(mappable.GetSplit(Mappable).Receive(auxiliaryRequest.Value)))
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
