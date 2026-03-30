package distribute

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/utilities/name"
)

type dummy struct{}

var Auxiliary = baseHelpers.NewAuxiliary(
	name.GetPackageName(dummy{}),
	keeperPrototype,
)
