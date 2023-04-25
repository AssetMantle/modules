// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/schema/go/ids/base"
	"math"
	"math/rand"
	"time"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func GenerateRandomData(r *rand.Rand, randomPositiveInt int) data.Data {
	switch randomPositiveInt % 8 {
	case 0:
		return baseData.NewIDData(GenerateRandomID(r))
	case 1:
		return baseData.NewStringData(simulationTypes.RandStringOfLength(r, r.Intn(99)))
	case 2:
		return baseData.NewDecData(simulationTypes.RandomDecAmount(r, sdkTypes.NewDec(99)))
	case 3:
		return baseData.NewHeightData(baseTypes.NewHeight(r.Int63()))
	case 4:
		return baseData.NewBooleanData(random.GenerateRandomBool())
	case 5:
		return baseData.NewAccAddressData(simulationTypes.RandomAccounts(r, 1)[0].Address)
	case 6:
		return baseData.NewNumberData(int64(r.Intn(99)))
	case 7:
		return GenerateRandomListData(r)
	default:
		return nil
	}
}

func GenerateRandomCoinListString(listCount int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := baseData.NewListData(baseData.NewIDData(base.NewCoinID(base.NewStringID("stake")).ToAnyID()))

	for i := 0; i < listCount; i++ {
		list.Add(baseData.NewIDData(base.NewCoinID(base.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(127)))).ToAnyID()))
	}
	return list.AsString()
}

func GenerateRandomListData(r *rand.Rand) data.ListData {
	listDataType := int(math.Abs(float64(r.Int()))) % 7
	listData := baseData.PrototypeListData()

	for i := 0; i < r.Intn(10); i++ {
		listData.Add(GenerateRandomData(r, listDataType))
	}

	return listData
}
