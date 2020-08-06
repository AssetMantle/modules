package base

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Trait = (*trait)(nil)

type trait struct {
	ID       types.ID       `json:"id"`
	Property types.Property `json:"property"`
	Mutable  bool           `json:"mutable"`
}

func (trait trait) String() string {
	bytes, Error := json.Marshal(trait)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (trait trait) GetID() types.ID             { return trait.ID }
func (trait trait) GetProperty() types.Property { return trait.Property }
func (trait trait) IsMutable() bool             { return trait.Mutable }

func NewTrait(ID types.ID, property types.Property, mutable bool) types.Trait {
	return trait{
		ID:       ID,
		Property: property,
		Mutable:  mutable,
	}
}
