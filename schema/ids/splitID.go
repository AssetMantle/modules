package ids

import "github.com/AssetMantle/modules/schema/ids/base"

type SplitID interface {
	ID
	GetOwnableID() *base.OwnableID
	IsSplitID()
	SplitIDString() string
}
