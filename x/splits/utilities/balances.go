// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/schema/ids"
)

func GetAllBalancesForIdentity(collection helpers.Collection, ownerID ids.IdentityID) helpers.Collection {
	return collection.IterateAll(func(record helpers.Record) bool {
		return key.GetSplitIDFromKey(record.GetKey()).GetOwnerID().Compare(ownerID) == 0
	})
}
