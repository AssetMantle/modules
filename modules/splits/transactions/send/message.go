package send

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type message struct {
	From      sdkTypes.AccAddress `json:"from" valid:"required~required field fromaddress missing matches(^commit[a-z0-9]{39}$)~invalid field fromaddress"`
	ToID      types.ID            `json:"toid" valid:"required~required field toid missing"`
	OwnableID types.ID            `json:"ownableid" valid:"required~required field ownableid missing"`
	Split     sdkTypes.Dec        `json:"split" valid:"required~required field split missing"`
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

func newMessage(from sdkTypes.AccAddress, toID types.ID, ownableID types.ID, split sdkTypes.Dec) sdkTypes.Msg {
	return message{
		From:      from,
		ToID:      toID,
		OwnableID: ownableID,
		Split:     split,
	}
}
