// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func GenerateRandomProperty(r *rand.Rand) properties.Property {
	return baseProperties.NewMesaProperty(baseIDs.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(99))), GenerateRandomData(r))
}
