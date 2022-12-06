package base

import (
	ids2 "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type dataIDI ids2.HashID

var _ ids.DataID = (*dataIDI)(nil)

func (dataIDI *dataIDI) Compare(listable traits.Listable) int {
	return dataIDI.Impl.(ids.DataID).Compare(listable)
}

func (dataIDI *dataIDI) String() string {
	return dataIDI.Impl.(ids.DataID).String()
}

func (dataIDI *dataIDI) Bytes() []byte {
	return dataIDI.Impl.(ids.DataID).Bytes()
}

func (dataIDI *dataIDI) GetHashID() ids.HashID {
	return dataIDI.Impl.(ids.DataID).GetHashID()
}

func (dataIDI *dataIDI) IsDataID() {
}
