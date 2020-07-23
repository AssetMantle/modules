package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type identities struct {
	ID   types.ID
	List []mappables.InterIdentity

	mapper  utilities.Mapper
	context sdkTypes.Context
}

var _ mappers.InterIdentities = (*identities)(nil)

func (identities identities) GetID() types.ID { return identities.ID }
func (identities identities) Get(id types.ID) mappables.InterIdentity {
	identityID := identityIDFromInterface(id)
	for _, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identityID) == 0 {
			return oldIdentity
		}
	}
	return nil
}
func (identities identities) GetList() []mappables.InterIdentity {
	return identities.List
}

func (identities identities) Fetch(id types.ID) mappers.InterIdentities {
	var identityList []mappables.InterIdentity
	identitiesID := identityIDFromInterface(id)
	if len(identitiesID.HashID.Bytes()) > 0 {
		mappable := identities.mapper.Read(identities.context, identitiesID)
		if mappable != nil {
			identityList = append(identityList, mappable.(identity))
		}
	} else {
		appendIdentityList := func(mappable traits.Mappable) bool {
			identityList = append(identityList, mappable.(identity))
			return false
		}
		identities.mapper.Iterate(identities.context, identitiesID, appendIdentityList)
	}
	identities.ID, identities.List = id, identityList
	return identities
}
func (identities identities) Add(identity mappables.InterIdentity) mappers.InterIdentities {
	identities.ID = readIdentityID("")
	identities.mapper.Create(identities.context, identity)
	identities.List = append(identities.List, identity)
	return identities
}
func (identities identities) Remove(identity mappables.InterIdentity) mappers.InterIdentities {
	identities.mapper.Delete(identities.context, identity.GetID())
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			identities.List = append(identities.List[:i], identities.List[i+1:]...)
			break
		}
	}
	return identities
}
func (identities identities) Mutate(identity mappables.InterIdentity) mappers.InterIdentities {
	identities.mapper.Update(identities.context, identity)
	for i, oldIdentity := range identities.List {
		if oldIdentity.GetID().Compare(identity.GetID()) == 0 {
			identities.List[i] = identity
			break
		}
	}
	return identities
}

func NewIdentities(mapper utilities.Mapper, context sdkTypes.Context) mappers.InterIdentities {
	return identities{
		ID:      readIdentityID(""),
		List:    []mappables.InterIdentity{},
		mapper:  mapper,
		context: context,
	}
}
