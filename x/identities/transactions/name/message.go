// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package name

import (
	"github.com/AssetMantle/modules/helpers/constants"
	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) Type() string { return Transaction.GetName() }
func (message *Message) GetFromAddress() sdkTypes.AccAddress {
	from, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic(err)
	}
	return from
}
func (message *Message) ValidateBasic() error {
	if _, err := sdkTypes.AccAddressFromBech32(message.From); err != nil {
		return err
	}

	if message.Name == nil || message.Name.AsString() == "" {
		return constants.InvalidRequest.Wrapf("name cannot be empty")
	}

	if err := message.Name.ValidateBasic(); err != nil {
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
func NewMessage(from sdkTypes.AccAddress, name ids.StringID) sdkTypes.Msg {
	return &Message{
		From: from.String(),
		Name: name.(*baseIDs.StringID),
	}
}
