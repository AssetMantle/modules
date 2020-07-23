package mint

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
)

//TODO make private
type Message struct {
	From             sdkTypes.AccAddress `json:"from" valid:"required~Enter the FromAddress,matches(^commit[a-z0-9]{39}$)~FromAddress is Invalid"`
	MaintainersID    schema.ID           `json:"maintainers id" valid:"required~Enter the MaintainersID"`
	ClassificationID schema.ID           `json:"classification id" valid:"required~Enter the ClassificationID"`
	Properties       schema.Properties   `json:"properties" valid:"required~Enter the Properties"`
	Lock             schema.Height       `json:"lock" valid:"required~Enter the Lock"`
	Burn             schema.Height       `json:"burn" valid:"required~Enter the Burn"`
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

func newMessage(from sdkTypes.AccAddress, maintainersID schema.ID, classificationID schema.ID, properties schema.Properties, lock schema.Height, burn schema.Height) sdkTypes.Msg {
	return Message{
		From:             from,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		Properties:       properties,
		Lock:             lock,
		Burn:             burn,
	}
}
