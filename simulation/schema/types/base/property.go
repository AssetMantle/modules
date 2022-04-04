// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"

	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

func GenerateRandomProperty(r *rand.Rand) types.Property {
	return base.NewProperty(GenerateRandomID(r), GenerateRandomData(r))
}
