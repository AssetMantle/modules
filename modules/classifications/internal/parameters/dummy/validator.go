/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package dummy

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func validator(i interface{}) error {
	switch value := i.(type) {
	case types.Parameter:
		data, Error := value.GetData().AsDec()
		if Error != nil || !value.GetID().Equals(ID) || data.IsNegative() {
			return errors.InvalidParameter
		}

		return nil
	case types.Data:
		data, Error := value.AsDec()
		if Error != nil || data.IsNegative() {
			return errors.InvalidParameter
		}

		return nil
	default:
		return errors.IncorrectFormat
	}
}
