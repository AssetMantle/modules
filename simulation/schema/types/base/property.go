package base

import (
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomProperty(r *rand.Rand) base.Property {
	return *base.NewProperty(GenerateRandomID(r), GenerateRandomFact(r))
}
