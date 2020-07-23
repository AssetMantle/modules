package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Chains interface {
	GetID() types.ID
	Get(types.ID) entities.Chain
	Add(entities.Chain) Chains
	Remove(entities.Chain) Chains
	Mutate(entities.Chain) Chains
}
