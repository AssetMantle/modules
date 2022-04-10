// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type Mutables struct {
	Properties types.Properties `json:"properties"`
}

var _ qualified.Mutables = (*Mutables)(nil)

func (mutables Mutables) GetMutableProperties() types.Properties {
	if mutables.Properties == nil {
		return baseTypes.NewProperties()
	}

	return mutables.Properties
}
func (mutables Mutables) Mutate(propertyList ...types.Property) qualified.Mutables {
	for _, property := range propertyList {
		mutables.Properties = mutables.Properties.Mutate(property)
	}

	return mutables
}
