// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"encoding/json"

	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
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
	BaseReq      rest.BaseReq `json:"baseReq"`
	FromID       string       `json:"fromID" valid:"required~required field fromID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field fromID"`
	MakerAssetID string       `json:"makerAssetID" valid:"required~required field makerAssetID missing, matches(^(COI|AI)\|((?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4}|[A-Za-z0-9]{32}))$)~invalid field makerAssetID"`
	TakerCoinID  string       `json:"takerCoinID" valid:"required~required field takerCoinID missing, matches(^(COI|AI)\|((?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4}|[A-Za-z0-9]{32}))$)~invalid field takerCoinID"`
	MakerSplit   string       `json:"makerSplit" valid:"required~required field makerSplit missing, matches(^[0-9.]+$)~invalid field makerSplit"`
	TakerSplit   string       `json:"takerSplit" valid:"required~required field takerSplit missing, matches(^[0-9.]+$)~invalid field takerSplit"`
	ExpiryHeight int64        `json:"expiryHeight" valid:"required~required field expiryHeight missing, matches(^[0-9]+$)~invalid field expiryHeight"`
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
		cliCommand.ReadString(constants.MakerAssetID),
		cliCommand.ReadString(constants.TakerCoinID),
		cliCommand.ReadString(constants.MakerSplit),
		cliCommand.ReadString(constants.TakerSplit),
		cliCommand.ReadInt64(constants.ExpiryHeight),
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

	makerAssetID, err := baseIDs.PrototypeAssetID().FromString(transactionRequest.MakerAssetID)
	if err != nil {
		return nil, err
	}

	takerCoinID, err := baseIDs.PrototypeCoinID().FromString(transactionRequest.TakerCoinID)
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

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		makerAssetID.(ids.AssetID),
		takerCoinID.(ids.CoinID),
		makerSplit,
		takerSplit,
		baseTypes.NewHeight(transactionRequest.ExpiryHeight),
	), nil
}
func (transactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, makerAssetID string, takerCoinID string, makerSplit, takerSplit string, expiryHeight int64) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:      baseReq,
		FromID:       fromID,
		MakerAssetID: makerAssetID,
		TakerCoinID:  takerCoinID,
		MakerSplit:   makerSplit,
		TakerSplit:   takerSplit,
		ExpiryHeight: expiryHeight,
	}
}
