package mutate

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

type Message struct {
	From       sdkTypes.AccAddress
	AssetID    types.ID
	Properties types.Properties
	Lock       types.Height
	Burn       types.Height
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

func messagePrototype() sdkTypes.Msg {
	return Message{}
}

func messageFromInterface(msg sdkTypes.Msg) Message {
	switch value := msg.(type) {
	case Message:
		return value
	default:
		return Message{}
	}
}

func NewMessage(from sdkTypes.AccAddress, assetID types.ID, properties types.Properties, lock types.Height, burn types.Height) sdkTypes.Msg {
	return Message{
		From:       from,
		AssetID:    assetID,
		Properties: properties,
		Lock:       lock,
		Burn:       burn,
	}
}
