// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type transactionRequest struct {
	BaseReq              rest.BaseReq `json:"baseReq"`
	FromID               string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ToID                 string       `json:"toID" valid:"required~required field toID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field toID"`
	ClassificationID     string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^[A-Za-z0-9-_=.]+$)~invalid field classificationID"`
	MaintainedProperties string       `json:"maintainedProperties" valid:"required~required field maintainedProperties missing, matches(^.*$)~invalid field maintainedProperties"`
	CanMintAsset         bool         `json:"canMintAsset"`
	CanBurnAsset         bool         `json:"canBurnAsset"`
	CanRenumerateAsset   bool         `json:"canRenumerateAsset"`
	CanAddMaintainer     bool         `json:"canAddMaintainer"`
	CanRemoveMaintainer  bool         `json:"canRemoveMaintainer"`
	CanMutateMaintainer  bool         `json:"canMutateMaintainer"`
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
	if err != nil {
		return err
	}

	return nil
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context),
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.ToID),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.MaintainedProperties),
		cliCommand.ReadBool(constants.CanMintAsset),
		cliCommand.ReadBool(constants.CanBurnAsset),
		cliCommand.ReadBool(constants.CanRenumerateAsset),
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

	maintainedProperties, err := utilities.ReadMetaPropertyList(transactionRequest.MaintainedProperties)
	if err != nil {
		return nil, err
	}

	fromID, err := baseIDs.ReadIdentityID(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	toID, err := baseIDs.ReadIdentityID(transactionRequest.ToID)
	if err != nil {
		return nil, err
	}

	classificationID, err := baseIDs.ReadClassificationID(transactionRequest.ClassificationID)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		fromID,
		toID,
		classificationID,
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
func newTransactionRequest(baseReq rest.BaseReq, fromID string, toID string, classificationID string, maintainedProperties string, canMintAsset bool, canBurnAsset bool, canRenumerateAsset bool, canAddMaintainer bool, canRemoveMaintainer bool, canMutateMaintainer bool) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:              baseReq,
		FromID:               fromID,
		ToID:                 toID,
		ClassificationID:     classificationID,
		MaintainedProperties: maintainedProperties,
		CanMintAsset:         canMintAsset,
		CanBurnAsset:         canBurnAsset,
		CanRenumerateAsset:   canRenumerateAsset,
		CanAddMaintainer:     canAddMaintainer,
		CanRemoveMaintainer:  canRemoveMaintainer,
		CanMutateMaintainer:  canMutateMaintainer,
	}
}
