package utilities

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data/utilities"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func ReadMetaProperty(metaPropertyString string) (properties.MetaProperty, error) {
	propertyIDAndData := strings.Split(metaPropertyString, constants.PropertyIDAndDataSeparator)
	if len(propertyIDAndData) == 2 && propertyIDAndData[0] != "" {
		data, Error := utilities.ReadData(propertyIDAndData[1])
		if Error != nil {
			return nil, Error
		}

		return baseProperties.NewMetaProperty(baseIDs.NewID(propertyIDAndData[0]), data), nil
	}

	return nil, errors.IncorrectFormat
}
