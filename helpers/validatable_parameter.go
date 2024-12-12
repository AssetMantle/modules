// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/parameters"
)

type ValidatableParameter interface {
	GetParameter() parameters.Parameter
	Mutate(data.Data) ValidatableParameter
	Validate() error
}
