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

type mutables struct {
	lists.PropertyList
}

var _ qualified.Mutables = (*Mutables)(nil)

func (mutables *Mutables) GetMutablePropertyList() lists.List {
	if mutables.PropertyList == nil {
		return baseLists.NewPropertyList()
	}

	return mutables.PropertyList
}
func (mutables *Mutables) GetProperty(id ids.ID) properties.Property {
	return mutables.GetMutablePropertyList().(*baseLists.List).Impl.(lists.PropertyList).GetProperty(id)
}
func (mutables *Mutables) Mutate(propertyList ...properties.Property) qualified.Mutables {
	for _, property := range propertyList {
		mutables.PropertyList = mutables.PropertyList.Impl.(lists.PropertyList).Mutate(property).(*baseLists.List)
	}

	return mutables
}

func NewMutables(propertyList lists.List) qualified.Mutables {
	return &Mutables{
		PropertyList: propertyList.(*baseLists.List),
	}
}
