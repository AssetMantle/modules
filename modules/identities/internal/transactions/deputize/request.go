// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/lists/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type transactionRequest struct {
	BaseReq              rest.BaseReq `json:"baseReq"`
	FromID               string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ToID                 string       `json:"toID" valid:"required~required field toID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field toID"`
	ClassificationID     string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^[A-Za-z0-9-_=.]+$)~invalid field classificationID"`
	MaintainedProperties string       `json:"maintainedProperties" valid:"required~required field maintainedProperties missing, matches(^.*$)~invalid field maintainedProperties"`
	AddMaintainer        bool         `json:"addMaintainer"`
	RemoveMaintainer     bool         `json:"removeMaintainer"`
	MutateMaintainer     bool         `json:"mutateMaintainer"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate  godoc
// @Summary Deputize an identity
// @Description A transaction to deputize a maintainer for an identity classification.
// @Accept text/plain
// @Produce json
// @Tags Identities
// @Param body  transactionRequest true "Request body to deputize identity"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /identities/deputize [post]
func (transactionRequest transactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
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

	maintainedProperties, err := baseTypes.ReadProperties(transactionRequest.MaintainedProperties)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		baseIDs.NewID(transactionRequest.FromID),
		baseIDs.NewID(transactionRequest.ToID),
		baseIDs.NewID(transactionRequest.ClassificationID),
		maintainedProperties,
		transactionRequest.AddMaintainer,
		transactionRequest.RemoveMaintainer,
		transactionRequest.MutateMaintainer,
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, transactionRequest{})
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
