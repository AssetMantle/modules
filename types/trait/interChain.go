package trait

import "github.com/persistenceOne/persistenceSDK/types/schema"

type InterChain interface {
	GetChainID() schema.ID
}
