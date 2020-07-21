package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type identities struct {
	ID   schema.ID
	List []schema.InterIdentity

	mapper  identitiesMapper
	context sdkTypes.Context
}

var _ schema.InterIdentities = (*identities)(nil)

func (identities identities) GetID() schema.ID { return identities.ID }
func (identities identities) Get(id schema.ID) schema.InterIdentity {
	identityID := identityIDFromInterface(id)
	for _, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identityID) == 0 {
			return oldIdentity
		}
	}
	return nil
}
func (identities identities) GetList() []schema.InterIdentity {
	return identities.List
}

func (identities identities) Fetch(id schema.ID) schema.InterIdentities {
	var identityList []schema.InterIdentity
	identitiesID := identityIDFromInterface(id)
	if len(identitiesID.HashID.Bytes()) > 0 {
		identity := identities.mapper.read(identities.context, identitiesID)
		if identity != nil {
			identityList = append(identityList, identity)
		}
	} else {
		appendIdentityList := func(identity schema.InterIdentity) bool {
			identityList = append(identityList, identity)
			return false
		}
		identities.mapper.iterate(identities.context, identitiesID, appendIdentityList)
	}
	identities.ID, identities.List = id, identityList
	return identities
}
func (identities identities) Add(identity schema.InterIdentity) schema.InterIdentities {
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
func (identities identities) Remove(identity schema.InterIdentity) schema.InterIdentities {
	identities.mapper.delete(identities.context, identity.GetID())
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			identities.List = append(identities.List[:i], identities.List[i+1:]...)
			break
		}
	}
	return identities
}
func (identities identities) Mutate(identity schema.InterIdentity) schema.InterIdentities {
	identities.mapper.update(identities.context, identity)
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			identities.List[i] = identity
			break
		}
	}
	return identities
}

func NewIdentities(Mapper utility.Mapper, context sdkTypes.Context) schema.InterIdentities {
	switch mapper := Mapper.(type) {
	case identitiesMapper:
		return identities{
			ID:      readIdentityID(""),
			List:    []schema.InterIdentity{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleRoute)))
	}

}
