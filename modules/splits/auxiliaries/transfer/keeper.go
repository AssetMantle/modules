/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transfer

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	if auxiliaryRequest.Value.LTE(sdkTypes.ZeroDec()) {
		return newAuxiliaryResponse(errors.NotAuthorized)
	}

	fromSplitID := key.NewSplitID(auxiliaryRequest.FromID, auxiliaryRequest.OwnableID)
	splits := auxiliaryKeeper.mapper.NewCollection(context)

	fromSplit := splits.Fetch(key.FromID(fromSplitID)).Get(key.FromID(fromSplitID))
	if fromSplit == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}

	switch fromSplit = fromSplit.(mappables.Split).Send(auxiliaryRequest.Value).(mappables.Split); {
	case fromSplit.(mappables.Split).GetValue().LT(sdkTypes.ZeroDec()):
		return newAuxiliaryResponse(errors.NotAuthorized)
	case fromSplit.(mappables.Split).GetValue().Equal(sdkTypes.ZeroDec()):
		splits.Remove(fromSplit)
	default:
		splits.Mutate(fromSplit)
	}

	toSplitID := key.NewSplitID(auxiliaryRequest.ToID, auxiliaryRequest.OwnableID)

	if toSplit, ok := splits.Fetch(key.FromID(toSplitID)).Get(key.FromID(toSplitID)).(mappables.Split); !ok {
		splits.Add(mappable.NewSplit(toSplitID, auxiliaryRequest.Value))
	} else {
		splits.Mutate(toSplit.Receive(auxiliaryRequest.Value).(mappables.Split))
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
