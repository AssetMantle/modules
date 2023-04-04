package invariants

import (
	"github.com/AssetMantle/modules/modules/splits/internal/invariants/dummy"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/helpers/base"
)

func Prototype() helpers.Invariants {
	return base.NewInvariants(module.Name, "", dummy.Invariant)
}
