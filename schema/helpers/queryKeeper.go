package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type QueryKeeper interface {
	Enquire(sdkTypes.Context, QueryRequest) QueryResponse
}
