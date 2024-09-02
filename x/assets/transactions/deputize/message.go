// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) ValidateBasic() error {
	if _, err := sdkTypes.AccAddressFromBech32(message.From); err != nil {
		return err
	}
	if err := message.FromID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.ClassificationID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.ToID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.MaintainedProperties.ValidateBasic(); err != nil {
		return err
	}
	return nil
}
func (message *Message) GetSigners() []sdkTypes.AccAddress {
	from, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic(err)
	}
	return []sdkTypes.AccAddress{from}
}
func (message *Message) RegisterInterface(interfaceRegistry types.InterfaceRegistry) {
	interfaceRegistry.RegisterImplementations((*sdkTypes.Msg)(nil), message)
}

func messagePrototype() helpers.Message {
	return &Message{}
}
func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, toID ids.IdentityID, classificationID ids.ClassificationID, maintainedProperties lists.PropertyList, canMintAsset bool, canRenumerateAsset bool, canBurnAsset bool, canAddMaintainer bool, canRemoveMaintainer bool, canMutateMaintainer bool) sdkTypes.Msg {
	return &Message{
		From:                 from.String(),
		FromID:               fromID.(*baseIDs.IdentityID),
		ToID:                 toID.(*baseIDs.IdentityID),
		ClassificationID:     classificationID.(*baseIDs.ClassificationID),
		MaintainedProperties: maintainedProperties.(*baseLists.PropertyList),
		CanMintAsset:         canMintAsset,
		CanRenumerateAsset:   canRenumerateAsset,
		CanBurnAsset:         canBurnAsset,
		CanAddMaintainer:     canAddMaintainer,
		CanRemoveMaintainer:  canRemoveMaintainer,
		CanMutateMaintainer:  canMutateMaintainer,
	}
}
