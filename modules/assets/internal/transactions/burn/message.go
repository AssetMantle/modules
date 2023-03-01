// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) Type() string { return Transaction.GetName() }
func (message *Message) ValidateBasic() error {
	_, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		return err
	}

	if err :=
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
func newMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, assetID ids.AssetID) sdkTypes.Msg {
	return &Message{
		From:    from.String(),
		FromID:  fromID.(*base.IdentityID),
		AssetID: assetID.(*base.AssetID),
	}
}
