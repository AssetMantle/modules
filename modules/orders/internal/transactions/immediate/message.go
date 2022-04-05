// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	xprtErrors "github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/utilities/transaction"
)

type message struct {
	From                    sdkTypes.AccAddress  `json:"from" valid:"required~required field from missing"`
	FromID                  types.ID             `json:"fromID" valid:"required~required field fromID missing"`
	ClassificationID        types.ID             `json:"classificationID" valid:"required~required field classificationID missing"`
	MakerOwnableID          types.ID             `json:"makerOwnableID" valid:"required~required field makerOwnableID missing"`
	TakerOwnableID          types.ID             `json:"takerOwnableID" valid:"required~required field takerOwnableID missing"`
	ExpiresIn               types.Height         `json:"expiresIn" valid:"required~required field expiresIn missing"`
	MakerOwnableSplit       sdkTypes.Dec         `json:"makerOwnableSplit" valid:"required~required field makerOwnableSplit missing"`
	TakerOwnableSplit       sdkTypes.Dec         `json:"takerOwnableSplit" valid:"required~required field takerOwnableSplit missing"`
	ImmutableMetaProperties types.MetaProperties `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing"`
	ImmutableProperties     types.Properties     `json:"immutableProperties" valid:"required~required field immutableProperties missing"`
	MutableMetaProperties   types.MetaProperties `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing"`
	MutableProperties       types.Properties     `json:"mutableProperties" valid:"required~required field mutableProperties missing"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}

	if message.MakerOwnableID.Compare(message.TakerOwnableID) == 0 {
		return errors.Wrap(xprtErrors.IncorrectMessage, "")
	}

	if message.TakerOwnableSplit.LTE(sdkTypes.ZeroDec()) || message.MakerOwnableSplit.LTE(sdkTypes.ZeroDec()) {
		return errors.Wrap(xprtErrors.IncorrectMessage, "")
	}

	return nil
}
func (message message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (message) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, message{})
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

func newMessage(from sdkTypes.AccAddress, fromID types.ID, classificationID types.ID, makerOwnableID types.ID, takerOwnableID types.ID, expiresIn types.Height, makerOwnableSplit sdkTypes.Dec, takerOwnableSplit sdkTypes.Dec, immutableMetaProperties types.MetaProperties, immutableProperties types.Properties, mutableMetaProperties types.MetaProperties, mutableProperties types.Properties) sdkTypes.Msg {
	return message{
		From:                    from,
		FromID:                  fromID,
		ClassificationID:        classificationID,
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
