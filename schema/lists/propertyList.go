package lists

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
)

type PropertyList interface {
	GetProperty(ids.PropertyID) properties.Property
	GetList() []*base.Property
	GetPropertyIDList() IDList

	Add(...properties.Property) PropertyList
	Remove(...properties.Property) PropertyList
	Mutate(...properties.Property) PropertyList

	ScrubData() PropertyList
}
