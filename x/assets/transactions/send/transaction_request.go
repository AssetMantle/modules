// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

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
	ToID                             string `json:"toID"`
	AssetID                          string `json:"assetID"`
	Value                            string `json:"value"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Send asset transaction
// @Description Send asset transaction
// @Accept text/plain
// @Produce json
// @Tags Assets
// @Param body body  transactionRequest true "Request body to send asset"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /assets/send [post]
func (transactionRequest transactionRequest) Validate() error {
	return helpers.ValidateTransactionRequest(transactionRequest)
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadCommonTransactionRequest(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.ToIdentityID),
		cliCommand.ReadString(constants.AssetID),
		cliCommand.ReadString(constants.Value),
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

	value, ok := sdkTypes.NewIntFromString(transactionRequest.Value)
	if !ok {
		return nil, constants.IncorrectFormat.Wrapf("send value %s is not a valid integer", transactionRequest.Value)
	}

	fromID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	toID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.ToID)
	if err != nil {
		return nil, err
	}

	assetID, err := baseIDs.PrototypeAssetID().FromString(transactionRequest.AssetID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		toID.(ids.IdentityID),
		assetID.(ids.AssetID),
		value,
	), nil
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(commonTransactionRequest helpers.CommonTransactionRequest, fromID string, toID string, assetID string, value string) helpers.TransactionRequest {
	return transactionRequest{
		CommonTransactionRequest: commonTransactionRequest,
		FromID:                   fromID,
		ToID:                     toID,
		AssetID:                  assetID,
		Value:                    value,
	}
}
