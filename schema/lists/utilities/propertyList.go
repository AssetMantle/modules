package utilities

import (
	"github.com/AssetMantle/modules/schema/lists"
)

func ReadProperties(propertiesString string) (lists.PropertyList, error) {
	Properties, err := ReadMetaProperties(propertiesString)
	if err != nil {
		return nil, err
	}

	return Properties.ToPropertyList(), nil
}
