/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Trait = (*trait)(nil)

type trait struct {
	Property types.Property `json:"property"`
	Mutable  bool           `json:"mutable"`
}

func (trait trait) GetID() types.ID             { return trait.Property.GetID() }
func (trait trait) GetProperty() types.Property { return trait.Property }
func (trait trait) IsMutable() bool             { return trait.Mutable }

func NewTrait(property types.Property, mutable bool) types.Trait {
	return trait{
		Property: property,
		Mutable:  mutable,
	}
}
