package trait

import "github.com/persistenceOne/persistenceSDK/types/schema"

type Lockable interface {
	CanSend(schema.Height) bool
	GetLock() schema.Height
}
