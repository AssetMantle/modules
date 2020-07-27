package issue

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type message struct {
	From             sdkTypes.AccAddress `json:"from" valid:"required~required field from missing matches(^commit[a-z0-9]{39}$)~invalid field from"`
	To               sdkTypes.AccAddress `json:"to" valid:"required~required field to missing matches(^commit[a-z0-9]{39}$)~invalid field to"`
	MaintainersID    types.ID            `json:"maintainersID" valid:"required~required field maintainersID missing"`
	ClassificationID types.ID            `json:"classificationID" valid:"required~required field classificationID missing"`
	Properties       types.Properties    `json:"properties" valid:"required~required field properties missing"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return Transaction.GetModuleName() }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessage, Error.Error())
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

func newMessage(from sdkTypes.AccAddress, to sdkTypes.AccAddress, maintainersID types.ID, classificationID types.ID, properties types.Properties) sdkTypes.Msg {
	return message{
		From:             from,
		To:               to,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		Properties:       properties,
	}
}
