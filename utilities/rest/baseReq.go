package rest

import (
	"github.com/cosmos/cosmos-sdk/types"
	"net/http"
)

type BaseReq struct {
	From          string         `json:"from"`
	Memo          string         `json:"memo"`
	TimeoutHeight uint64         `json:"timeout_height"`
	AccountNumber uint64         `json:"account_number"`
	Sequence      uint64         `json:"sequence"`
	ChainID       string         `json:"chain_id"`
	Gas           string         `json:"gas"`
	Fees          types.Coins    `json:"fees"`
	GasPrices     types.DecCoins `json:"gas_prices"`
	Simulate      bool           `json:"simulate"`
	GasAdjustment string         `json:"gas_adjustment"`
}

func (BaseReq) Validate() error {
	return nil
}

func (BaseReq) Sanitize() BaseReq {
	return BaseReq{}
}

func (BaseReq) ValidateBasic(w http.ResponseWriter) bool {
	return false
}

func NewBaseReq(from, memo, chainID, gas, gasAdjustment string, timeoutHeight, accountNumber, sequence uint64, fees types.Coins, gasPrices types.DecCoins, simulate bool) BaseReq {
	return BaseReq{
		From:          from,
		Memo:          memo,
		TimeoutHeight: timeoutHeight,
		AccountNumber: accountNumber,
		Sequence:      sequence,
		ChainID:       chainID,
		Gas:           gas,
		Fees:          fees,
		GasPrices:     gasPrices,
		Simulate:      simulate,
		GasAdjustment: gasAdjustment,
	}
}
