package base

import (
	"math"
	"math/rand"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func TestGenerateRandomData(t *testing.T) {
	r := rand.New(rand.NewSource(7))
	randomPositiveInt := int(math.Abs(float64(r.Int())))

	switch randomPositiveInt % 4 {
	case 1:
		require.Equal(t, GenerateRandomData(r), base.NewStringData(simulation.RandStringOfLength(r, r.Intn(99))))
	case 2:
		require.Equal(t, GenerateRandomData(r), base.NewDecData(simulation.RandomDecAmount(r, sdkTypes.NewDec(99))))
	case 3:
		require.Equal(t, GenerateRandomData(r), base.NewHeightData(base.NewHeight(r.Int63())))
	}

}
