// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type transactionRequest struct {
	BaseReq    rest.BaseReq `json:"baseReq"`
	To         string       `json:"to" valid:"required~required field to missing, matches(^[a-z0-9]+$)~invalid field to"`
	IdentityID string       `json:"identityID" valid:"required~required field identityID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field identityID"`
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
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.To),
		cliCommand.ReadString(constants.IdentityID),
	), nil
}
func (transactionRequest transactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if err := json.Unmarshal(rawMessage, &transactionRequest); err != nil {
		return nil, err
	}

	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if err != nil {
		return nil, err
	}

	to, err := sdkTypes.AccAddressFromBech32(transactionRequest.To)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		to,
		baseIDs.NewID(transactionRequest.IdentityID),
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, to string, identityID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:    baseReq,
		To:         to,
		IdentityID: identityID,
	}
}
