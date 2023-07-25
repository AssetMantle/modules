package utilities

import (
	"github.com/AssetMantle/schema/go/ids"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/mappable"
)

func GetAllBalancesForIdentity(collection helpers.Collection, identityID ids.IdentityID) helpers.Collection {
	return collection.IterateAll(func(Mappable helpers.Mappable) bool {
		return mappable.GetSplit(Mappable).GetOwnerID().Compare(identityID) == 0
	})
}
