package property

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func Duplicate(propertyList []types.Property) bool {
	propertyIDMap := map[types.ID]bool{}
	for _, property := range propertyList {
		if _, ok := propertyIDMap[property.GetID()]; ok {
			return true
		}
		propertyIDMap[property.GetID()] = true
	}
	return false
}
