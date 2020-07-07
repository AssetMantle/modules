package assets

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mutate"
)

func NewHandler(keeper Keeper) sdkTypes.Handler {
	return func(context sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
		context = context.WithEventManager(sdkTypes.NewEventManager())

		switch message := msg.(type) {
		case burn.Message:
			return burn.Transaction.HandleMessage(context, keeper.getBurnKeeper(), message)
		case mint.Message:
			return mint.Transaction.HandleMessage(context, keeper.getMintKeeper(), message)
		case mutate.Message:
			return mutate.Transaction.HandleMessage(context, keeper.getMutateKeeper(), message)

		default:
			return nil, errors.Wrapf(constants.UnknownMessage, "%T", msg)
		}
	}
}
