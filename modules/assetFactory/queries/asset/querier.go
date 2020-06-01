package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Querier interface {
	Query(sdkTypes.Context, abciTypes.RequestQuery) ([]byte, error)
}

func NewQuerier(mapper mapper.Mapper) Querier {
	return querier{mapper: mapper}
}

type querier struct {
	mapper mapper.Mapper
}
