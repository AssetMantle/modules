package record

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/key"
	"github.com/AssetMantle/modules/x/metas/mappable"
	"github.com/AssetMantle/schema/data"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
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
func (record *Record) ReadFromIterator(iterator sdkTypes.Iterator) helpers.Record {
	return NewRecord(mappable.GetData(helpers.ReadMappableFromIterator(iterator, record.GetMappable())))
}
func (record *Record) Read(kvStore sdkTypes.KVStore) helpers.Record {
	if record.GetKey() == nil || len(record.GetKey().GeneratePrefixedStoreKeyBytes()) == 0 {
		return Prototype()
	}
	Bytes := kvStore.Get(record.GetKey().GeneratePrefixedStoreKeyBytes())
	if Bytes == nil {
		return Prototype()
	}

	Mappable := record.GetMappable()
	base.CodecPrototype().MustUnmarshal(Bytes, Mappable)
	record.Mappable = Mappable.(*mappable.Mappable)

	return record
}
func (record *Record) Write(kvStore sdkTypes.KVStore) helpers.Record {
	Bytes := base.CodecPrototype().MustMarshal(record.GetMappable())
	kvStore.Set(record.GetKey().GeneratePrefixedStoreKeyBytes(), Bytes)
	return record
}
func (record *Record) Delete(kvStore sdkTypes.KVStore) {
	kvStore.Delete(record.GetKey().GeneratePrefixedStoreKeyBytes())
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

func NewRecord(data data.Data) *Record {
	return &Record{
		Key:      key.NewKey(baseIDs.GenerateDataID(data)).(*key.Key),
		Mappable: mappable.NewMappable(data).(*mappable.Mappable),
	}
}
