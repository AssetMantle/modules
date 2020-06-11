package mutate

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

type Message struct {
	from             sdkTypes.AccAddress
	chainID          types.ID
	maintainersID    types.ID
	classificationID types.ID
	properties       types.Properties
	lock             types.Height
	burn             types.Height
}

var _ sdkTypes.Msg = Message{}

func (message Message) Route() string { return constants.ModuleName }
func (message Message) Type() string  { return constants.MutateTransaction }
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
