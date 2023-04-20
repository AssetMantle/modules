// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"math"
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func GenerateRandomData(r *rand.Rand) data.Data {
	randomPositiveInt := int(math.Abs(float64(r.Int())))

	switch randomPositiveInt % 5 {
	case 0:
		return baseData.NewIDData(GenerateRandomID(r))
	case 1:
		return baseData.NewStringData(simulationTypes.RandStringOfLength(r, r.Intn(99)))
	case 2:
		return baseData.NewDecData(simulationTypes.RandomDecAmount(r, sdkTypes.NewDec(99)))
	case 3:
		return baseData.NewHeightData(baseTypes.NewHeight(r.Int63()))
	case 4:
		return GenerateRandomListData(r)
	default:
		return nil
	}
}

func GenerateRandomListData(r *rand.Rand) data.ListData {
	listData := baseData.NewListData(baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID(sdkTypes.DefaultBondDenom)).ToAnyID()))

	for i := 0; i < int(math.Abs(float64(r.Int()))); i++ {
		listData.Add(baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(99))))))
	}

	return listData
}
