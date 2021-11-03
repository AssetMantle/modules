package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	metaTypes "github.com/persistenceOne/persistenceSDK/modules/metas/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Keeper struct {
	storeKey      sdkTypes.StoreKey
	cdc           codec.BinaryMarshaler
	paramSubspace paramTypes.Subspace
}

func (k Keeper) GetMeta(ctx sdkTypes.Context, id metaTypes.MetaID) (metaTypes.Meta, error) {
	metaID := metaTypes.MetaIDFromInterface(id)
	store := ctx.KVStore(k.storeKey)
	value := store.Get(metaID.GenerateStoreKeyBytes())
	if value == nil {
		return metaTypes.Meta{}, errors.ErrKeyNotFound
	}
	var meta metaTypes.Meta
	err := k.cdc.UnmarshalBinaryBare(value, &meta)
	if err != nil {
		return metaTypes.Meta{}, errors.ErrKeyNotFound
	}
	return meta, nil
}

func (k Keeper) SetMeta(ctx sdkTypes.Context, meta metaTypes.Meta) {
	metaID := meta.GetMetaID()
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&meta)
	store.Set(metaID.GenerateStoreKeyBytes(), bz)
}

func (k Keeper) GetParameters(ctx sdkTypes.Context) metaTypes.Parameters {
	var maxStringLengthParameter metaTypes.Parameter
	k.paramSubspace.Get(ctx, metaTypes.MaxStringLengthID.Bytes(), &maxStringLengthParameter)
	return metaTypes.NewParameters(maxStringLengthParameter)
}

func (k Keeper) GetParameter(ctx sdkTypes.Context, id types.ID) metaTypes.Parameter {
	var parameter metaTypes.Parameter
	k.paramSubspace.Get(ctx, id.Bytes(), &parameter)
	return parameter
}

func (k Keeper) SetParameters(ctx sdkTypes.Context, parameters metaTypes.Parameters) {
	k.paramSubspace.SetParamSet(ctx, &parameters)
}

func (k Keeper) GetAllMetas(ctx sdkTypes.Context) (metas []metaTypes.Meta) {
	store := ctx.KVStore(k.storeKey)

	emptyMeta := metaTypes.Meta{}
	iterator := sdkTypes.KVStorePrefixIterator(store, emptyMeta.GetMetaID().GenerateStoreKeyBytes())
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var meta metaTypes.Meta
		err := k.cdc.UnmarshalBinaryBare(iterator.Value(), &meta)
		if err != nil {
			panic(err)
		}
		metas = append(metas, meta)
	}

	return metas
}

func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdkTypes.StoreKey, paramStore paramTypes.Subspace) Keeper {
	if !paramStore.HasKeyTable() {
		paramStore = paramStore.WithKeyTable(metaTypes.ParamsKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramSubspace: paramStore,
	}
}
