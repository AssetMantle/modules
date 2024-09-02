// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	"encoding/json"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"io"
	"net/http"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	helpers.CommonTransactionRequest `json:"commonTransactionRequest"`
	FromID                           string `json:"fromID"`
	TakerSplit                       string `json:"takerSplit"`
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
	if msg, err := transactionRequest.MakeMsg(); err != nil {
		return err
	} else if err := msg.(helpers.Message).ValidateBasic(); err != nil {
		return err
	}

	return nil
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadCommonTransactionRequest(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.TakerSplit),
		cliCommand.ReadString(constants.OrderID),
	), nil
}
func (transactionRequest transactionRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.TransactionRequest, error) {
	body, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &transactionRequest); err != nil {
		return nil, err
	}

	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetCommonTransactionRequest() helpers.CommonTransactionRequest {
	return transactionRequest.CommonTransactionRequest
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetCommonTransactionRequest().GetFrom())
	if err != nil {
		return nil, err
	}

	takerSplit, ok := sdkTypes.NewIntFromString(transactionRequest.TakerSplit)
	if !ok {
		return nil, constants.IncorrectFormat.Wrapf("taker split %s is not a valid integer", transactionRequest.TakerSplit)
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
		takerSplit,
		orderID.(ids.OrderID),
	), nil
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(commonTransactionRequest helpers.CommonTransactionRequest, fromID string, takerSplit string, orderID string) helpers.TransactionRequest {
	return transactionRequest{
		CommonTransactionRequest: commonTransactionRequest,
		FromID:                   fromID,
		TakerSplit:               takerSplit,
		OrderID:                  orderID,
	}
}
