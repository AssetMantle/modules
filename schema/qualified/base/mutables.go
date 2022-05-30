// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/lists"
	baseTypes "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

type Mutables struct {
	Properties lists.PropertyList `json:"properties"`
}

var _ qualified.Mutables = (*Mutables)(nil)

func (mutables Mutables) GetMutablePropertyList() lists.PropertyList {
	if mutables.Properties == nil {
		return baseTypes.NewPropertyList()
	}

	return mutables.Properties
}
func (mutables Mutables) Mutate(propertyList ...properties.Property) qualified.Mutables {
	for _, property := range propertyList {
		mutables.Properties = mutables.Properties.Mutate(property)
	}

	return mutables
}
