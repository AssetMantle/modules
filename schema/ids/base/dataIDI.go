package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.DataID = (*DataIDI)(nil)

func (dataIDI *DataIDI) Compare(listable traits.Listable) int {
	return dataIDI.Impl.(ids.DataID).Compare(listable)
}

func (dataIDI *DataIDI) Bytes() []byte {
	return dataIDI.Impl.(ids.DataID).Bytes()
}

func (dataIDI *DataIDI) GetHashID() ids.HashID {
	return dataIDI.Impl.(ids.DataID).GetHashID()
}

func (dataIDI *DataIDI) IsDataID() {
}
