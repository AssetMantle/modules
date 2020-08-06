/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Property = (*property)(nil)

type property struct {
	ID   types.ID   `json:"id"`
	Fact types.Fact `json:"fact"`
}

func (property property) String() string {
	bytes, Error := json.Marshal(property)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (property property) GetID() types.ID     { return property.ID }
func (property property) GetFact() types.Fact { return property.Fact }
func NewProperty(id types.ID, fact types.Fact) types.Property {
	return property{
		ID:   id,
		Fact: fact,
	}
}
