// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/modules/metas/module/module"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	codecUtilities "github.com/AssetMantle/modules/utilities"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) GenerateOnSuccessEvents() sdkTypes.Events {
	return nil
}
func (message *Message) Route() string   { return module.Name }
func (message *Message) GetType() string { return Transaction.GetName() }
func (message *Message) ValidateBasic() error {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return sdkErrors.Wrap(constants.IncorrectMessage, err.Error())
	}

	return nil
}
func (message *Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(codecUtilities.MakeMessageCodec(messagePrototype).MustMarshalJSON(message))
}
func (message *Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{sdkTypes.MustAccAddressFromBech32(message.From)}
}
func (*Message) RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, &Message{})
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

func newMessage(from []byte, data data.Data) sdkTypes.Msg {
	return &Message{
		From: from,
		Data: data,
	}
}
