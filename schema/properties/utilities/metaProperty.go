package utilities

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data/utlities"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func ReadMetaProperty(metaPropertyString string) (properties.MetaProperty, error) {
	propertyIDAndData := strings.Split(metaPropertyString, stringUtilities.PropertyIDAndDataSeparator)
	if len(propertyIDAndData) == 2 && propertyIDAndData[0] != "" {
		data, err := utlities.ReadData(propertyIDAndData[1])
		if err != nil {
			return nil, err
		}

		return baseProperties.NewMetaProperty(baseIDs.NewID(propertyIDAndData[0]), data), nil
	}

	return nil, errorConstants.IncorrectFormat
}
