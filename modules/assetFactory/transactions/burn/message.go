package burn

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

type Message struct {
	from    sdkTypes.AccAddress
	assetID types.ID
}

var _ sdkTypes.Msg = Message{}

func (message Message) Route() string { return constants.ModuleName }
func (message Message) Type() string  { return constants.BurnTransaction }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessageCode, Error.Error())
	}
	return nil
}
func (message Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(packageCodec.MustMarshalJSON(message))
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.from}
}
