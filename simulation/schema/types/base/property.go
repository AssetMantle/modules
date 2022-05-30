// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"

	"github.com/AssetMantle/modules/schema/properties"
)

func GenerateRandomProperty(r *rand.Rand) properties.Property {
	return baseTypes.NewProperty(GenerateRandomID(r), GenerateRandomData(r))
}
