/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	messageUtilities "github.com/persistenceOne/persistenceSDK/utilities/message"
)

type message struct {
	From                    sdkTypes.AccAddress  `json:"from" valid:"required~required field from missing"`
	FromID                  types.ID             `json:"fromID" valid:"required~required field fromID missing"`
	ClassificationID        types.ID             `json:"classificationID" valid:"required~required field classificationID missing matches(^[A-Za-z]$)~invalid field classificationID"`
	MakerOwnableID          types.ID             `json:"makerOwnableID" valid:"required~required field makerOwnableID missing"`
	TakerOwnableID          types.ID             `json:"takerOwnableID" valid:"required~required field takerOwnableID missing"`
	ExpiresIn               types.Height         `json:"expiresIn" valid:"required~required field expiresIn missing"`
	MakerOwnableSplit       sdkTypes.Dec         `json:"makerOwnableSplit" valid:"required~required field makerOwnableSplit missing"`
	ImmutableMetaProperties types.MetaProperties `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing matches(^[A-Za-z]$)~invalid field immutableMetaProperties"`
	ImmutableProperties     types.Properties     `json:"immutableProperties" valid:"required~required field immutableProperties missing matches(^[A-Za-z]$)~invalid field immutableProperties"`
	MutableMetaProperties   types.MetaProperties `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing matches(^[A-Za-z]$)~invalid field mutableMetaProperties"`
	MutableProperties       types.Properties     `json:"mutableProperties" valid:"required~required field mutableProperties missing matches(^[A-Za-z]$)~invalid field mutableProperties"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}
	return nil
}
func (message message) GetSignBytes() []byte {
	messageCodec := messageUtilities.RegisterCodec(messagePrototype)
	return sdkTypes.MustSortJSON(messageCodec.MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (message) RegisterCodec(codec *codec.Codec) {
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

func newMessage(from sdkTypes.AccAddress, fromID types.ID, classificationID types.ID, makerOwnableID types.ID, takerOwnableID types.ID, expiresIn types.Height, makerOwnableSplit sdkTypes.Dec, immutableMetaProperties types.MetaProperties, immutableProperties types.Properties, mutableMetaProperties types.MetaProperties, mutableProperties types.Properties) sdkTypes.Msg {
	return message{
		From:                    from,
		FromID:                  fromID,
		ClassificationID:        classificationID,
		MakerOwnableID:          makerOwnableID,
		TakerOwnableID:          takerOwnableID,
		ExpiresIn:               expiresIn,
		MakerOwnableSplit:       makerOwnableSplit,
		ImmutableMetaProperties: immutableMetaProperties,
		ImmutableProperties:     immutableProperties,
		MutableMetaProperties:   mutableMetaProperties,
		MutableProperties:       mutableProperties,
	}
}
