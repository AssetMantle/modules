// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	xprtErrors "github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/utilities/transaction"
)

type message struct {
	From       sdkTypes.AccAddress `json:"from" valid:"required~required field from missing"`
	To         sdkTypes.AccAddress `json:"to" valid:"required~required field to missing, matches(^[a-z0-9]*$)~invalid field to"`
	IdentityID ids.ID              `json:"identityID" valid:"required~required field identityID missing"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, err.Error())
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
func newMessage(from sdkTypes.AccAddress, to sdkTypes.AccAddress, identityID ids.ID) sdkTypes.Msg {
	return message{
		From:       from,
		To:         to,
		IdentityID: identityID,
	}
}
