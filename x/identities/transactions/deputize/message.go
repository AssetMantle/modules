// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) GetFromAddress() sdkTypes.AccAddress {
	from, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil || from.Empty() {
		// NOTE: This should never happen as the message is validated before it is sent
		return nil
	}
	return from
}
func (message *Message) GetFromIdentityID() ids.IdentityID {
	return message.FromID
}
func (message *Message) ValidateBasic() error {
	if message.GetFromAddress() == nil {
		return constants.InvalidMessage.Wrapf("from address %s is not a valid address", message.From)
	}
	if err := message.GetFromIdentityID().ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.ClassificationID.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.MaintainedProperties.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.ToID.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	return nil
}
func (message *Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.GetFromAddress()}
}
func (message *Message) RegisterInterface(interfaceRegistry types.InterfaceRegistry) {
	interfaceRegistry.RegisterImplementations((*sdkTypes.Msg)(nil), message)
}

func messagePrototype() helpers.Message {
	return &Message{}
}
func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, toID ids.IdentityID, classificationID ids.ClassificationID, maintainedProperties lists.PropertyList, canIssueIdentity bool, canQuashIdentity bool, canAddMaintainer bool, canRemoveMaintainer bool, canMutateMaintainer bool) sdkTypes.Msg {
	return &Message{
		From:                 from.String(),
		FromID:               fromID.(*baseIDs.IdentityID),
		ToID:                 toID.(*baseIDs.IdentityID),
		ClassificationID:     classificationID.(*baseIDs.ClassificationID),
		MaintainedProperties: maintainedProperties.(*baseLists.PropertyList),
		CanIssueIdentity:     canIssueIdentity,
		CanQuashIdentity:     canQuashIdentity,
		CanAddMaintainer:     canAddMaintainer,
		CanRemoveMaintainer:  canRemoveMaintainer,
		CanMutateMaintainer:  canMutateMaintainer,
	}
}
