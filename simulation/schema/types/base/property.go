// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"math/rand"

	"github.com/AssetMantle/modules/utilities/random"
)

func GenerateRandomMesaProperty(r *rand.Rand) properties.Property {
	return baseProperties.NewMesaProperty(baseIDs.NewStringID(RandValidStringID(r)), GenerateRandomData(r, r.Intn(8)))
}
func GenerateRandomMetaProperty(r *rand.Rand) properties.Property {
	return baseProperties.NewMetaProperty(baseIDs.NewStringID(RandValidStringID(r)), GenerateRandomData(r, r.Intn(8)))
}
func GenerateRandomMetaPropertyWithoutData(r *rand.Rand) properties.Property {
	// Use MesaProperty (stores DataID hash only) instead of MetaProperty with prototype data.
	// MetaProperty with prototype data has nil internal AnyID.Impl which panics
	// when GenerateHashID calls GetDataID().
	return GenerateRandomMesaProperty(r)
}
func GenerateRandomProperty(r *rand.Rand) properties.Property {
	if random.GenerateRandomBool() {
		return GenerateRandomMesaProperty(r)
	}
	return GenerateRandomMetaProperty(r)
}
