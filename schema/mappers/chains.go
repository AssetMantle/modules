package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Chains interface {
	GetID() types.ID
	Get(types.ID) mappables.Chain
	Add(mappables.Chain) Chains
	Remove(mappables.Chain) Chains
	Mutate(mappables.Chain) Chains
}
