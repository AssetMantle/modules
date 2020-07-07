package assets

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/queries/asset"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdkTypes.Querier {
	return func(context sdkTypes.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		switch path[0] {
		case asset.Query.GetName():
			return asset.Query.HandleMessage(context, keeper.getAssetQuerier(), requestQuery)

		default:
			return nil, errors.Wrapf(constants.UnknownQuery, "%v", path[0])
		}
	}
}
