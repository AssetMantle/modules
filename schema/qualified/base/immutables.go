// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ qualified.Immutables = (*Immutables)(nil)

// TODO write test case
func (immutables Immutables) GetImmutablePropertyList() lists.PropertyList {
	if immutables.PropertyList.GetList() == nil {
		return baseLists.NewPropertyList()
	}

	return immutables.PropertyList
}
func (immutables Immutables) GetProperty(id ids.ID) properties.Property {
	return immutables.GetImmutablePropertyList().GetProperty(id)
}
func (immutables Immutables) GenerateHashID() ids.ID {
	metaList := make([][]byte, len(immutables.PropertyList.GetList()))

	for i, immutableProperty := range immutables.PropertyList.GetList() {
		metaList[i] = immutableProperty.GetDataID().(ids.DataID).GetHashID().Bytes()
	}

	return baseIDs.GenerateHashID(metaList...)
}

func NewImmutables(propertyList lists.PropertyList) qualified.Immutables {
	return &Immutables{
		PropertyList: propertyList.(*baseLists.PropertyList),
	}
}
