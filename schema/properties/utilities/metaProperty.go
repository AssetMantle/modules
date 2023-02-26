package utilities

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
)

func ReadMetaProperty(metaPropertyString string) (properties.MetaProperty, error) {
	propertyIDString, dataString := SplitMetaProperty(metaPropertyString)
	if propertyIDString != "" {
		data, err := base.PrototypeAnyData().FromString(dataString)
		if err != nil {
			return nil, err
		}

		return baseProperties.NewMetaProperty(baseIDs.NewStringID(propertyIDString), data), nil
	}

	return nil, errorConstants.IncorrectFormat.Wrapf("propertyID missing from metaPropertyString %s", metaPropertyString)
}

func SplitMetaProperty(metaPropertyString string) (propertyIDString, dataString string) {
	if propertyIDAndData := strings.Split(metaPropertyString, constants.PropertyIDAndDataSeparator); len(propertyIDAndData) < 2 {
		return "", ""
	} else {
		return propertyIDAndData[0], propertyIDAndData[1]
	}

}
