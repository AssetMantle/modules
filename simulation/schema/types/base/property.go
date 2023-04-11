// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"

	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/AssetMantle/schema/x/properties"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func GenerateRandomProperty(r *rand.Rand) properties.Property {
	return baseProperties.NewMesaProperty(baseIDs.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(99))), GenerateRandomData(r))
}
