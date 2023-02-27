// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) Type() string { return Transaction.GetName() }
func (message *Message) ValidateBasic() error {
	var _, err = govalidator.ValidateStruct(message)
	if err != nil {
		return errorConstants.IncorrectMessage.Wrapf(err.Error())
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
func newMessage(from sdkTypes.AccAddress, data data.Data) sdkTypes.Msg {
	return &Message{
		From: from.String(),
		Data: data.ToAnyData().(*baseData.AnyData),
	}
}
