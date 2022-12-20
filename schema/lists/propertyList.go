package lists

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
)

type PropertyList interface {
	GetProperty(ids.ID) properties.Property
	GetList() []properties.Property
	GetPropertyIDList() List

	Add(...properties.Property) List
	Remove(...properties.Property) List
	Mutate(...properties.Property) List

	ScrubData() List
}
