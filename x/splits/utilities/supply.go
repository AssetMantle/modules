// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func GetTotalSupply(collection helpers.Collection, assetID ids.AssetID) math.Int {
	value := sdkTypes.ZeroInt()

	collection.Iterate(key.NewKey(baseIDs.NewSplitID(assetID, baseIDs.PrototypeIdentityID())), func(record helpers.Record) bool {
		value = value.Add(mappable.GetSplit(record.GetMappable()).GetValue())
		return false
	})

	return value
}
