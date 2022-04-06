// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/types"
)

func validator(i interface{}) error {
	switch value := i.(type) {
	case types.Parameter:
		datum := value.GetData().(data.DecData).Get()
		if value.GetID().Compare(ID) != 0 || datum.IsNegative() {
			return errors.InvalidParameter
		}

		return nil
	case types.Data:
		datum := value.(data.DecData).Get()
		if datum.IsNegative() {
			return errors.InvalidParameter
		}

		return nil
	default:
		return errors.IncorrectFormat
	}
}

// TODO search and replace Datum
