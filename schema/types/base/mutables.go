/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type mutables struct {
	Properties types.Properties `json:"properties"`
}

var _ types.Mutables = (*mutables)(nil)

func (mutables mutables) Get() types.Properties {
	return mutables.Properties
}
func (mutables mutables) Mutate(propertyList []types.Property) types.Mutables {
	for _, property := range propertyList {
		mutables.Properties.Mutate(property)
	}
	return mutables
}
func NewMutables(properties types.Properties) types.Mutables {
	return mutables{
		Properties: properties,
	}
}
