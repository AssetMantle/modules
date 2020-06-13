package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"github.com/pkg/errors"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type query struct {
	ID types.ID
}

var _ Querier = (*querier)(nil)

func (querier querier) Query(context sdkTypes.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	var query query
	if Error := packageCodec.UnmarshalJSON(requestQuery.Data, &query); Error != nil {
		return nil, errors.Wrap(constants.IncorrectQueryCode, Error.Error())
	}
	responseBytes := querier.mapper.QueryAssets(context, query.ID)
	return responseBytes, nil
}
