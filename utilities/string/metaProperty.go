package string

import (
	"strings"
)

func SplitMetaProperty(metaPropertyString string) (propertyIDString, dataString string) {
	propertyIDAndData := strings.Split(metaPropertyString, propertyIDAndDataSeparator)
	if len(propertyIDAndData) == 2 {
		return propertyIDAndData[0], propertyIDAndData[1]
	}

	return "", ""
}
