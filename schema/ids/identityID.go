package ids

import "github.com/AssetMantle/modules/schema/ids/base"

type IdentityID interface {
	ID
	GetHashID() *base.HashID
	IsIdentityID()
	IDString() string
}
