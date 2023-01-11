// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
)

type Parameter interface {
	AsString() string

	Equal(Parameter) bool
	Validate() error

	GetID() ids.ID
	GetValidator() func(interface{}) error
	GetData() data.AnyData

	Mutate(data.Data) Parameter
}
