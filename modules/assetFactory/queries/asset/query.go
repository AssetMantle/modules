package asset

import (
	"github.com/persistenceOne/persistenceSDK/types"
)

type request struct {
	ID types.ID
}

var _ types.QueryRequest = (*request)(nil)
