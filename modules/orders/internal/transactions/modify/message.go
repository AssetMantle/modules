// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/utilities/transaction"
)

type message struct {
	From                  sdkTypes.AccAddress `json:"from" valid:"required~required field from missing"`
	FromID                ids.IdentityID      `json:"fromID" valid:"required~required field fromID missing"`
	ids.OrderID           `json:"orderID" valid:"required~required field orderID missing"`
	MakerOwnableSplit     sdkTypes.Dec       `json:"makerOwnableSplit" valid:"required~required field makerOwnableSplit missing"`
	TakerOwnableSplit     sdkTypes.Dec       `json:"takerOwnableSplit" valid:"required~required field takerOwnableSplit missing"`
	ExpiresIn             types.Height       `json:"expiresIn" valid:"required~required field expiresIn missing"`
	MutableMetaProperties lists.PropertyList `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing"`
	MutableProperties     lists.PropertyList `json:"mutableProperties" valid:"required~required field mutableProperties missing"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return sdkErrors.Wrap(errorConstants.IncorrectMessage, err.Error())
	}

	if message.TakerOwnableSplit.LTE(sdkTypes.ZeroDec()) || message.MakerOwnableSplit.LTE(sdkTypes.ZeroDec()) {
		return sdkErrors.Wrap(errorConstants.IncorrectMessage, "")
	}

	if message.ExpiresIn.Compare(baseTypes.NewHeight(0)) <= 0 {
		return sdkErrors.Wrap(errorConstants.IncorrectMessage, "")
	}

	return nil
}
func (message message) GetSignBytes() []byte {
	if len(message.MutableMetaProperties.GetList()) == 0 {
		message.MutableMetaProperties = base.NewPropertyList(nil)
	}
	if len(message.MutableProperties.GetList()) == 0 {
		message.MutableProperties = base.NewPropertyList(nil)
	}
	return sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (message) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, message{})
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

func newMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, orderID ids.OrderID, takerOwnableSplit sdkTypes.Dec, makerOwnableSplit sdkTypes.Dec, expiresIn types.Height, mutableMetaProperties lists.PropertyList, mutableProperties lists.PropertyList) sdkTypes.Msg {
	return message{
		From:                  from,
		FromID:                fromID,
		OrderID:               orderID,
		TakerOwnableSplit:     takerOwnableSplit,
		MakerOwnableSplit:     makerOwnableSplit,
		ExpiresIn:             expiresIn,
		MutableMetaProperties: mutableMetaProperties,
		MutableProperties:     mutableProperties,
	}
}
