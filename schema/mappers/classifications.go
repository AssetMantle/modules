package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Classifications interface {
	GetID() types.ID

	Get(types.ID) mappables.Classification
	GetList() []mappables.Classification

	Fetch(types.ID) Classifications
	Add(mappables.Classification) Classifications
	Remove(mappables.Classification) Classifications
	Mutate(mappables.Classification) Classifications
}
