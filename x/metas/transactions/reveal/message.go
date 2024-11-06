// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
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
	return nil
}
func (message *Message) ValidateBasic() error {
	if message.GetFromAddress() == nil {
		return constants.InvalidMessage.Wrapf("from address %s is not a valid address", message.From)
	}
	if err := message.Data.ValidateBasic(); err != nil {
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
func NewMessage(from sdkTypes.AccAddress, data data.Data) sdkTypes.Msg {
	return &Message{
		From: from.String(),
		Data: data.ToAnyData().(*baseData.AnyData),
	}
}
