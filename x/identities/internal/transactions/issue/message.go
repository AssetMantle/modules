// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/AssetMantle/schema/x/lists"
	baseLists "github.com/AssetMantle/schema/x/lists/base"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) Type() string { return Transaction.GetName() }
func (message *Message) ValidateBasic() error {
	if _, err := sdkTypes.AccAddressFromBech32(message.From); err != nil {
		return err
	}
	if _, err := sdkTypes.AccAddressFromBech32(message.To); err != nil {
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
func (*Message) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Message{})
}
func (message *Message) RegisterInterface(interfaceRegistry types.InterfaceRegistry) {
	interfaceRegistry.RegisterImplementations((*sdkTypes.Msg)(nil), message)
}

func messageFromInterface(msg sdkTypes.Msg) *Message {
	switch value := msg.(type) {
	case *Message:
		return value
	default:
		return &Message{}
	}
}
func messagePrototype() helpers.Message {
	return &Message{}
}
func newMessage(from sdkTypes.AccAddress, to sdkTypes.AccAddress, fromID ids.IdentityID, classificationID ids.ClassificationID, immutableMetaProperties lists.PropertyList, immutableProperties lists.PropertyList, mutableMetaProperties lists.PropertyList, mutableProperties lists.PropertyList) sdkTypes.Msg {
	return &Message{
		From:                    from.String(),
		To:                      to.String(),
		FromID:                  fromID.(*baseIDs.IdentityID),
		ClassificationID:        classificationID.(*baseIDs.ClassificationID),
		ImmutableMetaProperties: immutableMetaProperties.(*baseLists.PropertyList),
		ImmutableProperties:     immutableProperties.(*baseLists.PropertyList),
		MutableMetaProperties:   mutableMetaProperties.(*baseLists.PropertyList),
		MutableProperties:       mutableProperties.(*baseLists.PropertyList),
	}
}
