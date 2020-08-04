package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type traits struct {
	TraitList []types.Trait
}

var _ types.Traits = (*traits)(nil)

func (traits traits) Get(id types.ID) types.Trait {
	for _, trait := range traits.GetList() {
		if trait.GetID().Compare(id) == 0 {
			return trait
		}
	}
	return nil
}
func (traits traits) GetList() []types.Trait {
	var traitList []types.Trait
	for _, baseTrait := range traits.TraitList {
		traitList = append(traitList, baseTrait)
	}
	return traitList
}
func (traits traits) Add(trait types.Trait) types.Traits {
	traitList := traits.GetList()
	traitList = append(traitList, trait)
	return NewTraits(traitList)
}
func (traits traits) Remove(trait types.Trait) types.Traits {
	traitList := traits.GetList()
	for i, oldTrait := range traitList {
		if oldTrait.GetID().Compare(trait.GetID()) == 0 {
			traitList = append(traitList[:i], traitList[i+1:]...)
		}
	}
	return NewTraits(traitList)
}
func (traits traits) Mutate(trait types.Trait) types.Traits {
	traitList := traits.GetList()
	for i, oldTrait := range traitList {
		if oldTrait.GetID().Compare(trait.GetID()) == 0 {
			traitList[i] = trait
		}
	}
	return NewTraits(traitList)
}
func NewTraits(traitList []types.Trait) types.Traits {
	return traits{
		TraitList: traitList,
	}
}
