/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Property = (*Property)(nil)

func (property Property) GetID() types.ID     { return &property.Id }
func (property Property) GetFact() types.Fact { return &property.Fact }
func NewProperty(id types.ID, fact types.Fact) *Property {
	newID := *NewID(id.String())
	newFact := *NewFactProperty(fact.GetHashID(), fact.GetTypeID(), fact.GetSignatures())
	return &Property{
		Id:   newID,
		Fact: newFact,
	}
}
func ReadProperty(propertyString string) (types.Property, error) {
	property, Error := ReadMetaProperty(propertyString)
	if Error != nil {
		return nil, Error
	}

	return property.RemoveData(), nil
}
