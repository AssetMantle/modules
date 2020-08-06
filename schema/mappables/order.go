package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Order interface {
	GetID() types.ID
	traits.InterChain
	traits.HasMutables
	traits.HasImmutables
	traits.Mappable
}
