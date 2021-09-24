/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package renumerate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	splits := auxiliaryKeeper.mapper.NewCollection(context)

	switch totalSplitsValue := utilities.GetOwnableTotalSplitsValue(splits, auxiliaryRequest.OwnableID); {
	case totalSplitsValue.LT(auxiliaryRequest.Value):
		if _, Error := utilities.AddSplits(splits, auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID, auxiliaryRequest.Value.Sub(totalSplitsValue)); Error != nil {
			return newAuxiliaryResponse(Error)
		}
	case totalSplitsValue.GT(auxiliaryRequest.Value):
		if _, Error := utilities.SubtractSplits(splits, auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID, totalSplitsValue.Sub(auxiliaryRequest.Value)); Error != nil {
			return newAuxiliaryResponse(Error)
		}
	case totalSplitsValue.IsZero():
		return newAuxiliaryResponse(errors.EntityNotFound)
	default:
		return newAuxiliaryResponse(nil)
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
