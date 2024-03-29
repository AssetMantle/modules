// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"encoding/json"

	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	BaseReq               rest.BaseReq `json:"baseReq"`
	FromID                string       `json:"fromID" valid:"required~required field fromID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field fromID"`
	OrderID               string       `json:"orderID" valid:"required~required field orderID missing, matches(^[A-Za-z0-9-_|=.*]+$)~invalid field orderID"`
	TakerSplit            string       `json:"takerSplit" valid:"required~required field takerSplit missing, matches(^[0-9.]+$)~invalid field takerSplit"`
	MakerSplit            string       `json:"makerSplit" valid:"required~required field makerSplit missing, matches(^[0-9.]+$)~invalid field makerSplit"`
	ExpiresIn             int64        `json:"expiresIn" valid:"required~required field expiresIn missing, matches(^[0-9]+$)~invalid field expiresIn"`
	MutableMetaProperties string       `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing, matches(^.*$)~invalid field mutableMetaProperties"`
	MutableProperties     string       `json:"mutableProperties" valid:"required~required field mutableProperties missing, matches(^.*$)~invalid field mutableProperties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Modify order transaction
// @Description Modify order transaction
// @Accept text/plain
// @Produce json
// @Tags Orders
// @Param body transactionRequest true "Request body to modify order transaction"
// @Success 200 {object} transactionResponse "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /orders/modify [post]
func (transactionRequest transactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.OrderID),
		cliCommand.ReadString(constants.TakerSplit),
		cliCommand.ReadString(constants.MakerSplit),
		cliCommand.ReadInt64(constants.ExpiresIn),
		cliCommand.ReadString(constants.MutableMetaProperties),
		cliCommand.ReadString(constants.MutableProperties),
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

	makerSplit, ok := sdkTypes.NewIntFromString(transactionRequest.MakerSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("maker split %s is not a valid integer", transactionRequest.MakerSplit)
	}

	takerSplit, ok := sdkTypes.NewIntFromString(transactionRequest.TakerSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("taker split %s is not a valid integer", transactionRequest.TakerSplit)
	}

	mutableMetaProperties, err := base.PrototypePropertyList().FromMetaPropertiesString(transactionRequest.MutableMetaProperties)
	if err != nil {
		return nil, err
	}

	mutableProperties, err := base.PrototypePropertyList().FromMetaPropertiesString(transactionRequest.MutableProperties)
	if err != nil {
		return nil, err
	}
	mutableProperties = mutableProperties.ScrubData()

	fromID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	orderID, err := baseIDs.PrototypeOrderID().FromString(transactionRequest.OrderID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		orderID.(ids.OrderID),
		takerSplit,
		makerSplit,
		baseTypes.NewHeight(transactionRequest.ExpiresIn),
		mutableMetaProperties,
		mutableProperties,
	), nil
}
func (transactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, orderID string, takerSplit string, makerSplit string, expiresIn int64, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:               baseReq,
		FromID:                fromID,
		OrderID:               orderID,
		TakerSplit:            takerSplit,
		MakerSplit:            makerSplit,
		ExpiresIn:             expiresIn,
		MutableMetaProperties: mutableMetaProperties,
		MutableProperties:     mutableProperties,
	}
}
