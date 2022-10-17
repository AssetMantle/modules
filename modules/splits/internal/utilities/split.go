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
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"

	"github.com/AssetMantle/modules/schema/helpers"
)

func AddSplits(splits helpers.Collection, ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroDec()) {
		return nil, constants.NotAuthorized
	}

	splitID := baseIDs.NewSplitID(ownerID, ownableID)

	split := splits.Fetch(key.NewKey(splitID)).Get(key.NewKey(splitID))
	if split == nil {
		splits.Add(mappable.NewMappable(base.NewSplit(ownerID, ownableID, value)))
	} else {
		splits.Mutate(mappable.NewMappable(split.(types.Split).Receive(value).(types.Split)))
	}

	return splits, nil
}

func SubtractSplits(splits helpers.Collection, ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroDec()) {
		return nil, constants.NotAuthorized
	}

	splitsKey := key.NewKey(baseIDs.NewSplitID(ownerID, ownableID))

	Mappable := splits.Fetch(splitsKey).Get(splitsKey)
	if Mappable == nil {
		return nil, constants.EntityNotFound
	}
	split := Mappable.(types.Split)

	switch split = split.Send(value).(types.Split); {
	case split.GetValue().LT(sdkTypes.ZeroDec()):
		return nil, constants.NotAuthorized
	case split.GetValue().Equal(sdkTypes.ZeroDec()):
		splits.Remove(mappable.NewMappable(split))
	default:
		splits.Mutate(mappable.NewMappable(split))
	}

	return splits, nil
}
