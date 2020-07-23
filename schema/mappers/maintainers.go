package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Maintainers interface {
	ID() types.ID

	Get(types.ID) entities.Maintainer

	Add(entities.Maintainer) Maintainers
	Remove(entities.Maintainer) Maintainers
	Mutate(entities.Maintainer) Maintainers
}
