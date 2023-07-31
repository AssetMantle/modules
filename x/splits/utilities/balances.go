package utilities

import (
	"github.com/AssetMantle/schema/go/ids"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
)

func GetAllBalancesForIdentity(collection helpers.Collection, ownerID ids.IdentityID) helpers.Collection {
	return collection.IterateAll(func(record helpers.Record) bool {
		return key.GetSplitIDFromKey(record.GetKey()).GetOwnerID().Compare(ownerID) == 0
	})
}
