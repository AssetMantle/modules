/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package establish

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

//TODO make private
type Message struct {
	From             sdkTypes.AccAddress `json:"from" valid:"required~required field from missing matches(^commit[a-z0-9]{39}$)~invalid field from"`
	FromID           types.ID            `json:"fromID" valid:"required~required field fromID missing"`
	ToID             types.ID            `json:"toID" valid:"required~required field toID missing"`
	MaintainersID    types.ID            `json:"maintainersID" valid:"required~required field maintainersID missing"`
	ClassificationID types.ID            `json:"classificationID" valid:"required~required field classificationID missing"`
	Properties       types.Properties    `json:"properties" valid:"required~required field properties missing"`
	Lock             types.Height        `json:"lock" valid:"required~required field lock missing"`
	Burn             types.Height        `json:"burn" valid:"required~required field burn missing"`
}

var _ sdkTypes.Msg = Message{}

func (message Message) Route() string { return Transaction.GetModuleName() }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessage, Error.Error())
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

func newMessage(from sdkTypes.AccAddress, fromID types.ID, toID types.ID, maintainersID types.ID, classificationID types.ID, properties types.Properties, lock types.Height, burn types.Height) sdkTypes.Msg {
	return Message{
		From:             from,
		FromID:           fromID,
		ToID:             toID,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		Properties:       properties,
		Lock:             lock,
		Burn:             burn,
	}
}
