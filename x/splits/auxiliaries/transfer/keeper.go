// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"context"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	if auxiliaryRequest.Value.LTE(sdkTypes.ZeroInt()) {
		return nil, errorConstants.InvalidRequest.Wrapf("transfer value must be greater than zero")
	}

	splits := auxiliaryKeeper.mapper.NewCollection(context)

	fromSplitID := baseIDs.NewSplitID(auxiliaryRequest.FromID, auxiliaryRequest.OwnableID)
	Mappable := splits.Fetch(key.NewKey(fromSplitID)).Get(key.NewKey(fromSplitID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("split with ID %s not found", fromSplitID.AsString())
	}
	fromSplit := mappable.GetSplit(Mappable)

	switch fromSplit = fromSplit.Send(auxiliaryRequest.Value); {
	case fromSplit.GetValue().LT(sdkTypes.ZeroInt()):
		return nil, errorConstants.InsufficientBalance.Wrapf("insufficient balance")
	case fromSplit.GetValue().Equal(sdkTypes.ZeroInt()):
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
