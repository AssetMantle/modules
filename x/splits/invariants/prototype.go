package invariants

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/invariants/dummy"
)

func Prototype() helpers.Invariants {
	return base.NewInvariants(constants.ModuleName, "", dummy.Invariant)
}
