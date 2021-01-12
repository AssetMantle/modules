package properties

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func Duplicate(properties types.Properties) bool {
	propertyIDMap := map[types.ID]bool{}
	for _, property := range properties.GetList() {
		if _, ok := propertyIDMap[property.GetID()]; ok {
			return true
		}
		propertyIDMap[property.GetID()] = true
	}
	return false
}
