package record

import (
	"github.com/AssetMantle/modules/helpers"
)

var _ helpers.Record = (*Record)(nil)

func (record *Record) GetKey() helpers.Key {
	return record.Key
}
func (record *Record) GetMappable() helpers.Mappable {
	return record.Mappable
}

func RecordsFromInterface(records []helpers.Record) []*Record {
	Records := make([]*Record, len(records))
	for index, record := range records {
		Records[index] = record.(*Record)
	}
	return Records
}
