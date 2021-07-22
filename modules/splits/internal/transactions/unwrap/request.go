/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package unwrap

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type transactionRequest struct {
	BaseReq   rest.BaseReq `json:"baseReq"`
	FromID    string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	OwnableID string       `json:"ownableID" valid:"required~required field ownableID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field ownableID"`
	Value     string       `json:"value" valid:"required~required field value missing, matches(^[0-9]+$)~invalid field value"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Transaction Request godoc
// @Summary unwrap split transaction
// @Descrption unwrap split transaction
// @Accept text/plain
// @Produce json
// @Tags Splits
// @Param body body  transactionRequest true "request body"
// @Success 200 {object} transactionResponse   "A successful response."
// @Failure default {object} transactionResponse "An unexpected error response."
// @Router /splits/unwrap [post]
func (transactionRequest transactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.OwnableID),
		cliCommand.ReadString(flags.Value),
	), nil
}
func (transactionRequest transactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if Error := json.Unmarshal(rawMessage, &transactionRequest); Error != nil {
		return nil, Error
	}

	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		return nil, Error
	}

	value, ok := sdkTypes.NewIntFromString(transactionRequest.Value)
	if !ok {
		return nil, xprtErrors.InvalidRequest
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.OwnableID),
		value,
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, fromID string, ownableID string, value string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:   baseReq,
		FromID:    fromID,
		OwnableID: ownableID,
		Value:     value,
	}
}
