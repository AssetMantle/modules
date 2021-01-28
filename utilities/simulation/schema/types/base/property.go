package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"math/rand"
)

func GenerateRandomProperty(r *rand.Rand) types.Property {
	return base.NewProperty(GenerateRandomID(r), GenerateRandomFact(r))
}
