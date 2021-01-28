package base

import (
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomImmutables(r *rand.Rand) types.Immutables {
	return base.NewImmutables(GenerateRandomProperties(r))
}
