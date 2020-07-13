package mint

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
)

//TODO make private
type Message struct {
	From sdkTypes.AccAddress
	//SellEntity types.InterNFT
	//BuyEntity  types.InterNFT
	SellCoins sdkTypes.Coin
	BuyCoins  sdkTypes.Coin
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

func newMessage(from sdkTypes.AccAddress, sellCoinDenom string, sellCoinAmount sdkTypes.Int, buyCoinDenom string, buyCoinAmount sdkTypes.Int) sdkTypes.Msg {
	return Message{
		From:      from,
		SellCoins: sdkTypes.NewCoin(sellCoinDenom, sellCoinAmount),
		BuyCoins:  sdkTypes.NewCoin(buyCoinDenom, buyCoinAmount),
	}
}
