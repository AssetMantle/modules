// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

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
	if err := message.ImmutableMetaProperties.ValidateBasic(); err != nil {
		return err
	}
	if err := message.MutableMetaProperties.ValidateBasic(); err != nil {
		return err
	}
	if err := message.ImmutableProperties.ValidateBasic(); err != nil {
		return err
	}
	if err := message.MutableProperties.ValidateBasic(); err != nil {
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
func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, classificationID ids.ClassificationID, immutableMetaProperties lists.PropertyList, immutableProperties lists.PropertyList, mutableMetaProperties lists.PropertyList, mutableProperties lists.PropertyList) sdkTypes.Msg {
	return &Message{
		From:                    from.String(),
		FromID:                  fromID.(*baseIDs.IdentityID),
		ClassificationID:        classificationID.(*baseIDs.ClassificationID),
		ImmutableMetaProperties: immutableMetaProperties.(*baseLists.PropertyList),
		ImmutableProperties:     immutableProperties.(*baseLists.PropertyList),
		MutableMetaProperties:   mutableMetaProperties.(*baseLists.PropertyList),
		MutableProperties:       mutableProperties.(*baseLists.PropertyList),
	}
}
