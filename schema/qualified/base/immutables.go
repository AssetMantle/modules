// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/qualified"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type Immutables struct {
	lists.PropertyList
}

var _ qualified.Immutables = (*Immutables)(nil)

// TODO write test case
func (immutables Immutables) GetImmutablePropertyList() lists.PropertyList {
	if immutables.PropertyList.GetList() == nil {
		return base.NewPropertyList()
	}

	return immutables.PropertyList
}
func (immutables Immutables) GenerateHashID() ids.ID {
	metaList := make([]string, len(immutables.PropertyList.GetList()))

	for i, immutableProperty := range immutables.PropertyList.GetList() {
		metaList[i] = immutableProperty.GetHash().String()
	}

	return baseIDs.NewID(stringUtilities.Hash(metaList...))
}
