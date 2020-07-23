package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type identities struct {
	ID      types.ID                 `json:"id" valid:"required~Enter the ID"`
	List    []entities.InterIdentity `json:"list" valid:"required~Enter the List"`
	mapper  identitiesMapper         `json:"mapper" valid:"required~Enter the Mapper"`
	context sdkTypes.Context         `json:"context" valid:"required~Enter the Context"`
}

var _ mappers.InterIdentities = (*identities)(nil)

func (identities identities) GetID() types.ID { return identities.ID }
func (identities identities) Get(id types.ID) entities.InterIdentity {
	identityID := identityIDFromInterface(id)
	for _, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identityID) == 0 {
			return oldIdentity
		}
	}
	return nil
}
func (identities identities) GetList() []entities.InterIdentity {
	return identities.List
}

func (identities identities) Fetch(id types.ID) mappers.InterIdentities {
	var identityList []entities.InterIdentity
	identitiesID := identityIDFromInterface(id)
	if len(identitiesID.HashID.Bytes()) > 0 {
		identity := identities.mapper.read(identities.context, identitiesID)
		if identity != nil {
			identityList = append(identityList, identity)
		}
	} else {
		appendIdentityList := func(identity entities.InterIdentity) bool {
			identityList = append(identityList, identity)
			return false
		}
		identities.mapper.iterate(identities.context, identitiesID, appendIdentityList)
	}
	identities.ID, identities.List = id, identityList
	return identities
}
func (identities identities) Add(identity entities.InterIdentity) mappers.InterIdentities {
	identities.ID = readIdentityID("")
	identities.mapper.create(identities.context, identity)
	identities.List = append(identities.List, identity)
	return identities
}
func (identities identities) Remove(identity entities.InterIdentity) mappers.InterIdentities {
	identities.mapper.delete(identities.context, identity.GetID())
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			identities.List = append(identities.List[:i], identities.List[i+1:]...)
			break
		}
	}
	return identities
}
func (identities identities) Mutate(identity entities.InterIdentity) mappers.InterIdentities {
	identities.mapper.update(identities.context, identity)
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			identities.List[i] = identity
			break
		}
	}
	return identities
}

func NewIdentities(Mapper utilities.Mapper, context sdkTypes.Context) mappers.InterIdentities {
	switch mapper := Mapper.(type) {
	case identitiesMapper:
		return identities{
			ID:      readIdentityID(""),
			List:    []entities.InterIdentity{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleRoute)))
	}

}
