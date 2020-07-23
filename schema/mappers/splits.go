package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Splits interface {
	GetID() types.ID

	Get(types.ID) entities.Split
	GetList() []entities.Split

	Fetch(types.ID) Splits
	Add(entities.Split) Splits
	Remove(entities.Split) Splits
	Mutate(entities.Split) Splits
}
