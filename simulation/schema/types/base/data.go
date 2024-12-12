// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math"
	"math/rand"
	"time"

	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	constantsData "github.com/AssetMantle/schema/data/constants"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/ids/constants"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/AssetMantle/modules/utilities/random"
)

func GenerateRandomDataForTypeID(r *rand.Rand, id ids.StringID) data.Data {
	switch id {
	case base.NewStringID(constants.AnyIDType):
		return GenerateRandomData(r, 0)
	case constantsData.StringDataTypeID:
		return GenerateRandomData(r, 1)
	case constantsData.DecDataTypeID:
		return GenerateRandomData(r, 2)
	case constantsData.HeightDataTypeID:
		return GenerateRandomData(r, 3)
	case constantsData.BooleanDataTypeID:
		return GenerateRandomData(r, 4)
	case constantsData.AccAddressDataTypeID:
		return GenerateRandomData(r, 5)
	case constantsData.NumberDataTypeID:
		return GenerateRandomData(r, 6)
	case constantsData.ListDataTypeID:
		return GenerateRandomData(r, 7)
	}

	return nil
}

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
		return baseData.NewNumberData(sdkTypes.NewInt(int64(r.Intn(99))))
	case 7:
		return GenerateRandomListData(r)
	default:
		return nil
	}
}

func GenerateRandomPrototypeData(r *rand.Rand, randomPositiveInt int) data.Data {
	switch randomPositiveInt % 8 {
	case 0:
		return baseData.PrototypeIDData()
	case 1:
		return baseData.PrototypeStringData()
	case 2:
		return baseData.PrototypeDecData()
	case 3:
		return baseData.PrototypeHeightData()
	case 4:
		return baseData.PrototypeBooleanData()
	case 5:
		return baseData.PrototypeAccAddressData()
	case 6:
		return baseData.PrototypeNumberData()
	case 7:
		return baseData.PrototypeListData()
	default:
		return nil
	}
}

func GenerateRandomCoinListString(listCount int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := baseData.NewListData(baseData.NewIDData(base.NewStringID("stake").ToAnyID()))

	for i := 0; i < listCount; i++ {
		list.Add(baseData.NewIDData(base.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(127)))))
	}
	return list.AsString()
}

func GenerateRandomListData(r *rand.Rand) data.ListData {
	listDataType := int(math.Abs(float64(r.Int()))) % 7
	listData := baseData.PrototypeListData()

	for i := 0; i < r.Intn(10); i++ {
		listData.Add(GenerateRandomData(r, listDataType).(data.ListableData))
	}

	return listData
}
