package base

import (
	"math/rand"
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomID(r *rand.Rand) types.ID {
	return base.NewID(simulation.RandStringOfLength(r, r.Int()))
}

func GenerateRandomDecID(r *rand.Rand) types.ID {
	return base.NewDecID(sdkTypes.MustNewDecFromStr(strconv.FormatInt(r.Int63(), 10)))
}

func GenerateRandomHeightID(r *rand.Rand) types.ID {
	return base.NewHeightID(r.Int63())
}
