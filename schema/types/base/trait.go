package base

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Trait = (*trait)(nil)

type trait struct {
	ID       types.ID
	Property types.Property
	mutable  bool
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
func (trait trait) IsMutable() bool             { return trait.mutable }
func NewTrait(id types.ID, property types.Property) types.Trait {
	return trait{
		ID:       id,
		Property: property,
	}
}
