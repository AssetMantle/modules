package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.AssetID = (*AssetIDI)(nil)

func (assetID *AssetIDI) Compare(listable traits.Listable) int {
	return assetID.Impl.(ids2.AssetID).Compare(listable)
}

func (assetID *AssetIDI) Bytes() []byte {
	return assetID.Impl.(ids2.AssetID).Bytes()
}

func (assetID *AssetIDI) IsOwnableID() {
}

func (assetID *AssetIDI) IsAssetID() {
}
