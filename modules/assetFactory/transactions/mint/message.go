package mint

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

type message struct {
	from             sdkTypes.AccAddress
	chainID          types.ID
	maintainersID    types.ID
	classificationID types.ID
	properties       [2][]string
	lock             int
	burn             int
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return constants.ModuleName }
func (message message) Type() string  { return constants.MintTransaction }
func (message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessageCode, Error.Error())
	}
	return nil
}
func (message message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(packageCodec.MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.from}
}
