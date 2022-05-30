// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
)

func Duplicate(propertyList []properties.Property) bool {
	propertyIDMap := map[ids.ID]bool{}

	for _, property := range propertyList {
		if _, ok := propertyIDMap[property.GetID()]; ok {
			return true
		}

		propertyIDMap[property.GetID()] = true
	}

	return false
}
