package assetFactory

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"
)

func NewHandler(keeper Keeper) sdkTypes.Handler {
	return func(context sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
		context = context.WithEventManager(sdkTypes.NewEventManager())

		switch message := msg.(type) {
		case mint.Message:
			return mint.Transaction.HandleMessage(context, keeper.getMintKeeper(), message)

		default:
			return nil, errors.Wrapf(constants.UnknownMessageCode, "%T", msg)
		}
	}
}
