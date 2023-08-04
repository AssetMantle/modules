// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

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
	BaseReq                 rest.BaseReq `json:"baseReq"`
	FromID                  string       `json:"fromID" valid:"required~required field fromID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field fromID"`
	ClassificationID        string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field classificationID"`
	TakerID                 string       `json:"takerID" valid:"required~required field takerID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field takerID"`
	MakerOwnableID          string       `json:"makerOwnableID" valid:"required~required field makerOwnableID missing, matches(^(COI|AI)\|((?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4}|[A-Za-z0-9]{32}))$)~invalid field makerOwnableID"`
	TakerOwnableID          string       `json:"takerOwnableID" valid:"required~required field takerOwnableID missing, matches(^(COI|AI)\|((?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4}|[A-Za-z0-9]{32}))$)~invalid field takerOwnableID"`
	ExpiresIn               int64        `json:"expiresIn" valid:"required~required field expiresIn missing, matches(^[0-9]+$)~invalid field expiresIn"`
	MakerOwnableSplit       string       `json:"makerOwnableSplit" valid:"required~required field makerOwnableSplit missing, matches(^[0-9.]+$)~invalid field makerOwnableSplit"`
	TakerOwnableSplit       string       `json:"takerOwnableSplit" valid:"required~required field takerOwnableSplit missing, matches(^[0-9.]+$)~invalid field takerOwnableSplit"`
	ImmutableMetaProperties string       `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing, matches(^.*$)~invalid field immutableMetaProperties"`
	ImmutableProperties     string       `json:"immutableProperties" valid:"required~required field immutableProperties missing, matches(^.*$)~invalid field immutableProperties"`
	MutableMetaProperties   string       `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing, matches(^.*$)~invalid field mutableMetaProperties"`
	MutableProperties       string       `json:"mutableProperties" valid:"required~required field mutableProperties missing, matches(^.*$)~invalid field mutableProperties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Make order transaction
// @Description Make order transaction
// @Accept text/plain
// @Produce json
// @Tags Orders
// @Param body  transactionRequest true "Request body to make order"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /orders/make [post]
func (transactionRequest transactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.TakerID),
		cliCommand.ReadString(constants.MakerOwnableID),
		cliCommand.ReadString(constants.TakerOwnableID),
		cliCommand.ReadInt64(constants.ExpiresIn),
		cliCommand.ReadString(constants.MakerSplit),
		cliCommand.ReadString(constants.TakerSplit),
		cliCommand.ReadString(constants.ImmutableMetaProperties),
		cliCommand.ReadString(constants.ImmutableProperties),
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

	makerOwnableSplit, ok := sdkTypes.NewIntFromString(transactionRequest.MakerOwnableSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("maker ownable split %s is not a valid integer", transactionRequest.MakerOwnableSplit)
	}

	takerOwnableSplit, ok := sdkTypes.NewIntFromString(transactionRequest.TakerOwnableSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("taker ownable split %s is not a valid integer", transactionRequest.TakerOwnableSplit)
	}

	immutableMetaProperties, err := base.PrototypePropertyList().FromMetaPropertiesString(transactionRequest.ImmutableMetaProperties)
	if err != nil {
		return nil, err
	}

	immutableProperties, err := base.PrototypePropertyList().FromMetaPropertiesString(transactionRequest.ImmutableProperties)
	if err != nil {
		return nil, err
	}
	immutableProperties = immutableProperties.ScrubData()

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

	classificationID, err := baseIDs.PrototypeClassificationID().FromString(transactionRequest.ClassificationID)
	if err != nil {
		return nil, err
	}

	takerID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.TakerID)
	if err != nil {
		return nil, err
	}

	makerOwnableID, err := baseIDs.PrototypeAnyOwnableID().FromString(transactionRequest.MakerOwnableID)
	if err != nil {
		return nil, err
	}

	takerOwnableID, err := baseIDs.PrototypeAnyOwnableID().FromString(transactionRequest.TakerOwnableID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		classificationID.(ids.ClassificationID),
		takerID.(ids.IdentityID),
		makerOwnableID.(ids.OwnableID).ToAnyOwnableID(),
		takerOwnableID.(ids.OwnableID).ToAnyOwnableID(),
		baseTypes.NewHeight(transactionRequest.ExpiresIn),
		makerOwnableSplit,
		takerOwnableSplit,
		immutableMetaProperties,
		immutableProperties,
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

func newTransactionRequest(baseReq rest.BaseReq, fromID string, classificationID string, takerID string, makerOwnableID string, takerOwnableID string, expiresIn int64, makerOwnableSplit, takerOwnableSplit string, immutableMetaProperties string, immutableProperties string, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:                 baseReq,
		FromID:                  fromID,
		ClassificationID:        classificationID,
		TakerID:                 takerID,
		MakerOwnableID:          makerOwnableID,
		TakerOwnableID:          takerOwnableID,
		ExpiresIn:               expiresIn,
		MakerOwnableSplit:       makerOwnableSplit,
		TakerOwnableSplit:       takerOwnableSplit,
		ImmutableMetaProperties: immutableMetaProperties,
		ImmutableProperties:     immutableProperties,
		MutableMetaProperties:   mutableMetaProperties,
		MutableProperties:       mutableProperties,
	}
}
