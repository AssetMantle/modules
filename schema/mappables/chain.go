package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Chain interface {
	String() string
	GetID() types.ID
	GetTrustHeight() types.Height
	traits.Mappable
}
