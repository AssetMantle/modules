/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package utilities

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

func AddSplits(splits helpers.Collection, ownerID types.ID, ownableID types.ID, value sdkTypes.Dec) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroDec()) {
		return nil, errors.NotAuthorized
	}

	splitID := key.NewSplitID(ownerID, ownableID)

	split := splits.Fetch(key.FromID(splitID)).Get(key.FromID(splitID))
	if split == nil {
		splits.Add(mappable.NewSplit(splitID, value))
	} else {
		splits.Mutate(split.(mappables.Split).Receive(value).(mappables.Split))
	}

	return splits, nil
}

func SubtractSplits(splits helpers.Collection, ownerID types.ID, ownableID types.ID, value sdkTypes.Dec) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroDec()) {
		return nil, errors.NotAuthorized
	}

	splitsKey := key.FromID(key.NewSplitID(ownerID, ownableID))

	split := splits.Fetch(splitsKey).Get(splitsKey)
	if split == nil {
		return nil, errors.EntityNotFound
	}

	switch split = split.(mappables.Split).Send(value).(mappables.Split); {
	case split.(mappables.Split).GetValue().LT(sdkTypes.ZeroDec()):
		return nil, errors.NotAuthorized
	case split.(mappables.Split).GetValue().Equal(sdkTypes.ZeroDec()):
		splits.Remove(split)
	default:
		splits.Mutate(split)
	}

	return splits, nil
}
