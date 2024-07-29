package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"net/http"
)

type CommonTransactionRequest interface {
	GetFrom() string
	SetFrom(string) CommonTransactionRequest
	GetChainID() string
	SetChainID(string) CommonTransactionRequest
	GetFees() sdkTypes.Coins
	IsSimulated() bool
	GetGasPrices() sdkTypes.DecCoins
	GetTimeoutHeight() uint64
	GetSequence() uint64
	GetMemo() string
	GetGas() string
	GetGasAdjustment() string
	GetAccountNumber() uint64
	SetAccountNumber(uint64) CommonTransactionRequest
	Validate() error
	Sanitize() CommonTransactionRequest
	ValidateBasic(responseWriter http.ResponseWriter) bool
}
