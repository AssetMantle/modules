package ids

type AssetID interface {
	OwnableID
	IsAssetID()
}
