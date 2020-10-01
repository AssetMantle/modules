/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transfer

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mapper"
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
	fromSplitID := mapper.NewSplitID(auxiliaryRequest.FromID, auxiliaryRequest.OwnableID)
	splits := mapper.NewSplits(auxiliaryKeeper.mapper, context).Fetch(fromSplitID)
	fromSplit := splits.Get(fromSplitID)
	if fromSplit == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}
	fromSplit = fromSplit.Send(auxiliaryRequest.Split).(mappables.Split)
	if fromSplit.GetSplit().LT(sdkTypes.ZeroDec()) {
		return newAuxiliaryResponse(errors.NotAuthorized)
	} else if fromSplit.GetSplit().Equal(sdkTypes.ZeroDec()) {
		splits.Remove(fromSplit)
	} else {
		splits.Mutate(fromSplit)
	}

	toSplitID := mapper.NewSplitID(auxiliaryRequest.ToID, auxiliaryRequest.OwnableID)
	toSplit := splits.Fetch(toSplitID).Get(toSplitID)
	if toSplit == nil {
		splits.Add(mapper.NewSplit(toSplitID, auxiliaryRequest.Split))
	} else {
		splits.Mutate(toSplit.Receive(auxiliaryRequest.Split).(mappables.Split))
	}
	return newAuxiliaryResponse(nil)
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
