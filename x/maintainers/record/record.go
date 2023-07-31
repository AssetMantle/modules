package record

import (
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/documents/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/maintainers/key"
	"github.com/AssetMantle/modules/x/maintainers/mappable"
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
	Bytes := iterator.Value()
	if Bytes == nil {
		return Prototype()
	}

	Mappable := record.GetMappable()
	base.CodecPrototype().MustUnmarshal(iterator.Value(), Mappable)
	record.Mappable = Mappable.(*mappable.Mappable)
	return record
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

func NewRecord(Maintainer documents.Maintainer) helpers.Record {
	return &Record{
		Key: key.NewKey(baseIDs.NewMaintainerID(constants.MaintainerClassificationID,
			baseQualified.NewImmutables(baseLists.NewPropertyList(
				baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(baseDocuments.NewMaintainerFromDocument(Maintainer).GetMaintainedClassificationID())),
				baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(baseDocuments.NewMaintainerFromDocument(Maintainer).GetIdentityID())),
			)))).(*key.Key),
		Mappable: mappable.NewMappable(Maintainer).(*mappable.Mappable),
	}
}
