package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type hashIDI base.HashID

var _ ids2.HashID = (*hashIDI)(nil)

func (hashIDI *hashIDI) Compare(listable traits.Listable) int {
	return hashIDI.Impl.(ids2.HashID).Compare(listable)
}

func (hashIDI *hashIDI) String() string {
	return hashIDI.Impl.(ids2.HashID).String()
}

func (hashIDI *hashIDI) Bytes() []byte {
	return hashIDI.Impl.(ids2.HashID).Bytes()
}

func (hashIDI *hashIDI) IsHashID() {
}
