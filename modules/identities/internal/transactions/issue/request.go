// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

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
	"github.com/AssetMantle/modules/schema/lists/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type transactionRequest struct {
	BaseReq                 rest.BaseReq `json:"baseReq"`
	To                      string       `json:"to" valid:"required~required field to missing, matches(^[a-z0-9]+$)~invalid field to"`
	FromID                  string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID "`
	ClassificationID        string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^[A-Za-z0-9-_=.]+$)~invalid field classificationID"`
	ImmutableMetaProperties string       `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing, matches(^.*$)~invalid field immutableMetaProperties"`
	ImmutableProperties     string       `json:"immutableProperties" valid:"required~required field immutableProperties missing, matches(^.*$)~invalid field immutableProperties"`
	MutableMetaProperties   string       `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing, matches(^.*$)~invalid field mutableMetaProperties"`
	MutableProperties       string       `json:"mutableProperties" valid:"required~required field mutableProperties missing, matches(^.*$)~invalid field mutableProperties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Issue an identity
// @Description Issue identities with mutable immutable properties
// @Accept text/plain
// @Produce json
// @Tags Identities
// @Param body body  transactionRequest true "A transaction to issue an identity."
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /identities/issue [post]
func (transactionRequest transactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.To),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.ClassificationID),
		cliCommand.ReadString(flags.ImmutableMetaProperties),
		cliCommand.ReadString(flags.ImmutableProperties),
		cliCommand.ReadString(flags.MutableMetaProperties),
		cliCommand.ReadString(flags.MutableProperties),
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

	immutableMetaProperties, err := base.ReadMetaProperties(transactionRequest.ImmutableMetaProperties)
	if err != nil {
		return nil, err
	}

	immutableProperties, err := base.ReadProperties(transactionRequest.ImmutableProperties)
	if err != nil {
		return nil, err
	}

	mutableMetaProperties, err := base.ReadMetaProperties(transactionRequest.MutableMetaProperties)
	if err != nil {
		return nil, err
	}

	mutableProperties, err := base.ReadProperties(transactionRequest.MutableProperties)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		to,
		baseIDs.NewID(transactionRequest.FromID),
		baseIDs.NewID(transactionRequest.ClassificationID),
		immutableMetaProperties,
		immutableProperties,
		mutableMetaProperties,
		mutableProperties,
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, to string, fromID string, classificationID string, immutableMetaProperties string, immutableProperties string, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:                 baseReq,
		To:                      to,
		FromID:                  fromID,
		ClassificationID:        classificationID,
		ImmutableMetaProperties: immutableMetaProperties,
		ImmutableProperties:     immutableProperties,
		MutableMetaProperties:   mutableMetaProperties,
		MutableProperties:       mutableProperties,
	}
}
