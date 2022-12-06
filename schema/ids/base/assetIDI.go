package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type assetIDI base.AssetIDI

var _ ids2.AssetID = (*assetIDI)(nil)

func (assetID *assetIDI) Compare(listable traits.Listable) int {
	return assetID.Impl.(ids2.AssetID).Compare(listable)
}

func (assetID *assetIDI) String() string {
	return assetID.Impl.(ids2.AssetID).String()
}

func (assetID *assetIDI) Bytes() []byte {
	return assetID.Impl.(ids2.AssetID).Bytes()
}

func (assetID *assetIDI) IsOwnableID() {
}

func (assetID *assetIDI) IsAssetID() {
}
