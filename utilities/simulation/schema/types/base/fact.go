package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"math"
	"math/rand"
)

func GenerateRandomFact(r *rand.Rand) types.Fact {
	randomPositiveInt := int(math.Abs(float64(r.Int())))
	var data types.Data

	switch randomPositiveInt % 4 {
	case 0:
		data = base.NewIDData(GenerateRandomID(r))
	case 1:
		data = base.NewStringData(simulation.RandStringOfLength(r, r.Int()))
	case 2:
		data = base.NewDecData(simulation.RandomDecAmount(r, sdkTypes.NewDec(9999999999)))
	case 3:
		data = base.NewHeightData(base.NewHeight(r.Int63()))
	default:
		return nil
	}

	return base.NewFact(data)
}
