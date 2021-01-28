/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GetOwnableTotalSplitsValue(collection helpers.Collection, ownableID types.ID) sdkTypes.Dec {
	value := sdkTypes.ZeroDec()
	accumulator := func(mappable helpers.Mappable) bool {
		if key.ReadOwnableID(key.ToID(mappable.GetKey())).Equals(ownableID) {
			value = value.Add(mappable.(mappables.Split).GetValue())
		}

		return false
	}
	collection.Iterate(key.FromID(base.NewID("")), accumulator)

	return value
}
