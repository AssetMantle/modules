package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

func storeKey(orderID types.ID) []byte {
	return append(StoreKeyPrefix, orderIDFromInterface(orderID).Bytes()...)
}

type ordersMapper interface {
	utilities.Mapper
	create(sdkTypes.Context, entities.Order)
	read(sdkTypes.Context, types.ID) entities.Order
	update(sdkTypes.Context, entities.Order)
	delete(sdkTypes.Context, types.ID)
	iterate(sdkTypes.Context, types.ID, func(entities.Order) bool)
}

type mapper struct {
	StoreKey sdkTypes.StoreKey
	Codec    *codec.Codec
}

var _ ordersMapper = (*mapper)(nil)

func (mapper mapper) InitializeMapper(codec *codec.Codec, storeKey sdkTypes.StoreKey) utilities.Mapper {
	mapper.StoreKey = storeKey
	mapper.Codec = codec
	return mapper
}
func (mapper mapper) create(context sdkTypes.Context, order entities.Order) {
	bytes, Error := mapper.Codec.MarshalBinaryBare(order)
	if Error != nil {
		panic(Error)
	}
	kvStore := context.KVStore(mapper.StoreKey)
	kvStore.Set(storeKey(order.GetID()), bytes)
}
func (mapper mapper) read(context sdkTypes.Context, orderID types.ID) entities.Order {
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
func (mapper mapper) update(context sdkTypes.Context, order entities.Order) {
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
func (mapper mapper) iterate(context sdkTypes.Context, orderID types.ID, accumulator func(entities.Order) bool) {
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

func newMapper() utilities.Mapper {
	return mapper{}
}

var Mapper = newMapper()
