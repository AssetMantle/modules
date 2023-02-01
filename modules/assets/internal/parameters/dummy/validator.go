// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

func validator(i interface{}) error {
	switch value := i.(type) {
	case helpers.Parameter:
		if value.GetMetaProperty().GetID().Compare(ID) != 0 || value.GetMetaProperty().GetData().Get().(data.DecData).Get().IsNegative() {
			return constants.InvalidParameter
		}

		return nil
	case data.DecData:
		if value.Get().IsNegative() {
			return constants.InvalidParameter
		}

		return nil
	default:
		return constants.IncorrectFormat
	}
}
