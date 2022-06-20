// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/parameters"
)

func validator(i interface{}) error {
	switch value := i.(type) {
	case parameters.Parameter:
		if value.GetID().Compare(ID) != 0 || value.GetData().(data.DecData).Get().IsNegative() {
			return constants.InvalidParameter
		}

		return nil
	case data.Data:
		if value.(data.DecData).Get().IsNegative() {
			return constants.InvalidParameter
		}

		return nil
	default:
		return constants.IncorrectFormat
	}
}
