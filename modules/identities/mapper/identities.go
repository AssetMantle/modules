package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type identities struct {
	ID   types.ID
	List []types.InterIdentity

	mapper  identitiesMapper
	context sdkTypes.Context
}

var _ types.InterIdentities = (*identities)(nil)

func (Identities identities) GetID() types.ID { return Identities.ID }
func (Identities identities) Get(id types.ID) types.InterIdentity {
	identityID := identityIDFromInterface(id)
	for _, oldIdentity := range Identities.List {
		if oldIdentity.GetID().Compare(identityID) == 0 {
			return oldIdentity
		}
	}
	return nil
}
func (Identities identities) GetList() []types.InterIdentity {
	return Identities.List
}

func (Identities identities) Fetch(id types.ID) types.InterIdentities {
	var identityList []types.InterIdentity
	identitiesID := identityIDFromInterface(id)
	if len(identitiesID.HashID.Bytes()) > 0 {
		identity := Identities.mapper.read(Identities.context, identitiesID)
		if identity != nil {
			identityList = append(identityList, identity)
		}
	} else {
		appendIdentityList := func(identity types.InterIdentity) bool {
			identityList = append(identityList, identity)
			return false
		}
		Identities.mapper.iterate(Identities.context, identitiesID, appendIdentityList)
	}
	return identities{id, identityList, Identities.mapper, Identities.context}
}
func (Identities identities) Add(identity types.InterIdentity) types.InterIdentities {
	Identities.ID = nil
	Identities.mapper.create(Identities.context, identity)
	for i, oldIdentity := range Identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) < 0 {
			Identities.List = append(append(Identities.List[:i], identity), Identities.List[i+1:]...)
			break
		}
	}
	return Identities
}
func (Identities identities) Remove(identity types.InterIdentity) types.InterIdentities {
	Identities.mapper.delete(Identities.context, identity.GetID())
	for i, oldIdentity := range Identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			Identities.List = append(Identities.List[:i], Identities.List[i+1:]...)
			break
		}
	}
	return Identities
}
func (Identities identities) Mutate(identity types.InterIdentity) types.InterIdentities {
	Identities.mapper.update(Identities.context, identity)
	for i, oldIdentity := range Identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			Identities.List[i] = identity
			break
		}
	}
	return Identities
}

func NewIdentities(Mapper types.Mapper, context sdkTypes.Context) types.InterIdentities {
	switch mapper := Mapper.(type) {
	case identitiesMapper:
		return identities{
			ID:      nil,
			List:    nil,
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleRoute)))
	}

}
