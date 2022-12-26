package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

type invariants struct {
	moduleName    string
	route         string
	invariantList []sdkTypes.Invariant
}

var _ helpers.Invariants = (*invariants)(nil)

func (invariants invariants) Register(invariantRegistry sdkTypes.InvariantRegistry) {
	for _, invariant := range invariants.invariantList {
		// ****** TODO check if route should be different for each invariant
		invariantRegistry.RegisterRoute(invariants.moduleName, invariants.route, invariant)
	}
}

func NewInvariants(moduleName, route string, invariantList ...sdkTypes.Invariant) helpers.Invariants {
	return invariants{
		moduleName:    moduleName,
		route:         route,
		invariantList: invariantList,
	}
}
