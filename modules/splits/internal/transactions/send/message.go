// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) Type() string { return Transaction.GetName() }
func (message *Message) ValidateBasic() error {
	if _, err := sdkTypes.AccAddressFromBech32(message.From); err != nil {
		return err
	}
	if err := message.FromID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.ToID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.OwnableID.ValidateBasic(); err != nil {
		return err
	}
	if _, err := sdkTypes.NewDecFromStr(message.Value); err != nil {
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
func newMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, toID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) sdkTypes.Msg {
	return &Message{
		From:      from.String(),
		FromID:    fromID.(*baseIDs.IdentityID),
		ToID:      toID.(*baseIDs.IdentityID),
		OwnableID: ownableID.ToAnyOwnableID().(*baseIDs.AnyOwnableID),
		Value:     value.String(),
	}
}
