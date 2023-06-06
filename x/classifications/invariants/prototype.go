package invariants

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/constants"
	"github.com/AssetMantle/modules/x/classifications/invariants/dummy"
)

func Prototype() helpers.Invariants {
	return base.NewInvariants(constants.ModuleName, "", dummy.Invariant)
}
