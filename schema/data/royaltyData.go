package data

import "github.com/AssetMantle/modules/schema/ids"

type RoyaltyData interface {
	Data
	GetOwnableID() ids.OwnableID
	GetIdentityID() ids.IdentityID
	GetSplit() []DecData
}
