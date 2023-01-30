// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/schema/properties"
)

func IsDuplicate(propertyList []properties.AnyProperty) bool {
	propertyIDMap := map[string]bool{}

	var str = ""
	for _, property := range propertyList {
		str = property.GetID().GetKey().AsString() + ":" + property.GetID().GetType().AsString()
		if _, ok := propertyIDMap[str]; ok {
			return true
		}

		propertyIDMap[str] = true
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
