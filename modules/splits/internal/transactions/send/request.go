// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

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

type transactionRequest struct {
	BaseReq   rest.BaseReq `json:"baseReq"`
	FromID    string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ToID      string       `json:"toID" valid:"required~required field toID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field toID"`
	OwnableID string       `json:"ownableID" valid:"required~required field ownableID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field ownableID"`
	Value     string       `json:"value" valid:"required~required field value missing, matches(^[0-9.]+$)~invalid field value"`
}

var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

// Validate godoc
// @Summary Send split transaction
// @Description Send split transaction
// @Accept text/plain
// @Produce json
// @Tags Splits
// @Param body body  transactionRequest true "Request body to send split"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /splits/send [post]
func (transactionRequest *TransactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest *TransactionRequest) RegisterInterface(registry types.InterfaceRegistry) {
	//TODO implement me
	panic("implement me")
}
func (transactionRequest *TransactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context).From,
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.ToID),
		cliCommand.ReadString(constants.OwnableID),
		cliCommand.ReadString(constants.Value),
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
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.From)
	if err != nil {
		return nil, err
	}

	value, err := sdkTypes.NewDecFromStr(transactionRequest.Value)
	if err != nil {
		return nil, err
	}

	fromID, err := baseIDs.ReadIdentityID(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	toID, err := baseIDs.ReadIdentityID(transactionRequest.ToID)
	if err != nil {
		return nil, err
	}

	ownableID, err := baseIDs.ReadOwnableID(transactionRequest.OwnableID)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		fromID,
		toID,
		ownableID,
		value,
	), nil
}
func (*TransactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return &TransactionRequest{}
}
func newTransactionRequest(from string, fromID string, toID string, ownableID string, value string) helpers.TransactionRequest {
	return &TransactionRequest{
		From:      from,
		FromID:    fromID,
		ToID:      toID,
		OwnableID: ownableID,
		Value:     value,
	}
}
