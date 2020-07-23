package traits

import "github.com/persistenceOne/persistenceSDK/schema/types"

type Ownable interface {
	GetOwnerID() types.ID
	GetOwnableID() types.ID
}
