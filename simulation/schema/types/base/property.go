package base

import (
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomProperty(r *rand.Rand) types.Property {
	return base.NewProperty(GenerateRandomID(r), GenerateRandomFact(r))
}
