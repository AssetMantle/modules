// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types/base"
)

func AddSplits(splits helpers.Collection, ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroDec()) {
		return nil, errorConstants.InvalidRequest.Wrapf("value must be greater than zero")
	}

	splitID := baseIDs.NewSplitID(ownerID, ownableID)

	Mappable := splits.Fetch(key.NewKey(splitID)).Get(key.NewKey(splitID))
	if Mappable == nil {
		splits.Add(mappable.NewMappable(base.NewSplit(ownerID, ownableID, value)))
	} else {
		splits.Mutate(mappable.NewMappable(mappable.GetSplit(Mappable).Receive(value)))
	}

	return splits, nil
}

func SubtractSplits(splits helpers.Collection, ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroDec()) {
		return nil, errorConstants.InvalidRequest.Wrapf("value must be greater than zero")
	}

	splitID := baseIDs.NewSplitID(ownerID, ownableID)

	Mappable := splits.Fetch(key.NewKey(splitID)).Get(key.NewKey(splitID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("split with ID %s not found", splitID.AsString())
	}
	split := mappable.GetSplit(Mappable)

	switch split = split.Send(value); {
	case split.GetValue().LT(sdkTypes.ZeroDec()):
		return nil, errorConstants.InvalidRequest.Wrapf("split value cannot be negative")
	case split.GetValue().Equal(sdkTypes.ZeroDec()):
		splits.Remove(mappable.NewMappable(split))
	default:
		splits.Mutate(mappable.NewMappable(split))
	}

	return splits, nil
}
