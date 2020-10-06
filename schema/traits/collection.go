/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Collection interface {
	GetKey() helpers.Key
	Get(types.ID) Mappable
	GetList() []Mappable

	Fetch(types.ID) Collection
	Add(mappables.InterNFT) Collection
	Remove(mappables.InterNFT) Collection
	Mutate(mappables.InterNFT) Collection
	Initialize(sdkTypes.Context, helpers.Mapper) Collection
	RegisterCodec(*codec.Codec)
}
