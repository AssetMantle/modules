// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/qualified"
)

type immutables struct {
	lists.PropertyList
}

var _ qualified.Immutables = (*immutables)(nil)

// TODO write test case
func (immutables immutables) GetImmutablePropertyList() lists.PropertyList {
	if immutables.PropertyList.GetList() == nil {
		return base.NewPropertyList()
	}

	return immutables.PropertyList
}
func (immutables immutables) GenerateHashID() ids.HashID {
	metaList := make([]string, len(immutables.PropertyList.GetList()))

	for i, immutableProperty := range immutables.PropertyList.GetList() {
		metaList[i] = immutableProperty.GetHash().String()
	}

	return baseIDs.GenerateHashID(metaList...)
}

func NewImmutables(propertyList lists.PropertyList) qualified.Immutables {
	return immutables{
		PropertyList: propertyList,
	}
}
