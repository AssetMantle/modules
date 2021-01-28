package base

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomID(r *rand.Rand) types.ID {

	return base.NewID(simulation.RandStringOfLength(r, r.Int()))
}
