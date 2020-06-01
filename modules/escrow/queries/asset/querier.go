package interNFT

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
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
	interNFT, getInterNFTError := baseQuerier.mapper.Read(context, mapper.NewInterNFTAddress(query.Address))
	if getInterNFTError != nil {
		return nil, getInterNFTError
	}

	bytes, marshalJSONIndentError := codec.MarshalJSONIndent(packageCodec, interNFT)
	if marshalJSONIndentError != nil {
		panic(marshalJSONIndentError)
	}

	return bytes, nil
}
