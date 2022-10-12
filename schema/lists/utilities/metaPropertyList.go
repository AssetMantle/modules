package utilities

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/utilities"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func ReadMetaPropertyList(metaPropertiesString string) (lists.MetaPropertyList, error) {
	var metaPropertyList []properties.MetaProperty

	metaProperties := stringUtilities.SplitListString(metaPropertiesString)
	for _, metaPropertyString := range metaProperties {
		if metaPropertyString != "" {
			metaProperty, err := utilities.ReadMetaProperty(metaPropertyString)
			if err != nil {
				return nil, err
			}

			metaPropertyList = append(metaPropertyList, metaProperty)
		}
	}

	return base.NewMetaPropertyList(metaPropertyList...), nil
}
