package make

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

//TODO make private
type Message struct {
	From          sdkTypes.AccAddress
	MaintainersID types.ID
	MakerID       types.ID
	TakerID       types.ID
	MakerSplit    sdkTypes.Dec
	MakerSplitID  types.ID
	ExchangeRate  sdkTypes.Dec
	TakerSplitID  types.ID
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

func newMessage(from sdkTypes.AccAddress, maintainersID types.ID,
	fromID types.ID, toID types.ID, makerSplit sdkTypes.Dec, makerSplitID types.ID,
	exchangeRate sdkTypes.Dec, takerSplitID types.ID) sdkTypes.Msg {
	return Message{
		From:          from,
		MaintainersID: maintainersID,
		MakerID:       fromID,
		TakerID:       toID,
		MakerSplit:    makerSplit,
		MakerSplitID:  makerSplitID,
		ExchangeRate:  exchangeRate,
		TakerSplitID:  takerSplitID,
	}
}
