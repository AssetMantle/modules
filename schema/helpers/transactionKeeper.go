package helpers

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type TransactionKeeper interface {
	Transact(sdkTypes.Context, sdkTypes.Msg) error
}
