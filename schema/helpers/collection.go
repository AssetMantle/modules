/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

// Collection a list of mappable with create CRUD methods
type Collection interface {
	GetKey() Key
	Get(Key) Mappable
	GetList() []Mappable

	Iterate(Key, func(Mappable) bool)
	Fetch(Key) Collection
	Add(Mappable) Collection
	Remove(Mappable) Collection
	Mutate(Mappable) Collection
	Initialize(sdkTypes.Context, Mapper) Collection
}
