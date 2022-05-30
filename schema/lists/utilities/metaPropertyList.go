package utilities

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/utilities"
)

func ReadMetaProperties(metaPropertiesString string) (lists.MetaPropertyList, error) {
	var metaPropertyList []properties.MetaProperty

	metaProperties := strings.Split(metaPropertiesString, constants.PropertiesSeparator)
	for _, metaPropertyString := range metaProperties {
		if metaPropertyString != "" {
			metaProperty, err := utilities.ReadMetaProperty(metaPropertyString)
			if err != nil {
				return nil, err
			}

			metaPropertyList = append(metaPropertyList, metaProperty)
		}
	}

	return base.NewMetaProperties(metaPropertyList...), nil
}
