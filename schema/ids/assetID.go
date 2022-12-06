package ids

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	"github.com/AssetMantle/modules/schema/traits"
)

type AssetID interface {
	OwnableID
	IsAssetID()
}

type assetID ids.AssetID

func (assetID *assetID) Compare(listable traits.Listable) int {
	// TODO implement me

	panic("implement me")
}

func (assetID *assetID) String() string {
	// TODO implement me
	panic("implement me")
}

func (assetID *assetID) Bytes() []byte {
	// TODO implement me
	panic("implement me")
}

func (assetID *assetID) IsOwnableID() {
	// TODO implement me
	panic("implement me")
}

func (assetID *assetID) IsAssetID() {
	// TODO implement me
	panic("implement me")
}

var _ AssetID = (*assetID)(nil)
