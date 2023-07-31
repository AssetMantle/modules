package utilities

import (
	"github.com/AssetMantle/schema/go/ids"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/mappable"
)

func GetAllBalancesForIdentity(collection helpers.Collection, identityID ids.IdentityID) helpers.Collection {
	return collection.IterateAll(func(record helpers.Record) bool {
		return mappable.GetSplit(record.GetMappable()).GetOwnerID().Compare(identityID) == 0
	})
}
