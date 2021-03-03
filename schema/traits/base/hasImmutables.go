/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type HasImmutables struct {
	Properties types.Properties `json:"properties"`
}

var _ traits.HasImmutables = (*HasImmutables)(nil)

func (immutables HasImmutables) GetImmutableProperties() types.Properties {
	return immutables.Properties
}
func (immutables HasImmutables) GenerateHashID() types.ID {
	metaList := make([]string, len(immutables.Properties.GetList()))

	for i, immutableProperty := range immutables.Properties.GetList() {
		metaList[i] = immutableProperty.GetFact().GetHashID().String()
	}

	return base.NewID(metaUtilities.Hash(metaList...))
}
