// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

//type transactionRequest struct {
//	BaseReq    rest.BaseReq `json:"baseReq"`
//	FromID     string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID "`
//	IdentityID string       `json:"identityID" valid:"required~required field identityID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field identityID "`
//}

var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

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
func (transactionRequest *TransactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (m *TransactionRequest) RegisterInterface(registry types.InterfaceRegistry) {
	//TODO implement me
	panic("implement me")
}
func (transactionRequest *TransactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context).From,
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.IdentityID),
	), nil
}
func (transactionRequest *TransactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if err := json.Unmarshal(rawMessage, &transactionRequest); err != nil {
		return nil, err
	}

	return transactionRequest, nil
}
func (transactionRequest *TransactionRequest) GetBaseReq() rest.BaseReq {
	panic("Implement me")
}
func (transactionRequest *TransactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if err != nil {
		return nil, err
	}

	fromID, err := baseIDs.ReadIdentityID(transactionRequest.FromId)
	if err != nil {
		return nil, err
	}

	identityID, err := baseIDs.ReadIdentityID(transactionRequest.IdentityId)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		fromID,
		identityID,
	), nil
}
func (*TransactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, &TransactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return &TransactionRequest{}
}
func newTransactionRequest(from string, fromID string, identityID string) helpers.TransactionRequest {
	return &TransactionRequest{
		From:       from,
		FromId:     fromID,
		IdentityId: identityID,
	}
}
