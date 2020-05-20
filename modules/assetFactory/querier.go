package assetFactory

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdkTypes.Querier {
	return func(context sdkTypes.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		switch path[0] {
		case constants.AssetQuery:
			return keeper.getAssetQuerier().Query(context, requestQuery)

		default:
			return nil, errors.Wrapf(constants.UnknownQueryCode, "%v", path[0])
		}
	}
}
