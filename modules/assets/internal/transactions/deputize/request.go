/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type transactionRequest struct {
	BaseReq              rest.BaseReq `json:"baseReq"`
	FromID               string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ToID                 string       `json:"toID" valid:"required~required field toID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field toID"`
	ClassificationID     string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^[A-Za-z0-9-_=.]+$)~invalid field classificationID"`
	MaintainedProperties string       `json:"maintainedProperties" valid:"required~required field maintainedProperties missing, matches(^.*$)~invalid field maintainedProperties"`
	AddMaintainer        bool         `json:"addMaintainer" valid:"required~required field addMaintainer missing"`
	RemoveMaintainer     bool         `json:"removeMaintainer" valid:"required~required field removeMaintainer missing"`
	MutateMaintainer     bool         `json:"mutateMaintainer" valid:"required~required field mutateMaintainer missing"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Transaction Request godoc
// @Summary deputize asset transaction
// @Descrption deputize asset
// @Accept text/plain
// @Produce json
// @Tags Assets
// @Param body body  transactionRequest true "request body"
// @Success 200 {object} transactionResponse   "A successful response."
// @Failure default  {object}  transactionResponse "An unexpected error response."
// @Router /assets/deputize [post]
func (transactionRequest transactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.ToID),
		cliCommand.ReadString(flags.ClassificationID),
		cliCommand.ReadString(flags.MaintainedProperties),
		cliCommand.ReadBool(flags.AddMaintainer),
		cliCommand.ReadBool(flags.RemoveMaintainer),
		cliCommand.ReadBool(flags.MutateMaintainer),
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

	maintainedProperties, Error := base.ReadProperties(transactionRequest.MaintainedProperties)
	if Error != nil {
		return nil, Error
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.ToID),
		base.NewID(transactionRequest.ClassificationID),
		maintainedProperties,
		transactionRequest.AddMaintainer,
		transactionRequest.RemoveMaintainer,
		transactionRequest.MutateMaintainer,
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, toID string, classificationID string, maintainedProperties string, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:              baseReq,
		FromID:               fromID,
		ToID:                 toID,
		ClassificationID:     classificationID,
		MaintainedProperties: maintainedProperties,
		AddMaintainer:        addMaintainer,
		RemoveMaintainer:     removeMaintainer,
		MutateMaintainer:     mutateMaintainer,
	}
}
