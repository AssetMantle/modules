package assetFactory

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mutate"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"
)

func NewHandler(keeper Keeper) sdkTypes.Handler {
	return func(context sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
		context = context.WithEventManager(sdkTypes.NewEventManager())

		switch message := msg.(type) {
		case burn.Message:
			return burn.HandleMessage(context, keeper.getBurnKeeper(), message)
		case mint.Message:
			return mint.HandleMessage(context, keeper.getMintKeeper(), message)
		case mutate.Message:
			return mutate.HandleMessage(context, keeper.getMutateKeeper(), message)

		default:
			return nil, errors.Wrapf(constants.UnknownMessageCode, "%T", msg)
		}
	}
}
