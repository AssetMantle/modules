package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

func storeKey(orderID schema.ID) []byte {
	return append(StoreKeyPrefix, orderIDFromInterface(orderID).Bytes()...)
}

type ordersMapper interface {
	utility.Mapper
	create(sdkTypes.Context, schema.Order)
	read(sdkTypes.Context, schema.ID) schema.Order
	update(sdkTypes.Context, schema.Order)
	delete(sdkTypes.Context, schema.ID)
	iterate(sdkTypes.Context, schema.ID, func(schema.Order) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ ordersMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) utility.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, order schema.Order) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(order)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(order.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, orderID schema.ID) schema.Order {
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
func (mapper mapper) update(context sdkTypes.Context, order schema.Order) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(order)
	if Error != nil {
		panic(Error)
	}
	assetID := order.GetID()
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(assetID), bytes)
}
func (mapper mapper) delete(context sdkTypes.Context, orderID schema.ID) {
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Delete(storeKey(orderID))
}
func (mapper mapper) iterate(context sdkTypes.Context, orderID schema.ID, accumulator func(schema.Order) bool) {
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

func newMapper() utility.Mapper {
	return mapper{}
}

var Mapper = newMapper()
