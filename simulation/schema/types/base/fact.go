package base

import (
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomFact(r *rand.Rand) base.Fact {
	return *base.NewFact(GenerateRandomData(r))
}
