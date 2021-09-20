package base

import (
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomFact(r *rand.Rand) types.Fact {
	return base.NewFact(GenerateRandomData(r))
}
