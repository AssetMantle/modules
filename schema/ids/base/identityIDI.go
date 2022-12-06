package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.IdentityID = (*IdentityIDI)(nil)

func (identityIDI *IdentityIDI) Compare(listable traits.Listable) int {
	return identityIDI.Impl.(ids2.IdentityID).Compare(listable)
}
func (identityIDI *IdentityIDI) GetHashID() ids2.HashID {
	return identityIDI.Impl.(ids2.IdentityID).GetHashID()
}
func (identityIDI *IdentityIDI) Bytes() []byte {
	return identityIDI.Impl.(ids2.IdentityID).Bytes()
}

func (identityIDI *IdentityIDI) IsIdentityID() {
}
