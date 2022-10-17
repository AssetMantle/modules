// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	if auxiliaryRequest.Value.LTE(sdkTypes.ZeroDec()) {
		return newAuxiliaryResponse(constants.NotAuthorized)
	}

	splits := auxiliaryKeeper.mapper.NewCollection(context)

	fromSplitID := baseIDs.NewSplitID(auxiliaryRequest.FromID, auxiliaryRequest.OwnableID)
	Mappable := splits.Fetch(key.NewKey(fromSplitID)).Get(key.NewKey(fromSplitID))
	if Mappable == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}
	fromSplit := Mappable.(types.Split)

	switch fromSplit = fromSplit.(types.Split).Send(auxiliaryRequest.Value).(types.Split); {
	case fromSplit.(types.Split).GetValue().LT(sdkTypes.ZeroDec()):
		return newAuxiliaryResponse(constants.NotAuthorized)
	case fromSplit.(types.Split).GetValue().Equal(sdkTypes.ZeroDec()):
		splits.Remove(mappable.NewMappable(fromSplit))
	default:
		splits.Mutate(mappable.NewMappable(fromSplit))
	}

	toSplitID := baseIDs.NewSplitID(auxiliaryRequest.ToID, auxiliaryRequest.OwnableID)

	if toSplit, ok := splits.Fetch(key.NewKey(toSplitID)).Get(key.NewKey(toSplitID)).(types.Split); !ok {
		splits.Add(mappable.NewMappable(base.NewSplit(auxiliaryRequest.ToID, auxiliaryRequest.OwnableID, auxiliaryRequest.Value)))
	} else {
		splits.Mutate(mappable.NewMappable(toSplit.Receive(auxiliaryRequest.Value).(types.Split)))
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
