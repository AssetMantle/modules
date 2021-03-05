/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"math"
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GenerateRandomData(r *rand.Rand) types.Data {
	randomPositiveInt := int(math.Abs(float64(r.Int())))

	switch randomPositiveInt % 4 {
	case 0:
		return base.NewIDData(GenerateRandomID(r))
	case 1:
		return base.NewStringData(simulation.RandStringOfLength(r, r.Intn(99)))
	case 2:
		return base.NewDecData(simulation.RandomDecAmount(r, sdkTypes.NewDec(99)))
	case 3:
		return base.NewHeightData(base.NewHeight(r.Int63()))
	default:
		return nil
	}
}
