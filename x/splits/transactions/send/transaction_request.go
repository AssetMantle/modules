// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"encoding/json"

	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	BaseReq rest.BaseReq `json:"baseReq"`
	FromID  string       `json:"fromID" valid:"required~required field fromID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field fromID"`
	ToID    string       `json:"toID" valid:"required~required field toID missing, matches(^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$)~invalid field toID"`
	AssetID string       `json:"assetID" valid:"required~required field assetID missing, matches(^(COI|AI)\|((?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4}|[A-Za-z0-9]{32}))$)~invalid field assetID"`
	Value   string       `json:"value" valid:"required~required field value missing, matches(^[0-9.]+$)~invalid field value"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

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
func (transactionRequest transactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.ToIdentityID),
		cliCommand.ReadString(constants.AssetID),
		cliCommand.ReadString(constants.Value),
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

	value, ok := sdkTypes.NewIntFromString(transactionRequest.Value)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("send value %s is not a valid integer", transactionRequest.Value)
	}

	fromID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	toID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.ToID)
	if err != nil {
		return nil, err
	}

	assetID, err := baseIDs.PrototypeAssetID().FromString(transactionRequest.AssetID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		toID.(ids.IdentityID),
		assetID.(ids.AssetID),
		value,
	), nil
}
func (transactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, fromID string, toID string, assetID string, value string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq: baseReq,
		FromID:  fromID,
		ToID:    toID,
		AssetID: assetID,
		Value:   value,
	}
}
