package compliance

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/utilities/name"
)

var Auxiliary = baseHelpers.NewAuxiliary(
	name.GetPackageName(dummy{}),
	keeperPrototype,
)

type dummy struct{}
