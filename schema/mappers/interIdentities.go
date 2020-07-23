package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterIdentities interface {
	GetID() types.ID

	Get(types.ID) entities.InterIdentity
	GetList() []entities.InterIdentity

	Fetch(types.ID) InterIdentities
	Add(entities.InterIdentity) InterIdentities
	Remove(entities.InterIdentity) InterIdentities
	Mutate(entities.InterIdentity) InterIdentities
}
