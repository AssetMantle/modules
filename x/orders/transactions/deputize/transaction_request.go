// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

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
	rest.CommonTransactionRequest `json:"commonTransactionRequest"`
	FromID                        string `json:"fromID"`
	ToID                          string `json:"toID"`
	ClassificationID              string `json:"classificationID"`
	MaintainedProperties          string `json:"maintainedProperties"`
	CanMakeOrder                  bool   `json:"canMakeOrder"`
	CanCancelOrder                bool   `json:"canCancelOrder"`
	CanAddMaintainer              bool   `json:"canAddMaintainer"`
	CanRemoveMaintainer           bool   `json:"canRemoveMaintainer"`
	CanMutateMaintainer           bool   `json:"canMutateMaintainer"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Deputize order transaction
// @Description Deputize order transaction
// @Accept text/plain
// @Produce json
// @Tags Orders
// @Param body body  transactionRequest true "Request body to deputize an order"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /orders/deputize [post]
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
		cliCommand.ReadCommonTransactionRequest(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.ToIdentityID),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.MaintainedProperties),
		cliCommand.ReadBool(constants.CanMakeOrder),
		cliCommand.ReadBool(constants.CanCancelOrder),
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
func (transactionRequest transactionRequest) GetCommonTransactionRequest() rest.CommonTransactionRequest {
	return transactionRequest.CommonTransactionRequest
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetCommonTransactionRequest().GetFrom())
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
		transactionRequest.CanMakeOrder,
		transactionRequest.CanCancelOrder,
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
func newTransactionRequest(commonTransactionRequest rest.CommonTransactionRequest, fromID string, toID string, classificationID string, maintainedProperties string, canMakeOrder bool, canCancelOrder bool, canAddMaintainer bool, canRemoveMaintainer bool, canMutateMaintainer bool) helpers.TransactionRequest {
	return transactionRequest{
		CommonTransactionRequest: commonTransactionRequest,
		FromID:                   fromID,
		ToID:                     toID,
		ClassificationID:         classificationID,
		MaintainedProperties:     maintainedProperties,
		CanMakeOrder:             canMakeOrder,
		CanCancelOrder:           canCancelOrder,
		CanAddMaintainer:         canAddMaintainer,
		CanRemoveMaintainer:      canRemoveMaintainer,
		CanMutateMaintainer:      canMutateMaintainer,
	}
}
