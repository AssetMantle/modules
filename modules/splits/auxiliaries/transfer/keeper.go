/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
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

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	if auxiliaryRequest.Split.LTE(sdkTypes.ZeroDec()) {
		return newAuxiliaryResponse(errors.NotAuthorized)
	}
	fromSplitID := key.NewSplitID(auxiliaryRequest.FromID, auxiliaryRequest.OwnableID)
	splits := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.New(fromSplitID))
	fromSplit := splits.Get(key.New(fromSplitID))
	if fromSplit == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}
	fromSplit = fromSplit.(mappables.Split).Send(auxiliaryRequest.Split).(mappables.Split)
	if fromSplit.(mappables.Split).GetSplit().LT(sdkTypes.ZeroDec()) {
		return newAuxiliaryResponse(errors.NotAuthorized)
	} else if fromSplit.(mappables.Split).GetSplit().Equal(sdkTypes.ZeroDec()) {
		splits.Remove(fromSplit)
	} else {
		splits.Mutate(fromSplit)
	}

	toSplitID := key.NewSplitID(auxiliaryRequest.ToID, auxiliaryRequest.OwnableID)
	toSplit := splits.Fetch(key.New(toSplitID)).Get(key.New(toSplitID)).(mappables.Split)
	if toSplit == nil {
		splits.Add(mappable.NewSplit(toSplitID, auxiliaryRequest.Split))
	} else {
		splits.Mutate(toSplit.Receive(auxiliaryRequest.Split).(mappables.Split))
	}
	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
