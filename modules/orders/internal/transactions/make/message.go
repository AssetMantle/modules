// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	typesSchema "github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) Type() string { return Transaction.GetName() }
func (message *Message) ValidateBasic() error {
	var _, err = govalidator.ValidateStruct(message)
	if err != nil {
		return sdkErrors.Wrap(errorConstants.IncorrectMessage, err.Error())
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
func (message *Message) GenerateOnSuccessEvents() sdkTypes.Events {
	return sdkTypes.Events{sdkTypes.NewEvent(
		sdkTypes.EventTypeMessage,
		sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, module.Name+"."+message.Type()),
	)}
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
func newMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, classificationID ids.ClassificationID, takerID ids.IdentityID, makerOwnableID ids.OwnableID, takerOwnableID ids.OwnableID, expiresIn typesSchema.Height, makerOwnableSplit sdkTypes.Dec, takerOwnableSplit sdkTypes.Dec, immutableMetaProperties lists.PropertyList, immutableProperties lists.PropertyList, mutableMetaProperties lists.PropertyList, mutableProperties lists.PropertyList) sdkTypes.Msg {
	return &Message{
		From:                    from,
		FromID:                  fromID,
		ClassificationID:        classificationID,
		TakerID:                 takerID,
		MakerOwnableID:          makerOwnableID,
		TakerOwnableID:          takerOwnableID,
		ExpiresIn:               expiresIn,
		MakerOwnableSplit:       makerOwnableSplit,
		TakerOwnableSplit:       takerOwnableSplit,
		ImmutableMetaProperties: immutableMetaProperties,
		ImmutableProperties:     immutableProperties,
		MutableMetaProperties:   mutableMetaProperties,
		MutableProperties:       mutableProperties,
	}
}
