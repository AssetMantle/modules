package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type identityIDI ids.IdentityID

var _ ids2.IdentityID = (*identityIDI)(nil)

func (identityIDI *identityIDI) Compare(listable traits.Listable) int {
	return identityIDI.Impl.(ids2.IdentityID).Compare(listable)
}

func (identityIDI *identityIDI) String() string {
	return identityIDI.Impl.(ids2.IdentityID).String()
}

func (identityIDI *identityIDI) Bytes() []byte {
	return identityIDI.Impl.(ids2.IdentityID).Bytes()
}

func (identityIDI *identityIDI) GetHashID() HashID {
	return identityIDI.Impl.(ids2.IdentityID).GetHashID()
}

func (identityIDI *identityIDI) IsIdentityID() {
}
