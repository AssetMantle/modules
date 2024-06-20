// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math"
	"math/rand"

	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/AssetMantle/modules/utilities/random"
)

func GenerateRandomMesaProperty(r *rand.Rand) properties.Property {
	return baseProperties.NewMesaProperty(baseIDs.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(99))), GenerateRandomData(r, int(math.Abs(float64(r.Int())))))
}
func GenerateRandomMetaProperty(r *rand.Rand) properties.Property {
	return baseProperties.NewMetaProperty(baseIDs.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(99))), GenerateRandomData(r, int(math.Abs(float64(r.Intn(99))))))
}
func GenerateRandomMetaPropertyWithoutData(r *rand.Rand) properties.Property {
	return baseProperties.NewMetaProperty(baseIDs.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(99))), GenerateRandomPrototypeData(r, int(math.Abs(float64(r.Int())))))
}
func GenerateRandomProperty(r *rand.Rand) properties.Property {
	if random.GenerateRandomBool() {
		return GenerateRandomMesaProperty(r)
	}
	return GenerateRandomMetaProperty(r)
}
