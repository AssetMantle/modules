// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

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
	To                               string `json:"to"`
	IdentityID                       string `json:"identityID"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Provision an identity
// @Description Provision identities
// @Accept text/plain
// @Produce json
// @Tags Identities
// @Param body body  transactionRequest true "Request body for provision identity"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /identities/provision [post]
func (transactionRequest transactionRequest) Validate() error {
	return helpers.ValidateTransactionRequest(transactionRequest)
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadCommonTransactionRequest(context),
		cliCommand.ReadString(constants.To),
		cliCommand.ReadString(constants.IdentityID),
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

	to, err := sdkTypes.AccAddressFromBech32(transactionRequest.To)
	if err != nil {
		return nil, err
	}

	identityID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.IdentityID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		to,
		identityID.(ids.IdentityID),
	), nil
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(commonTransactionRequest helpers.CommonTransactionRequest, to string, identityID string) helpers.TransactionRequest {
	return transactionRequest{
		CommonTransactionRequest: commonTransactionRequest,
		To:                       to,
		IdentityID:               identityID,
	}
}
