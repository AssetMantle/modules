/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	splitID := key.NewSplitID(auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID)
	splits := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(splitID))

	split := splits.Get(key.FromID(splitID))
	if split == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}

	newSplit := split.(mappables.Split).Send(auxiliaryRequest.Value)
	switch split = mappable.NewSplit(base.NewID(splitID.String()), newSplit.GetValue()); {
	case split.(mappables.Split).GetValue().LT(sdkTypes.ZeroDec()):
		return newAuxiliaryResponse(errors.InsufficientBalance)
	case split.(mappables.Split).GetValue().Equal(sdkTypes.ZeroDec()):
		splits.Remove(split)
	default:
		splits.Mutate(split)
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
