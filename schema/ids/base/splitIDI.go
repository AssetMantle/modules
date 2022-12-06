package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.SplitID = (*SplitIDI)(nil)

func (splitIDI *SplitIDI) Compare(listable traits.Listable) int {
	return splitIDI.Impl.(ids2.SplitID).Compare(listable)
}

func (splitIDI *SplitIDI) Bytes() []byte {
	return splitIDI.Impl.(ids2.SplitID).Bytes()
}

func (splitIDI *SplitIDI) GetOwnableID() ids2.ID {
	return splitIDI.Impl.(ids2.SplitID).GetOwnableID()
}

func (splitIDI *SplitIDI) IsSplitID() {
}
