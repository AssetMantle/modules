package share

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/share/constants"
	"github.com/persistenceOne/persistenceSDK/modules/share/mapper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Querier interface {
	Query(sdkTypes.Context, abciTypes.RequestQuery) ([]byte, error)
}

type baseQuerier struct {
	mapper mapper.Mapper
}

func NewQuerier(mapper mapper.Mapper) Querier {
	return baseQuerier{mapper: mapper}
}

var _ Querier = (*baseQuerier)(nil)

func (baseQuerier baseQuerier) Query(context sdkTypes.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	var query query
	if Error := packageCodec.UnmarshalJSON(requestQuery.Data, &query); Error != nil {
		return nil, errors.Wrap(constants.IncorrectQueryCode, Error.Error())
	}
	share, getShareError := baseQuerier.mapper.Read(context, mapper.NewShareAddress(query.Address))
	if getShareError != nil {
		return nil, getShareError
	}

	bytes, marshalJSONIndentError := codec.MarshalJSONIndent(packageCodec, share)
	if marshalJSONIndentError != nil {
		panic(marshalJSONIndentError)
	}

	return bytes, nil
}
