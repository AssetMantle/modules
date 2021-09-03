package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomProperty(r *rand.Rand) types.Property {
	return base.NewProperty(GenerateRandomID(r), GenerateRandomFact(r))
}
