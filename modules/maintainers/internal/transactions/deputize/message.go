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

type message struct {
	From             sdkTypes.AccAddress `json:"from" valid:"required~required field from missing matches(^commit[a-z0-9]{39}$)~invalid field from"`
	FromID           types.ID            `json:"fromID" valid:"required~required field fromID missing"`
	ToID             types.ID            `json:"toID" valid:"required~required field toID missing"`
	ClassificationID types.ID            `json:"classificationID" valid:"required~required field classificationID missing"`
	MaintainedTraits types.Properties    `json:"maintainedTraits" valid:"required~required field maintainedTraits missing"`
	AddMaintainer    bool                `json:"addMaintainer" valid:"required~required field addMaintainer missing"`
	RemoveMaintainer bool                `json:"removeMaintainer" valid:"required~required field removeMaintainer missing"`
	MutateMaintainer bool                `json:"mutateMaintainer" valid:"required~required field mutateMaintainer missing"`
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

func newMessage(from sdkTypes.AccAddress, fromID types.ID, toID types.ID, classificationID types.ID, maintainedTraits types.Properties, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) sdkTypes.Msg {
	return message{
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
