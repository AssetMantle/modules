// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"encoding/json"

	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	BaseReq              rest.BaseReq `json:"baseReq"`
	FromID               string       `json:"fromID"`
	ToID                 string       `json:"toID"`
	ClassificationID     string       `json:"classificationID"`
	MaintainedProperties string       `json:"maintainedProperties"`
	CanMintAsset         bool         `json:"canMintAsset"`
	CanRenumerateAsset   bool         `json:"canRenumerateAsset"`
	CanBurnAsset         bool         `json:"canBurnAsset"`
	CanAddMaintainer     bool         `json:"canAddMaintainer"`
	CanRemoveMaintainer  bool         `json:"canRemoveMaintainer"`
	CanMutateMaintainer  bool         `json:"canMutateMaintainer"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Deputize an asset transaction
// @Description Deputize asset
// @Accept text/plain
// @Produce json
// @Tags Assets
// @Param body body  transactionRequest true "request body"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /assets/deputize [post]
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
		cliCommand.ReadString(constants.ToIdentityID),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.MaintainedProperties),
		cliCommand.ReadBool(constants.CanMintAsset),
		cliCommand.ReadBool(constants.CanRenumerateAsset),
		cliCommand.ReadBool(constants.CanBurnAsset),
		cliCommand.ReadBool(constants.CanAddMaintainer),
		cliCommand.ReadBool(constants.CanRemoveMaintainer),
		cliCommand.ReadBool(constants.CanMutateMaintainer),
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

	maintainedProperties, err := base.NewPropertyList().FromMetaPropertiesString(transactionRequest.MaintainedProperties)
	if err != nil {
		return nil, err
	}

	fromID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	toID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.ToID)
	if err != nil {
		return nil, err
	}

	classificationID, err := baseIDs.PrototypeClassificationID().FromString(transactionRequest.ClassificationID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		toID.(ids.IdentityID),
		classificationID.(ids.ClassificationID),
		maintainedProperties,
		transactionRequest.CanMintAsset,
		transactionRequest.CanBurnAsset,
		transactionRequest.CanRenumerateAsset,
		transactionRequest.CanAddMaintainer,
		transactionRequest.CanRemoveMaintainer,
		transactionRequest.CanMutateMaintainer,
	), nil
}
func (transactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, toID string, classificationID string, maintainedProperties string, canMintAsset bool, canRenumerateAsset bool, canBurnAsset bool, canAddMaintainer bool, canRemoveMaintainer bool, canMutateMaintainer bool) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:              baseReq,
		FromID:               fromID,
		ToID:                 toID,
		ClassificationID:     classificationID,
		MaintainedProperties: maintainedProperties,
		CanMintAsset:         canMintAsset,
		CanRenumerateAsset:   canRenumerateAsset,
		CanBurnAsset:         canBurnAsset,
		CanAddMaintainer:     canAddMaintainer,
		CanRemoveMaintainer:  canRemoveMaintainer,
		CanMutateMaintainer:  canMutateMaintainer,
	}
}
