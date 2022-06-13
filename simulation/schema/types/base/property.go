// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"

	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func GenerateRandomProperty(r *rand.Rand) properties.Property {
	return baseProperties.NewProperty(GenerateRandomID(r), GenerateRandomData(r))
}
