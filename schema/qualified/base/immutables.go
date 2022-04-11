// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseTypes "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
	metaUtilities "github.com/AssetMantle/modules/utilities/meta"
)

type Immutables struct {
	Properties lists.PropertyList `json:"properties"`
}

var _ qualified.Immutables = (*Immutables)(nil)

func (immutables Immutables) GetImmutablePropertyList() lists.PropertyList {
	if immutables.Properties == nil {
		return baseTypes.NewPropertyList()
	}

	return immutables.Properties
}
func (immutables Immutables) GenerateHashID() types.ID {
	metaList := make([]string, len(immutables.Properties.GetList()))

	for i, immutableProperty := range immutables.Properties.GetList() {
		metaList[i] = immutableProperty.GetHash().String()
	}

	return baseIDs.NewID(metaUtilities.Hash(metaList...))
}
