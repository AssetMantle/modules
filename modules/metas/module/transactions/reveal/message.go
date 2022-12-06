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

type message struct {
	From sdkTypes.AccAddress `json:"from" valid:"required~required field from missing"`
	Data data.Data           `json:"data" valid:"required~required field data missing"`
}

func (message message) GenerateOnSuccessEvents() sdkTypes.Events {
	// TODO implement me
	panic("implement me")
}

func (message message) GetType() string {
	// TODO implement me
	panic("implement me")
}

func (message message) Reset() {
	// TODO implement me
	panic("implement me")
}

func (message message) String() string {
	// TODO implement me
	panic("implement me")
}

func (message message) ProtoMessage() {
	// TODO implement me
	panic("implement me")
}

var _ helpers.Message = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return sdkErrors.Wrap(constants.IncorrectMessage, err.Error())
	}

	return nil
}
func (message message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(codecUtilities.MakeMessageCodec(messagePrototype).MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (message) RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, message{})
}
func messageFromInterface(msg sdkTypes.Msg) message {
	switch value := msg.(type) {
	case message:
		return value
	default:
		return message{}
	}
}
func messagePrototype() helpers.Message {
	return message{}
}

func newMessage(from []byte, data data.Data) sdkTypes.Msg {
	return message{
		From: from,
		Data: data,
	}
}
