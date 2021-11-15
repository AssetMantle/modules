/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type message struct {
	From                sdkTypes.AccAddress  `json:"from" valid:"required~required field from missing "`
	FromID              types.ID             `json:"fromID" valid:"required~required field fromID missing"`
	ImmutableMetaTraits types.MetaProperties `json:"immutableMetaTraits" valid:"required~required field immutableMetaTraits missing"`
	ImmutableTraits     types.Properties     `json:"immutableTraits" valid:"required~required field immutableTraits missing"`
	MutableMetaTraits   types.MetaProperties `json:"mutableMetaTraits" valid:"required~required field mutableMetaTraits missing"`
	MutableTraits       types.Properties     `json:"mutableTraits" valid:"required~required field mutableTraits missing"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return Transaction.GetModuleName() }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}
	return nil
}
func (message message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(packageCodec.MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}

func messageFromInterface(msg sdkTypes.Msg) message {
	switch value := msg.(type) {
	case message:
		return value
	default:
		return message{}
	}
}

func newMessage(from sdkTypes.AccAddress, fromID types.ID, immutableMetaTraits types.MetaProperties, immutableTraits types.Properties, mutableMetaTraits types.MetaProperties, mutableTraits types.Properties) sdkTypes.Msg {
	return message{
		From:                from,
		FromID:              fromID,
		ImmutableMetaTraits: immutableMetaTraits,
		ImmutableTraits:     immutableTraits,
		MutableMetaTraits:   mutableMetaTraits,
		MutableTraits:       mutableTraits,
	}
}
