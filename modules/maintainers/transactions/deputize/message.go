/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

//TODO make private
type Message struct {
	From             sdkTypes.AccAddress `json:"from" valid:"required~required field from missing matches(^commit[a-z0-9]{39}$)~invalid field from"`
	FromID           types.ID            `json:"fromID" valid:"required~required field fromID missing"`
	ToID             types.ID            `json:"toID" valid:"required~required field toID missing"`
	ClassificationID types.ID            `json:"classificationID" valid:"required~required field classificationID missing"`
	MaintainedTraits types.Properties    `json:"maintainedTraits" valid:"required~required field maintainedTraits missing"`
	AddMaintainer    bool                `json:"addMaintainer" valid:"required field addMaintainer missing"`
	RemoveMaintainer bool                `json:"removeMaintainer" valid:"required field removeMaintainer missing"`
	MutateMaintainer bool                `json:"mutateMaintainer" valid:"required field mutateMaintainer missing"`
}

var _ sdkTypes.Msg = Message{}

func (message Message) Route() string { return Transaction.GetModuleName() }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}
	return nil
}
func (message Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(packageCodec.MustMarshalJSON(message))
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}

func messageFromInterface(msg sdkTypes.Msg) Message {
	switch value := msg.(type) {
	case Message:
		return value
	default:
		return Message{}
	}
}

func newMessage(from sdkTypes.AccAddress, fromID types.ID, toID types.ID, classificationID types.ID, maintainedTraits types.Properties, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) sdkTypes.Msg {
	return Message{
		From:             from,
		FromID:           fromID,
		ToID:             toID,
		ClassificationID: classificationID,
		MaintainedTraits: maintainedTraits,
		AddMaintainer:    addMaintainer,
		RemoveMaintainer: removeMaintainer,
		MutateMaintainer: mutateMaintainer,
	}
}
