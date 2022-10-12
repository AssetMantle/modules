// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

type mutables struct {
	lists.PropertyList
}

var _ qualified.Mutables = (*mutables)(nil)

func (mutables mutables) GetMutablePropertyList() lists.PropertyList {
	if mutables.PropertyList == nil {
		return baseLists.NewPropertyList()
	}

	return mutables.PropertyList
}
func (mutables mutables) Mutate(propertyList ...properties.Property) qualified.Mutables {
	for _, property := range propertyList {
		mutables.PropertyList = mutables.PropertyList.Mutate(property)
	}

	return mutables
}

func NewMutables(propertyList lists.PropertyList) qualified.Mutables {
	return mutables{
		PropertyList: propertyList,
	}
}
