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

func (identities identities) GetID() types.ID { return identities.ID }
func (identities identities) Get(id types.ID) types.InterIdentity {
	identityID := identityIDFromInterface(id)
	for _, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identityID) == 0 {
			return oldIdentity
		}
	}
	return nil
}
func (identities identities) GetList() []types.InterIdentity {
	return identities.List
}

func (identities identities) Fetch(id types.ID) types.InterIdentities {
	var identityList []types.InterIdentity
	identitiesID := identityIDFromInterface(id)
	if len(identitiesID.HashID.Bytes()) > 0 {
		identity := identities.mapper.read(identities.context, identitiesID)
		if identity != nil {
			identityList = append(identityList, identity)
		}
	} else {
		appendIdentityList := func(identity types.InterIdentity) bool {
			identityList = append(identityList, identity)
			return false
		}
		identities.mapper.iterate(identities.context, identitiesID, appendIdentityList)
	}
	identities.ID, identities.List = id, identityList
	return identities
}
func (identities identities) Add(identity types.InterIdentity) types.InterIdentities {
	identities.ID = readIdentityID("")
	identities.mapper.create(identities.context, identity)
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) < 0 {
			identities.List = append(append(identities.List[:i], identity), identities.List[i+1:]...)
			break
		}
	}
	return identities
}
func (identities identities) Remove(identity types.InterIdentity) types.InterIdentities {
	identities.mapper.delete(identities.context, identity.GetID())
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			identities.List = append(identities.List[:i], identities.List[i+1:]...)
			break
		}
	}
	return identities
}
func (identities identities) Mutate(identity types.InterIdentity) types.InterIdentities {
	identities.mapper.update(identities.context, identity)
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			identities.List[i] = identity
			break
		}
	}
	return identities
}

func NewIdentities(Mapper types.Mapper, context sdkTypes.Context) types.InterIdentities {
	switch mapper := Mapper.(type) {
	case identitiesMapper:
		return identities{
			ID:      readIdentityID(""),
			List:    []types.InterIdentity{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleRoute)))
	}

}
