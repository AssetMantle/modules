// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

// TODO move base to own package
// TODO rename to Mutables
type HasMutables struct {
	Properties types.Properties `json:"properties"`
}

var _ qualified.HasMutables = (*HasMutables)(nil)

func (mutables HasMutables) GetMutableProperties() types.Properties {
	if mutables.Properties == nil {
		return base.NewProperties()
	}

	return mutables.Properties
}
func (mutables HasMutables) Mutate(propertyList ...types.Property) qualified.HasMutables {
	for _, property := range propertyList {
		mutables.Properties = mutables.Properties.Mutate(property)
	}

	return mutables
}
