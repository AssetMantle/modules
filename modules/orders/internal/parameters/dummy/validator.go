// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/types"
)

func validator(i interface{}) error {
	switch value := i.(type) {
	case types.Parameter:
		data, err := value.GetData().AsDec()
		if err != nil || value.GetID().Compare(ID) != 0 || data.IsNegative() {
			return errors.InvalidParameter
		}

		return nil
	case types.Data:
		data, err := value.AsDec()
		if err != nil || data.IsNegative() {
			return errors.InvalidParameter
		}

		return nil
	default:
		return errors.IncorrectFormat
	}
}
