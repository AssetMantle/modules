package traits

import "github.com/persistenceOne/persistenceSDK/schema/types"

type Mappable interface {
	GetID() types.ID
}
