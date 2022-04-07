// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math"
	"math/rand"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func TestGenerateRandomData(t *testing.T) {
	r := rand.New(rand.NewSource(7))
	randomPositiveInt := int(math.Abs(float64(r.Int())))

	switch randomPositiveInt % 4 {
	case 1:
		require.Equal(t, GenerateRandomData(r), baseData.NewStringData(simulation.RandStringOfLength(r, r.Intn(99))))
	case 2:
		require.Equal(t, GenerateRandomData(r), baseData.NewDecData(simulation.RandomDecAmount(r, sdkTypes.NewDec(99))))
	case 3:
		require.Equal(t, GenerateRandomData(r), baseData.NewHeightData(baseTypes.NewHeight(r.Int63())))
	}
}
