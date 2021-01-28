package base

import (
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomMutables(r *rand.Rand) types.Mutables {
	return base.NewMutables(GenerateRandomProperties(r))
}
