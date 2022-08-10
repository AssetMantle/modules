// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
)

func GetOwnableTotalSplitsValue(collection helpers.Collection, ownableID ids.ID) sdkTypes.Dec {
	value := sdkTypes.ZeroDec()
	accumulator := func(mappable helpers.Mappable) bool {
		if mappable.(mappables.Split).GetOwnableID().Compare(ownableID) == 0 {
			value = value.Add(mappable.(mappables.Split).GetValue())
		}

		return false
	}
	// TODO test nil ID components
	collection.Iterate(key.NewKey(baseIDs.NewSplitID(nil, nil)), accumulator)

	return value
}
