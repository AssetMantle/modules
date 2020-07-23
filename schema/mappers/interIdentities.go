package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterIdentities interface {
	GetID() types.ID

	Get(types.ID) mappables.InterIdentity
	GetList() []mappables.InterIdentity

	Fetch(types.ID) InterIdentities
	Add(mappables.InterIdentity) InterIdentities
	Remove(mappables.InterIdentity) InterIdentities
	Mutate(mappables.InterIdentity) InterIdentities
}
