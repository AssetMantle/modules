package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.OwnableID = (*OwnableIDI)(nil)

func (ownableIdI *OwnableIDI) Compare(listable traits.Listable) int {
	return ownableIdI.Impl.(ids2.OwnableID).Compare(listable)
}

func (ownableIdI *OwnableIDI) Bytes() []byte {
	return ownableIdI.Impl.(ids2.OwnableID).Bytes()
}

func (ownableIdI *OwnableIDI) IsOwnableID() {
}
