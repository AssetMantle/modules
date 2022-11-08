// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/utilities/transaction"
)

type message struct {
	From                 sdkTypes.AccAddress  `json:"from" valid:"required~required field from missing"`
	FromID               ids.IdentityID       `json:"fromID" valid:"required~required field fromID missing"`
	ToID                 ids.IdentityID       `json:"toID" valid:"required~required field toID missing"`
	ClassificationID     ids.ClassificationID `json:"classificationID" valid:"required~required field classificationID missing"`
	MaintainedProperties lists.PropertyList   `json:"maintainedProperties" valid:"required~required field maintainedProperties missing"`
	CanMintAsset         bool                 `json:"canMintAsset"`
	CanBurnAsset         bool                 `json:"canBurnAsset"`
	CanRenumerateAsset   bool                 `json:"canRenumerateAsset"`
	CanAddMaintainer     bool                 `json:"canAddMaintainer"`
	CanRemoveMaintainer  bool                 `json:"canRemoveMaintainer"`
	CanMutateMaintainer  bool                 `json:"canMutateMaintainer"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return sdkErrors.Wrap(errorConstants.IncorrectMessage, err.Error())
	}

	return nil
}
func (message message) GetSignBytes() []byte {
	if len(message.MaintainedProperties.GetList()) == 0 {
		message.MaintainedProperties = base.NewPropertyList(nil)
	}
	return sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (message) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, message{})
}
func messageFromInterface(msg sdkTypes.Msg) message {
	switch value := msg.(type) {
	case message:
		return value
	default:
		return message{}
	}
}
func messagePrototype() helpers.Message {
	return message{}
}

func newMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, toID ids.IdentityID, classificationID ids.ClassificationID, maintainedProperties lists.PropertyList, canMintAsset bool, canBurnAsset bool, canRenumerateAsset bool, canAddMaintainer bool, canRemoveMaintainer bool, canMutateMaintainer bool) sdkTypes.Msg {
	return message{
		From:                 from,
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
