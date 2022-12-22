// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

//type transactionRequest struct {
//	BaseReq                 rest.BaseReq `json:"baseReq"`
//	FromID                  string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
//	ToID                    string       `json:"toID" valid:"required~required field toID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field toID"`
//	ClassificationID        string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^[A-Za-z0-9-_=.]+$)~invalid field classificationID"`
//	ImmutableMetaProperties string       `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing, matches(^.*$)~invalid field immutableMetaProperties"`
//	ImmutableProperties     string       `json:"immutableProperties" valid:"required~required field immutableProperties missing, matches(^.*$)~invalid field immutableProperties"`
//	MutableMetaProperties   string       `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing, matches(^.*$)~invalid field mutableMetaProperties"`
//	MutableProperties       string       `json:"mutableProperties" valid:"required~required field mutableProperties missing, matches(^.*$)~invalid field mutableProperties"`
//}

var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

// Validate godoc
// @Summary Mint an asset transaction
// @Description Mint asset with mutable immutable properties
// @Accept text/plain
// @Produce json
// @Tags Assets
// @Param body  transactionRequest true "A transaction to mint the asset."
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /assets/mint [post]
func (transactionRequest *TransactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	if err != nil {
		return err
	}
	inputValidator := base.NewInputValidator(constants.PropertyExpression)
	if !inputValidator.IsValid(transactionRequest.ImmutableProperties, transactionRequest.ImmutableMetaProperties, transactionRequest.MutableProperties, transactionRequest.MutableMetaProperties) {
		return errorConstants.IncorrectFormat
	}
	return nil
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
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.ImmutableMetaProperties),
		cliCommand.ReadString(constants.ImmutableProperties),
		cliCommand.ReadString(constants.MutableMetaProperties),
		cliCommand.ReadString(constants.MutableProperties),
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

	immutableMetaProperties, err := utilities.ReadMetaPropertyList(transactionRequest.ImmutableMetaProperties)
	if err != nil {
		return nil, err
	}

	immutableProperties, err := utilities.ReadMetaPropertyList(transactionRequest.ImmutableProperties)
	if err != nil {
		return nil, err
	}
	immutableProperties = immutableProperties.ScrubData()

	mutableMetaProperties, err := utilities.ReadMetaPropertyList(transactionRequest.MutableMetaProperties)
	if err != nil {
		return nil, err
	}

	mutableProperties, err := utilities.ReadMetaPropertyList(transactionRequest.MutableProperties)
	if err != nil {
		return nil, err
	}
	mutableProperties = mutableProperties.ScrubData()

	fromID, err := baseIDs.ReadIdentityID(transactionRequest.FromId)
	if err != nil {
		return nil, err
	}

	toID, err := baseIDs.ReadIdentityID(transactionRequest.ToId)
	if err != nil {
		return nil, err
	}

	classificationID, err := baseIDs.ReadClassificationID(transactionRequest.ClassificationId)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		fromID,
		toID,
		classificationID,
		immutableMetaProperties,
		immutableProperties,
		mutableMetaProperties,
		mutableProperties,
	), nil
}
func (*TransactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, &TransactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return &TransactionRequest{}
}

func newTransactionRequest(from string, fromID string, toID string, classificationID string, immutableMetaProperties string, immutableProperties string, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return &TransactionRequest{
		From:                    from,
		FromId:                  fromID,
		ToId:                    toID,
		ClassificationId:        classificationID,
		ImmutableMetaProperties: immutableMetaProperties,
		ImmutableProperties:     immutableProperties,
		MutableMetaProperties:   mutableMetaProperties,
		MutableProperties:       mutableProperties,
	}
}
