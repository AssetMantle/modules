package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type stringIDI ids.StringID

var _ ids2.StringID = (*stringIDI)(nil)

func (stringIDI *stringIDI) Compare(listable traits.Listable) int {
	return stringIDI.Impl.(ids2.StringID).Compare(listable)
}

func (stringIDI *stringIDI) String() string {
	return stringIDI.Impl.(ids2.StringID).String()
}

func (stringIDI *stringIDI) Bytes() []byte {
	return stringIDI.Impl.(ids2.StringID).Bytes()
}

func (stringIDI *stringIDI) IsStringID() {
}
