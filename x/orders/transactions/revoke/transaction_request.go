// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"encoding/json"

	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	FromID           string       `json:"fromID" valid:"required~required field fromID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field fromID"`
	ToID             string       `json:"toID" valid:"required~required field toID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field toID"`
	ClassificationID string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field classificationID"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Revoke order transaction
// @Description Revoke order transaction
// @Accept text/plain
// @Produce json
// @Tags Orders
// @Param body transactionRequest true "Request body to revoke order"
// @Success 200 {object} transactionResponse "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /orders/revoke [post]
func (transactionRequest transactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.ToIdentityID),
		cliCommand.ReadString(constants.ClassificationID),
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
func (transactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, toID string, classificationID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		FromID:           fromID,
		ToID:             toID,
		ClassificationID: classificationID,
	}
}
