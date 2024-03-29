// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/record"
)

func AddSplits(splits helpers.Collection, ownerID ids.IdentityID, assetID ids.AssetID, value sdkTypes.Int) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroInt()) {
		return nil, errorConstants.InvalidRequest.Wrapf("value must be greater than zero")
	}

	splitID := baseIDs.NewSplitID(assetID, ownerID)

	Mappable := splits.Fetch(key.NewKey(splitID)).GetMappable(key.NewKey(splitID))
	if Mappable == nil {
		splits.Add(record.NewRecord(baseIDs.NewSplitID(assetID, ownerID), base.NewSplit(value)))
	} else {
		splits.Mutate(record.NewRecord(splitID, mappable.GetSplit(Mappable).Add(value)))
	}

	return splits, nil
}

func SubtractSplits(splits helpers.Collection, ownerID ids.IdentityID, assetID ids.AssetID, value sdkTypes.Int) (helpers.Collection, error) {
	if value.LTE(sdkTypes.ZeroInt()) {
		return nil, errorConstants.InvalidRequest.Wrapf("value must be greater than zero")
	}

	splitID := baseIDs.NewSplitID(assetID, ownerID)

	Mappable := splits.Fetch(key.NewKey(splitID)).GetMappable(key.NewKey(splitID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("split with ID %s not found", splitID.AsString())
	}
	split := mappable.GetSplit(Mappable)

	switch split = split.Subtract(value); {
	case split.GetValue().LT(sdkTypes.ZeroInt()):
		return nil, errorConstants.InvalidRequest.Wrapf("split value cannot be negative")
	case split.GetValue().Equal(sdkTypes.ZeroInt()):
		splits.Remove(record.NewRecord(splitID, split))
	default:
		splits.Mutate(record.NewRecord(splitID, split))
	}

	return splits, nil
}
