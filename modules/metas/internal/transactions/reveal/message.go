// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/utilities/transaction"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

//type message struct {
//	From sdkTypes.AccAddress `json:"from" valid:"required~required field from missing"`
//	Data data.Data           `json:"data" valid:"required~required field data missing"`
//}

var _ helpers.Message = &Message{}

func (message *Message) RegisterInterface(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdkTypes.Msg)(nil),
		&Message{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Transaction_serviceDesc)
}

func (message *Message) GenerateOnSuccessEvents() sdkTypes.Events {
	return nil
}
func (message *Message) Route() string { return module.Name }
func (message *Message) Type() string  { return Transaction.GetName() }
func (message *Message) ValidateBasic() error {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return sdkErrors.Wrap(constants.IncorrectMessage, err.Error())
	}

	return nil
}
func (message *Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(message))
}
func (message *Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (*Message) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, &Message{})
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
		From: from,
		Data: data,
	}
}
