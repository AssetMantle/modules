// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"encoding/json"
	"github.com/AssetMantle/modules/utilities/rest"

	codecUtilities "github.com/AssetMantle/schema/codec/utilities"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	BaseReq               rest.BaseReq `json:"baseReq"`
	FromID                string       `json:"fromID"`
	AssetID               string       `json:"assetID"`
	MutableMetaProperties string       `json:"mutableMetaProperties"`
	MutableProperties     string       `json:"mutableProperties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Mutate an asset transaction
// @Description Mutate asset
// @Accept text/plain
// @Produce json
// @Tags Assets
// @Param body  transactionRequest true "A transaction to mutate an asset."
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /assets/mutate [post]
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
		cliCommand.ReadString(constants.AssetID),
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

	mutableMetaProperties, err := base.NewPropertyList().FromMetaPropertiesString(transactionRequest.MutableMetaProperties)
	if err != nil {
		return nil, err
	}

	mutableProperties, err := base.NewPropertyList().FromMetaPropertiesString(transactionRequest.MutableProperties)
	if err != nil {
		return nil, err
	}
	mutableProperties = mutableProperties.ScrubData()

	fromID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.FromID)
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
		assetID.(ids.AssetID),
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

func newTransactionRequest(baseReq rest.BaseReq, fromID string, assetID string, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:               baseReq,
		FromID:                fromID,
		AssetID:               assetID,
		MutableMetaProperties: mutableMetaProperties,
		MutableProperties:     mutableProperties,
	}
}
