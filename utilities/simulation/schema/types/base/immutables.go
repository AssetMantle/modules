package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"math/rand"
)

func GenerateRandomImmutables(r *rand.Rand) types.Immutables {

	return base.NewImmutables(GenerateRandomProperties(r))
}
