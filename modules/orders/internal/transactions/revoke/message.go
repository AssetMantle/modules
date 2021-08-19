/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package revoke

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	sdkTypesMsgService "github.com/cosmos/cosmos-sdk/types/msgservice"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"github.com/persistenceOne/persistenceSDK/utilities/transaction"
)

var _ helpers.Message = &Message{}

func (message Message) Route() string { return module.Name }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}

	return nil
}
func (message Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(message))
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From.AsSDKTypesAccAddress()}
}
func (Message) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, &Message{})
}
func (Message) RegisterInterface(registry codecTypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdkTypes.Msg)(nil),
		&Message{},
	)
	sdkTypesMsgService.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
func messageFromInterface(msg sdkTypes.Msg) Message {
	switch value := msg.(type) {
	case *Message:
		return *value
	default:
		return Message{}
	}
}
func messagePrototype() helpers.Message {
	return &Message{}
}

func newMessage(from sdkTypes.AccAddress, fromID types.ID, toID types.ID, classificationID types.ID) sdkTypes.Msg {
	return &Message{
		From:             base.NewAccAddressFromSDKTypesAccAddress(from),
		FromID:           fromID,
		ToID:             toID,
		ClassificationID: classificationID,
	}
}
