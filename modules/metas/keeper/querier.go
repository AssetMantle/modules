package keeper

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/modules/metas/types"
)

// creates a querier for staking REST endpoints
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdkTypes.Querier {
	return func(ctx sdkTypes.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParameters:
			return queryParameters(ctx, k, legacyQuerierCdc)

		case types.QueryMeta:
			return queryValidator(ctx, req, k, legacyQuerierCdc)

		default:
			return nil, sdkErrors.Wrapf(sdkErrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

func queryValidator(ctx sdkTypes.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var metaRequest types.QueryMetaRequest

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &metaRequest)
	if err != nil {
		return nil, sdkErrors.Wrap(sdkErrors.ErrJSONUnmarshal, err.Error())
	}

	meta, err := k.GetMeta(ctx, types.MetaIDFromInterface(base.NewID(metaRequest.MetaID)))
	if err != nil {
		return nil, errors.EntityNotFound
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, meta)
	if err != nil {
		return nil, sdkErrors.Wrap(sdkErrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryParameters(ctx sdkTypes.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParameters(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkErrors.Wrap(sdkErrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
