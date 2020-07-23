package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Burnable interface {
	CanBurn(types.Height) bool
	GetBurn() types.Height
}
