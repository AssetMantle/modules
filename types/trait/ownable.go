package trait

import "github.com/persistenceOne/persistenceSDK/types/schema"

type Ownable interface {
	GetOwnerID() schema.ID
	GetOwnableID() schema.ID
}
