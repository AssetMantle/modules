// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package get

import (
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"net/http"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	helpers.CommonTransactionRequest `json:"commonTransactionRequest"`
	FromID                           string `json:"fromID"`
	OrderID                          string `json:"orderID"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Take order transaction
// @Description Take order transaction
// @Accept text/plain
// @Produce json
// @Tags Orders
// @Param body  transactionRequest true "Request body to take order"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /orders/take [post]
func (transactionRequest transactionRequest) Validate() error {
	return helpers.ValidateTransactionRequest(transactionRequest)
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadCommonTransactionRequest(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.OrderID),
	), nil
}
func (transactionRequest transactionRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.TransactionRequest, error) {
	return helpers.TransactionRequestFromHTTPRequest(httpRequest, &transactionRequest)
}
func (transactionRequest transactionRequest) GetCommonTransactionRequest() helpers.CommonTransactionRequest {
	return transactionRequest.CommonTransactionRequest
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetCommonTransactionRequest().GetFrom())
	if err != nil {
		return nil, err
	}

	fromID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	orderID, err := baseIDs.PrototypeOrderID().FromString(transactionRequest.OrderID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		orderID.(ids.OrderID),
	), nil
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(commonTransactionRequest helpers.CommonTransactionRequest, fromID string, orderID string) helpers.TransactionRequest {
	return transactionRequest{
		CommonTransactionRequest: commonTransactionRequest,
		FromID:                   fromID,
		OrderID:                  orderID,
	}
}
