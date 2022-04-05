// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"

	"math"
	"math/rand"
)

func GenerateRandomData(r *rand.Rand) types.Data {
	randomPositiveInt := int(math.Abs(float64(r.Int())))

	switch randomPositiveInt % 4 {
	case 0:
		return baseData.NewIDData(GenerateRandomID(r))
	case 1:
		return baseData.NewStringData(simulation.RandStringOfLength(r, r.Intn(99)))
	case 2:
		return baseData.NewDecData(simulation.RandomDecAmount(r, sdkTypes.NewDec(99)))
	case 3:
		return baseData.NewHeightData(base.NewHeight(r.Int63()))
	default:
		return nil
	}
}
