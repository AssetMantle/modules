// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math"
	"math/rand"

	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
)

func GenerateRandomProperties(r *rand.Rand) lists.PropertyList {
	randomPositiveInt := int(math.Abs(float64(r.Int()))) % 11

	propertyList := make([]properties.Property, randomPositiveInt)

	for i := 0; i < randomPositiveInt; i++ {
		propertyList[i] = GenerateRandomProperty(r)
	}

	return baseLists.NewPropertyList(propertyList...)
}
