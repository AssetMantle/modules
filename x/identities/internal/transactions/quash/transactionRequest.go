// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/helpers/constants"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type transactionRequest struct {
	BaseReq    rest.BaseReq `json:"baseReq"`
	FromID     string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID "`
	IdentityID string       `json:"identityID" valid:"required~required field identityID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field identityID "`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Quash identity transaction
// @Description Quash identity transaction
// @Accept text/plain
// @Produce json
// @Tags Identities
// @Param body body  transactionRequest true "Request body for quash identity"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /identities/quash [post]
func (transactionRequest transactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context),
		cliCommand.ReadString(constants.FromID),
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

	fromID, err := baseIDs.ReadIdentityID(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	identityID, err := baseIDs.ReadIdentityID(transactionRequest.IdentityID)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		fromID,
		identityID,
	), nil
}
func (transactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, fromID string, identityID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:    baseReq,
		FromID:     fromID,
		IdentityID: identityID,
	}
}
