/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type HasMutables struct {
	Properties types.Properties `json:"properties"`
}

var _ traits.HasMutables = (*HasMutables)(nil)

func (mutables HasMutables) GetMutableProperties() types.Properties {
	return mutables.Properties
}
func (mutables HasMutables) Mutate(propertyList ...types.Property) traits.HasMutables {
	for _, property := range propertyList {
		mutables.Properties = mutables.Properties.Mutate(property)
	}

	return mutables
}
