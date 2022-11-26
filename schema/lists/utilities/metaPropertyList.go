package utilities

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/utilities"
)

func ReadMetaPropertyList(metaPropertiesString string) (lists.PropertyList, error) {
	var Properties []properties.Property

	metaProperties := SplitListString(metaPropertiesString)
	for _, metaPropertyString := range metaProperties {
		if metaPropertyString != "" {
			metaProperty, err := utilities.ReadMetaProperty(metaPropertyString)
			if err != nil {
				return nil, err
			}

			Properties = append(Properties, metaProperty)
		}
	}

	return base.NewPropertyList(Properties...), nil
}
