package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type HasImmutables interface {
	GetImmutables() types.Immutables
}
