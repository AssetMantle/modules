/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type immutables struct {
	Properties types.Properties `json:"properties"`
}

var _ types.Immutables = (*immutables)(nil)

func (immutables immutables) Get() types.Properties {
	return immutables.Properties
}
func (immutables immutables) GetHashID() types.ID {
	var facts []string
	for _, immutableProperty := range immutables.Properties.GetList() {
		facts = append(facts, immutableProperty.GetFact().GetHash())
	}
	return NewID(metaUtilities.Hash(facts...))
}
func NewImmutables(properties types.Properties) types.Immutables {
	return immutables{Properties: properties}
}
