package trait

import (
	"github.com/persistenceOne/persistenceSDK/types/schema"
)

type HasImmutables interface {
	GetImmutables() schema.Immutables
}
