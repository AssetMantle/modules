package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/pkg/errors"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Querier interface {
	Query(sdkTypes.Context, abciTypes.RequestQuery) ([]byte, error)
}

type querier struct {
	mapper mapper.Mapper
}

var _ Querier = (*querier)(nil)

func (querier querier) Query(context sdkTypes.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	var query request
	if Error := packageCodec.UnmarshalJSON(requestQuery.Data, &query); Error != nil {
		return nil, errors.Wrap(constants.IncorrectQueryCode, Error.Error())
	}
	responseBytes := querier.mapper.QueryAssets(context, query.ID)
	return responseBytes, nil
}
func NewQuerier(mapper mapper.Mapper) Querier {
	return querier{mapper: mapper}
}
