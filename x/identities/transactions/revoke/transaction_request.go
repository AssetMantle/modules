// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

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
	ClassificationID                 string `json:"classificationID"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Revoke identity transaction
// @Description Revoke identity transaction
// @Accept text/plain
// @Produce json
// @Tags Identities
// @Param body  transactionRequest true "Request body for revoke identity"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /identities/revoke [post]
func (transactionRequest transactionRequest) Validate() error {
	return helpers.ValidateTransactionRequest(transactionRequest)
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadCommonTransactionRequest(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.ToIdentityID),
		cliCommand.ReadString(constants.ClassificationID),
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

	toID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.ToID)
	if err != nil {
		return nil, err
	}
	classificationID, err := baseIDs.PrototypeClassificationID().FromString(transactionRequest.ClassificationID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		toID.(ids.IdentityID),
		classificationID.(ids.ClassificationID),
	), nil
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(commonTransactionRequest helpers.CommonTransactionRequest, fromID string, toID string, classificationID string) helpers.TransactionRequest {
	return transactionRequest{
		CommonTransactionRequest: commonTransactionRequest,
		FromID:                   fromID,
		ToID:                     toID,
		ClassificationID:         classificationID,
	}
}
