package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.HashID = (*HashIDI)(nil)

func (hashIDI *HashIDI) Compare(listable traits.Listable) int {
	return hashIDI.Impl.(ids2.HashID).Compare(listable)
}

func (hashIDI *HashIDI) Bytes() []byte {
	return hashIDI.Impl.(ids2.HashID).Bytes()
}

func (hashIDI *HashIDI) IsHashID() {
}
