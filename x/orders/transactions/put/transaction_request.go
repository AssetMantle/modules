// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"encoding/json"
	"github.com/AssetMantle/modules/utilities/rest"

	codecUtilities "github.com/AssetMantle/schema/codec/utilities"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	BaseReq      rest.BaseReq `json:"baseReq"`
	FromID       string       `json:"fromID"`
	MakerAssetID string       `json:"makerAssetID"`
	TakerAssetID string       `json:"takerAssetID"`
	MakerSplit   string       `json:"makerSplit"`
	TakerSplit   string       `json:"takerSplit"`
	ExpiryHeight int64        `json:"expiryHeight"`
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
	if msg, err := transactionRequest.MakeMsg(); err != nil {
		return err
	} else if err := msg.(helpers.Message).ValidateBasic(); err != nil {
		return err
	}

	return nil
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.MakerAssetID),
		cliCommand.ReadString(constants.TakerAssetID),
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

	takerAssetID, err := baseIDs.PrototypeAssetID().FromString(transactionRequest.TakerAssetID)
	if err != nil {
		return nil, err
	}

	makerSplit, ok := sdkTypes.NewIntFromString(transactionRequest.MakerSplit)
	if !ok {
		return nil, constants.IncorrectFormat.Wrapf("maker split %s is not a valid integer", transactionRequest.MakerSplit)
	}

	takerSplit, ok := sdkTypes.NewIntFromString(transactionRequest.TakerSplit)
	if !ok {
		return nil, constants.IncorrectFormat.Wrapf("taker split %s is not a valid integer", transactionRequest.TakerSplit)
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		makerAssetID.(ids.AssetID),
		takerAssetID.(ids.AssetID),
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

func newTransactionRequest(baseReq rest.BaseReq, fromID string, makerAssetID string, takerAssetID string, makerSplit, takerSplit string, expiryHeight int64) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:      baseReq,
		FromID:       fromID,
		MakerAssetID: makerAssetID,
		TakerAssetID: takerAssetID,
		MakerSplit:   makerSplit,
		TakerSplit:   takerSplit,
		ExpiryHeight: expiryHeight,
	}
}
