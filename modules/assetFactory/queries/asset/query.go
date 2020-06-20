package asset

import (
	"github.com/persistenceOne/persistenceSDK/types"
)

type query struct {
	ID types.ID
}

var _ types.Query = (*query)(nil)
