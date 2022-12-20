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

//type immutables struct {
//	lists.PropertyList
//}

var _ qualified.Immutables = (*Immutables)(nil)

// TODO write test case
func (immutables Immutables) GetImmutablePropertyList() lists.List {
	if immutables.PropertyList.Impl.(lists.PropertyList).GetList() == nil {
		return baseLists.NewPropertyList()
	}

	return immutables.PropertyList
}
func (immutables Immutables) GetProperty(id ids.ID) properties.Property {
	return immutables.GetImmutablePropertyList().(*baseLists.List).Impl.(lists.PropertyList).GetProperty(id)
}
func (immutables Immutables) GenerateHashID() ids.ID {
	metaList := make([][]byte, len(immutables.PropertyList.Impl.(lists.PropertyList).GetList()))

	for i, immutableProperty := range immutables.PropertyList.Impl.(lists.PropertyList).GetList() {
		metaList[i] = immutableProperty.GetDataID().(ids.DataID).GetHashID().Bytes()
	}

	return baseIDs.GenerateHashID(metaList...)
}

func NewImmutables(propertyList lists.List) qualified.Immutables {
	return &Immutables{
		PropertyList: propertyList.(*baseLists.List),
	}
}
