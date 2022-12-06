package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.ID = (*IDI)(nil)

func (idI *IDI) Compare(listable traits.Listable) int {
	return idI.Impl.(ids2.ID).Compare(listable)
}
func (idI *IDI) Bytes() []byte {
	return idI.Impl.(ids2.ID).Bytes()
}
