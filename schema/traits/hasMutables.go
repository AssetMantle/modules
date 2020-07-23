package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type HasMutables interface {
	GetMutables() types.Mutables
}
