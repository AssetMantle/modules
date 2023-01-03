// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ qualified.Mutables = (*Mutables)(nil)

func (mutables *Mutables) GetMutablePropertyList() lists.PropertyList {
	if mutables.PropertyList == nil {
		return baseLists.NewPropertyList()
	}

	return mutables.PropertyList
}
func (mutables *Mutables) GetProperty(id ids.PropertyID) properties.AnyProperty {
	return mutables.GetMutablePropertyList().GetProperty(id)
}
func (mutables *Mutables) Mutate(propertyList ...properties.Property) qualified.Mutables {
	for _, property := range propertyList {
		mutables.PropertyList = mutables.PropertyList.Mutate(property).(*baseLists.PropertyList)
	}

	return mutables
}

func NewMutables(propertyList lists.PropertyList) qualified.Mutables {
	return &Mutables{
		PropertyList: propertyList.(*baseLists.PropertyList),
	}
}
