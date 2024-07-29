package base

import (
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/rest"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"net/http"
	"strings"
)

type commonTransactionRequest struct {
	From          string            `json:"from"`
	Memo          string            `json:"memo"`
	TimeoutHeight uint64            `json:"timeout_height"`
	AccountNumber uint64            `json:"account_number"`
	Sequence      uint64            `json:"sequence"`
	ChainID       string            `json:"chain_id"`
	Gas           string            `json:"gas"`
	Fees          sdkTypes.Coins    `json:"fees"`
	GasPrices     sdkTypes.DecCoins `json:"gas_prices"`
	Simulate      bool              `json:"simulate"`
	GasAdjustment string            `json:"gas_adjustment"`
}

var _ helpers.CommonTransactionRequest = commonTransactionRequest{}

func (commonTransactionRequest commonTransactionRequest) SetFrom(from string) helpers.CommonTransactionRequest {
	commonTransactionRequest.From = from
	return commonTransactionRequest
}
func (commonTransactionRequest commonTransactionRequest) SetAccountNumber(accountNumber uint64) helpers.CommonTransactionRequest {
	commonTransactionRequest.AccountNumber = accountNumber
	return commonTransactionRequest
}
func (commonTransactionRequest commonTransactionRequest) GetAccountNumber() uint64 {
	return commonTransactionRequest.AccountNumber
}
func (commonTransactionRequest commonTransactionRequest) GetGas() string {
	return commonTransactionRequest.Gas
}
func (commonTransactionRequest commonTransactionRequest) GetGasAdjustment() string {
	return commonTransactionRequest.GasAdjustment
}
func (commonTransactionRequest commonTransactionRequest) GetMemo() string {
	return commonTransactionRequest.Memo
}
func (commonTransactionRequest commonTransactionRequest) GetSequence() uint64 {
	return commonTransactionRequest.Sequence
}
func (commonTransactionRequest commonTransactionRequest) GetTimeoutHeight() uint64 {
	return commonTransactionRequest.TimeoutHeight
}
func (commonTransactionRequest commonTransactionRequest) GetGasPrices() sdkTypes.DecCoins {
	return commonTransactionRequest.GasPrices
}
func (commonTransactionRequest commonTransactionRequest) IsSimulated() bool {
	return commonTransactionRequest.Simulate
}
func (commonTransactionRequest commonTransactionRequest) GetFrom() string {
	return commonTransactionRequest.From
}
func (commonTransactionRequest commonTransactionRequest) GetChainID() string {
	return commonTransactionRequest.ChainID
}
func (commonTransactionRequest commonTransactionRequest) SetChainID(chainIDString string) helpers.CommonTransactionRequest {
	commonTransactionRequest.ChainID = chainIDString
	return commonTransactionRequest
}
func (commonTransactionRequest commonTransactionRequest) GetFees() sdkTypes.Coins {
	return commonTransactionRequest.Fees
}
func (commonTransactionRequest commonTransactionRequest) Validate() error {
	if _, err := sdkTypes.AccAddressFromBech32(commonTransactionRequest.From); err != nil || len(commonTransactionRequest.From) == 0 {
		return fmt.Errorf("invalid from address: %s", commonTransactionRequest.From)
	}

	if !commonTransactionRequest.Simulate {
		if len(commonTransactionRequest.ChainID) == 0 {
			return fmt.Errorf("chain-id required but not specified")
		}
		if !commonTransactionRequest.Fees.IsZero() && !commonTransactionRequest.GasPrices.IsZero() {
			return fmt.Errorf("cannot provide both fees and gas prices")
		}
		if !commonTransactionRequest.Fees.IsValid() && !commonTransactionRequest.GasPrices.IsValid() {
			return fmt.Errorf("invalid fees or gas prices provided")
		}
	}

	return nil
}
func (commonTransactionRequest commonTransactionRequest) Sanitize() helpers.CommonTransactionRequest {
	return NewCommonTransactionRequest(commonTransactionRequest.From, commonTransactionRequest.Memo, commonTransactionRequest.ChainID, commonTransactionRequest.Gas, commonTransactionRequest.GasAdjustment, commonTransactionRequest.AccountNumber, commonTransactionRequest.TimeoutHeight, commonTransactionRequest.Sequence, commonTransactionRequest.Fees, commonTransactionRequest.GasPrices, commonTransactionRequest.Simulate)
}
func (commonTransactionRequest commonTransactionRequest) ValidateBasic(responseWriter http.ResponseWriter) bool {
	if err := commonTransactionRequest.Validate(); err != nil {
		rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
		return false
	}

	return true
}

func NewCommonTransactionRequest(from, memo, chainID, gas, gasAdjustment string, accountNumber, sequence, timeOutHeight uint64, fees sdkTypes.Coins, gasPrices sdkTypes.DecCoins, simulate bool) helpers.CommonTransactionRequest {
	return commonTransactionRequest{
		From:          strings.TrimSpace(from),
		Memo:          strings.TrimSpace(memo),
		ChainID:       strings.TrimSpace(chainID),
		Fees:          fees,
		GasPrices:     gasPrices,
		Gas:           strings.TrimSpace(gas),
		GasAdjustment: strings.TrimSpace(gasAdjustment),
		AccountNumber: accountNumber,
		Sequence:      sequence,
		TimeoutHeight: timeOutHeight,
		Simulate:      simulate,
	}
}
func PrototypeCommonTransactionRequest() helpers.CommonTransactionRequest {
	return &commonTransactionRequest{}
}

func NewCommonTransactionRequestFromContext(context client.Context) helpers.CommonTransactionRequest {
	return NewCommonTransactionRequest(context.From, "", context.ChainID, "", "", uint64(context.Height), 0, 0, sdkTypes.Coins{}, sdkTypes.DecCoins{}, context.Simulate)
}