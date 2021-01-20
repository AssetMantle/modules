/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

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
