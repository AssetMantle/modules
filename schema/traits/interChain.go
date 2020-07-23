package traits

import "github.com/persistenceOne/persistenceSDK/schema/types"

type InterChain interface {
	GetChainID() types.ID
}
