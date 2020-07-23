package traits

import "github.com/persistenceOne/persistenceSDK/schema/types"

type Lockable interface {
	CanSend(types.Height) bool
	GetLock() types.Height
}
