package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type queryKeeper struct {
	mapper mapper.Mapper
}

var _ types.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Query(context sdkTypes.Context, QueryRequest types.QueryRequest) ([]byte, error) {
	query := QueryRequest.(queryRequest)
	bytes, Error := packageCodec.MarshalJSON(queryResponse{Assets: mapper.NewAssets(queryKeeper.mapper, context).Read(query.ID)})
	if Error != nil {
		return nil, Error
	}
	return bytes, nil
}

func NewQueryKeeper(mapper mapper.Mapper) types.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
