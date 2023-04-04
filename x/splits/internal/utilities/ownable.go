// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/x/splits/internal/key"
	"github.com/AssetMantle/modules/x/splits/internal/mappable"
)

func GetOwnableTotalSplitsValue(collection helpers.Collection, ownableID ids.ID) sdkTypes.Dec {
	value := sdkTypes.ZeroDec()
	accumulator := func(Mappable helpers.Mappable) bool {
		if mappable.GetSplit(Mappable).GetOwnableID().Compare(ownableID) == 0 {
			value = value.Add(mappable.GetSplit(Mappable).GetValue())
		}

		return false
	}
	collection.Iterate(key.NewKey(baseIDs.PrototypeSplitID()), accumulator)

	return value
}
