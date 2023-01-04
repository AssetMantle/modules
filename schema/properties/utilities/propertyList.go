// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
)

func IsDuplicate(propertyList []properties.AnyProperty) bool {
	propertyIDMap := map[ids.ID]bool{}

	for _, property := range propertyList {
		if _, ok := propertyIDMap[property.GetID()]; ok {
			return true
		}

		propertyIDMap[property.GetID()] = true
	}

	return false
}

func AnyPropertyListToPropertyList(anyPropertyList ...properties.AnyProperty) []properties.Property {
	propertyList := make([]properties.Property, len(anyPropertyList))

	for i, property := range anyPropertyList {
		propertyList[i] = property.(properties.Property)
	}

	return propertyList
}
