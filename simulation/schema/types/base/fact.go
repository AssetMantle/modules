package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomFact(r *rand.Rand) types.Fact {
	return base.NewFact(GenerateRandomData(r))
}
