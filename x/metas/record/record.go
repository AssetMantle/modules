package record

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/key"
	"github.com/AssetMantle/modules/x/metas/mappable"
)

func (record *Record) GetKey() helpers.Key {
	return record.Key
}
func (record *Record) GetMappable() helpers.Mappable {
	return record.Mappable
}
func (record *Record) WithKey(Key helpers.Key) helpers.Record {
	record.Key = Key.(*key.Key)
	record.Mappable = mappable.Prototype().(*mappable.Mappable)
	return record
}
func (record *Record) WithMappable(Mappable helpers.Mappable) helpers.Record {
	record.Key = Mappable.GenerateKey().(*key.Key)
	record.Mappable = Mappable.(*mappable.Mappable)
	return record
}
func (record *Record) ReadFromIterator(iterator sdkTypes.Iterator) helpers.Record {
	Mappable := record.GetMappable()
	base.CodecPrototype().MustUnmarshal(iterator.Value(), Mappable)
	record.WithMappable(Mappable)
	return record
}
func (record *Record) Read(kvStore sdkTypes.KVStore) helpers.Record {
	Bytes := kvStore.Get(record.GetKey().GenerateStoreKeyBytes())
	if Bytes == nil {
		return Prototype()
	}
	Mappable := record.GetMappable()
	base.CodecPrototype().MustUnmarshal(Bytes, Mappable)
	record.WithMappable(Mappable)
	return record
}
func (record *Record) Write(kvStore sdkTypes.KVStore) helpers.Record {
	Bytes := base.CodecPrototype().MustMarshal(record.GetMappable())
	kvStore.Set(record.GetKey().GenerateStoreKeyBytes(), Bytes)
	return record
}
func (record *Record) Delete(kvStore sdkTypes.KVStore) {
	kvStore.Delete(record.GetKey().GenerateStoreKeyBytes())
}

func RecordsFromInterface(records []helpers.Record) []*Record {
	Records := make([]*Record, len(records))
	for index, record := range records {
		Records[index] = record.(*Record)
	}
	return Records
}

func Prototype() helpers.Record {
	return &Record{
		Key:      key.Prototype().(*key.Key),
		Mappable: mappable.Prototype().(*mappable.Mappable),
	}
}
