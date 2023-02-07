package lists

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
)

type PropertyList interface {
	Size() int
	GetProperty(ids.PropertyID) properties.Property
	GetList() []properties.Property
	GetPropertyIDList() IDList

	Add(...properties.Property) PropertyList
	Remove(...properties.Property) PropertyList
	Mutate(...properties.Property) PropertyList
	HasRepeats() bool
	ScrubData() PropertyList
}
