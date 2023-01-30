// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/AssetMantle/modules/schema/properties"
)

type Parameter interface {
	Validate() error
	GetMetaProperty() properties.MetaProperty
	Mutate(properties.MetaProperty) Parameter
}
