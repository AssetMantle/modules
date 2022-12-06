package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type ownableIDI ids.OwnableID

var _ ids2.OwnableID = (*ownableIDI)(nil)

func (ownableIdI *ownableIDI) Compare(listable traits.Listable) int {
	return ownableIdI.Impl.(ids2.OwnableID).Compare(listable)
}

func (ownableIdI *ownableIDI) String() string {
	return ownableIdI.Impl.(ids2.OwnableID).String()
}

func (ownableIdI *ownableIDI) Bytes() []byte {
	return ownableIdI.Impl.(ids2.OwnableID).Bytes()
}

func (ownableIdI *ownableIDI) IsOwnableID() {
}
