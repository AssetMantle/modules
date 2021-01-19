/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Property = (*property)(nil)

type property struct {
	ID   types.ID   `json:"id"`
	Fact types.Fact `json:"fact"`
}

func (property property) GetID() types.ID     { return property.ID }
func (property property) GetFact() types.Fact { return property.Fact }
func NewProperty(id types.ID, fact types.Fact) types.Property {
	return property{
		ID:   id,
		Fact: fact,
	}
}
func ReadProperty(propertyString string) (types.Property, error) {
	property, Error := ReadMetaProperty(propertyString)
	if Error != nil {
		return nil, Error
	}

	return property.RemoveData(), nil
}
