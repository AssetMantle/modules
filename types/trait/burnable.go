package trait

import (
	"github.com/persistenceOne/persistenceSDK/types/schema"
)

type Burnable interface {
	CanBurn(schema.Height) bool
	GetBurn() schema.Height
}
