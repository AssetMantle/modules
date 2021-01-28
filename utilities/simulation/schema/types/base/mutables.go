package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"math/rand"
)

func GenerateRandomMutables(r *rand.Rand) types.Mutables {
	return base.NewMutables(GenerateRandomProperties(r))
}
