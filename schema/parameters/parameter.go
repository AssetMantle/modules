// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
)

type Parameter interface {
	String() string

	Equal(Parameter) bool
	Validate() error

	GetID() ids.ID
	GetData() data.Data
	GetValidator() func(interface{}) error

	Mutate(data.Data) Parameter
}
