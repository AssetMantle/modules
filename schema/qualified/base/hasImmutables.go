// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	metaUtilities "github.com/AssetMantle/modules/utilities/meta"
)

type HasImmutables struct {
	Properties types.Properties `json:"properties"`
}

var _ qualified.HasImmutables = (*HasImmutables)(nil)

func (immutables HasImmutables) GetImmutableProperties() types.Properties {
	if immutables.Properties == nil {
		return baseTypes.NewProperties()
	}

	return immutables.Properties
}
func (immutables HasImmutables) GenerateHashID() types.ID {
	metaList := make([]string, len(immutables.Properties.GetList()))

	for i, immutableProperty := range immutables.Properties.GetList() {
		metaList[i] = immutableProperty.GetHash().String()
	}

	return baseIDs.NewID(metaUtilities.Hash(metaList...))
}
