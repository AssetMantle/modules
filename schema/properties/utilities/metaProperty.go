package utilities

import (
	"github.com/AssetMantle/modules/schema/data/utlities"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func ReadMetaProperty(metaPropertyString string) (properties.MetaProperty, error) {
	propertyIDString, dataString := stringUtilities.SplitMetaProperty(metaPropertyString)
	if propertyIDString != "" {
		data, err := utlities.ReadData(dataString)
		if err != nil {
			return nil, err
		}

		return baseProperties.NewMetaProperty(baseIDs.NewID(propertyIDString), data), nil
	}

	return nil, errorConstants.IncorrectFormat
}
