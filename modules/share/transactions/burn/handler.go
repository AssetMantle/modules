package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/share/constants"
)

func HandleMessage(context sdkTypes.Context, keeper Keeper, message Message) (*sdkTypes.Result, error) {

	if Error := keeper.transact(context, message); Error != nil {
		return nil, Error
	}

	context.EventManager().EmitEvent(
		sdkTypes.NewEvent(
			sdkTypes.EventTypeMessage,
			sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, constants.AttributeValueCategory),
		),
	)

	return &sdkTypes.Result{Events: context.EventManager().ABCIEvents()}, nil
}
