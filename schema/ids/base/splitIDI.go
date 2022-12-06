package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type splitIDI ids.SplitID

var _ ids2.SplitID = (*splitIDI)(nil)

func (splitIDI *splitIDI) Compare(listable traits.Listable) int {
	return splitIDI.Impl.(ids2.SplitID).Compare(listable)
}

func (splitIDI *splitIDI) String() string {
	return splitIDI.Impl.(ids2.SplitID).String()
}

func (splitIDI *splitIDI) Bytes() []byte {
	return splitIDI.Impl.(ids2.SplitID).Bytes()
}

func (splitIDI *splitIDI) GetOwnableID() ids2.ID {
	return splitIDI.Impl.(ids2.SplitID).GetOwnableID()
}

func (splitIDI *splitIDI) IsSplitID() {
}
