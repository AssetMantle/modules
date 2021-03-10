/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package burn

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"github.com/persistenceOne/persistenceSDK/utilities/transaction"
)

type message struct {
	From    sdkTypes.AccAddress `json:"from" valid:"required~required field from missing"`
	FromID  types.ID            `json:"fromID" valid:"required~required field fromID missing"`
	AssetID types.ID            `json:"assetID" valid:"required~required field assetID missing"`
}

var _ helpers.Message = message{}

func (Message message) Reset()         { Message = message{} }
func (Message message) String() string { return proto.CompactTextString(Message) }
func (Message message) ProtoMessage()  {}
func (Message message) Route() string  { return module.Name }
func (Message message) Type() string   { return Transaction.GetName() }
func (Message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(Message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}

	return nil
}
func (Message message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(Message))
}
func (Message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{Message.From}
}
func (message) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, message{})
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
func newMessage(from sdkTypes.AccAddress, fromID types.ID, assetID types.ID) sdkTypes.Msg {
	return message{
		From:    from,
		FromID:  fromID,
		AssetID: assetID,
	}
}
