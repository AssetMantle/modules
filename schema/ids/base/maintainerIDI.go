package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type maintainerIDI ids.MaintainerID

var _ ids2.MaintainerID = (*maintainerIDI)(nil)

func (maintainerIDI *maintainerIDI) Compare(listable traits.Listable) int {
	return maintainerIDI.Impl.(ids2.MaintainerID).Compare(listable)
}

func (maintainerIDI *maintainerIDI) String() string {
	return maintainerIDI.Impl.(ids2.MaintainerID).String()
}

func (maintainerIDI *maintainerIDI) Bytes() []byte {
	return maintainerIDI.Impl.(ids2.MaintainerID).Bytes()
}
func (maintainerIDI *maintainerIDI) IsMaintainerID() {
}
