package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type idI ids.ID

var _ ids2.ID = (*idI)(nil)

func (idI *idI) Compare(listable traits.Listable) int {
	return idI.Impl.(ids2.ID).Compare(listable)
}
func (idI *idI) String() string {
	return idI.Impl.(ids2.ID).String()
}
func (idI *idI) Bytes() []byte {
	return idI.Impl.(ids2.ID).Bytes()
}
