package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.MaintainerID = (*MaintainerIDI)(nil)

func (maintainerIDI *MaintainerIDI) Compare(listable traits.Listable) int {
	return maintainerIDI.Impl.(ids2.MaintainerID).Compare(listable)
}

func (maintainerIDI *MaintainerIDI) Bytes() []byte {
	return maintainerIDI.Impl.(ids2.MaintainerID).Bytes()
}
func (maintainerIDI *MaintainerIDI) IsMaintainerID() {
}
