// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/mappables"

	"github.com/AssetMantle/modules/schema/helpers"
)

func AddSplits(splits helpers.Collection, ownerID ids.ID, ownableID ids.ID, value sdkTypes.Dec) (helpers.Collection, error) {
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

func SubtractSplits(splits helpers.Collection, ownerID ids.ID, ownableID ids.ID, value sdkTypes.Dec) (helpers.Collection, error) {
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
