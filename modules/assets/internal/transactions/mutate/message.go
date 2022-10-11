// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/utilities/transaction"
)

type message struct {
	From                  sdkTypes.AccAddress    `json:"from" valid:"required~required field from missing"`
	FromID                ids.IdentityID         `json:"fromID" valid:"required~required field fromID missing"`
	AssetID               ids.AssetID            `json:"assetID" valid:"required~required field assetID missing"`
	MutableMetaProperties lists.MetaPropertyList `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing"`
	MutableProperties     lists.PropertyList     `json:"mutableProperties" valid:"required~required field mutableProperties missing"`
}

var _ helpers.Message = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return sdkErrors.Wrap(errors.IncorrectMessage, err.Error())
	}

	return nil
}
func (message message) GetSignBytes() []byte {
	if len(message.MutableProperties.GetList()) == 0 {
		message.MutableProperties = base.NewPropertyList(nil)
	}
	if len(message.MutableMetaProperties.GetList()) == 0 {
		message.MutableMetaProperties = base.NewMetaPropertyList(nil)
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
func newMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, assetID ids.AssetID, mutableMetaProperties lists.MetaPropertyList, mutableProperties lists.PropertyList) sdkTypes.Msg {
	return message{
		From:                  from,
		FromID:                fromID,
		AssetID:               assetID,
		MutableMetaProperties: mutableMetaProperties,
		MutableProperties:     mutableProperties,
	}
}
