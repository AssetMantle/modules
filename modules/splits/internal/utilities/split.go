// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"

	"github.com/AssetMantle/modules/schema/helpers"
)

func AddSplits(splits helpers.Collection, ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroDec()) {
		return nil, constants.NotAuthorized
	}

	splitID := baseIDs.NewSplitID(ownerID, ownableID)

	split := splits.Fetch(key.NewKey(splitID)).Get(key.NewKey(splitID))
	if split == nil {
		splits.Add(mappable.NewSplit(ownerID, ownableID, value))
	} else {
		splits.Mutate(split.(mappables.Split).Receive(value).(mappables.Split))
	}

	return splits, nil
}

func SubtractSplits(splits helpers.Collection, ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroDec()) {
		return nil, constants.NotAuthorized
	}

	splitsKey := key.NewKey(baseIDs.NewSplitID(ownerID, ownableID))

	split := splits.Fetch(splitsKey).Get(splitsKey)
	if split == nil {
		return nil, constants.EntityNotFound
	}

	switch split = split.(mappables.Split).Send(value).(mappables.Split); {
	case split.(mappables.Split).GetValue().LT(sdkTypes.ZeroDec()):
		return nil, constants.NotAuthorized
	case split.(mappables.Split).GetValue().Equal(sdkTypes.ZeroDec()):
		splits.Remove(split)
	default:
		splits.Mutate(split)
	}

	return splits, nil
}
