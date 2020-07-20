package utility

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type QueryKeeper interface {
	Query(sdkTypes.Context, QueryRequest) QueryResponse
}
