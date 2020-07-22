package trait

import (
	"github.com/persistenceOne/persistenceSDK/types/schema"
)

type HasMutables interface {
	GetMutables() schema.Mutables
}
