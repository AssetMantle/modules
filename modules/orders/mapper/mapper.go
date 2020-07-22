package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

func storeKey(assetID types.ID) []byte {
	return append(StoreKeyPrefix, orderIDFromInterface(assetID).Bytes()...)
}

type ordersMapper interface {
	types.Mapper
	create(sdkTypes.Context, types.Order)
	read(sdkTypes.Context, types.ID) types.Order
	update(sdkTypes.Context, types.Order)
	delete(sdkTypes.Context, types.ID)
	iterate(sdkTypes.Context, types.ID, func(types.Order) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ ordersMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) types.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, order types.Order) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(order)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(order.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, orderID types.ID) types.Order {
	kvStore := context.KVStore(mapper.StoreKey)
	bytes := kvStore.Get(storeKey(orderID))
	if bytes == nil {
		return order{}
	}
	order := order{}
	Error := mapper.Codec.UnmarshalBinaryBare(bytes, &order)
	if Error != nil {
		panic(Error)
	}
	return order
}
func (mapper mapper) update(context sdkTypes.Context, order types.Order) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(order)
	if Error != nil {
		panic(Error)
	}
	assetID := order.GetID()
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, orderID types.ID) {
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Delete(storeKey(orderID))
}
func (mapper mapper) iterate(context sdkTypes.Context, orderID types.ID, accumulator func(types.Order) bool) {
	store := context.KVStore(mapper.StoreKey)
	kvStorePrefixIterator := sdkTypes.KVStorePrefixIterator(store, storeKey(orderID))

	defer kvStorePrefixIterator.Close()
	for ; kvStorePrefixIterator.Valid(); kvStorePrefixIterator.Next() {
		asset := order{}
		Error := mapper.Codec.UnmarshalBinaryBare(kvStorePrefixIterator.Value(), &asset)
		if Error != nil {
			panic(Error)
		}
		if accumulator(asset) {
			break
		}
	}
}

func newMapper() types.Mapper {
	return mapper{}
}

var Mapper = newMapper()
