// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math"
	"math/rand"
	"strings"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func GenerateRandomData(r *rand.Rand) data.Data {
	randomPositiveInt := int(math.Abs(float64(r.Int())))

	switch randomPositiveInt % 4 {
	case 0:
		return baseData.NewIDData(GenerateRandomID(r))
	case 1:
		return baseData.NewStringData(simulationTypes.RandStringOfLength(r, r.Intn(99)))
	case 2:
		return baseData.NewDecData(simulationTypes.RandomDecAmount(r, sdkTypes.NewDec(99)))
	case 3:
		return baseData.NewHeightData(baseTypes.NewHeight(r.Int63()))
	default:
		return nil
	}
}

func GenerateRandomCoinListString(r *rand.Rand) string {
	prefix := "I|COI|"
	list := ""

	for i := 0; i < int(math.Abs(float64(r.Int()))); i++ {
		list += prefix + simulationTypes.RandStringOfLength(r, r.Intn(99)) + ","
	}

	return strings.TrimSuffix(list, ",")
}
