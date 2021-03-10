/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/kv"
)

type Mapper interface {
	NewCollection(sdkTypes.Context) Collection

	Create(sdkTypes.Context, Mappable)
	Read(sdkTypes.Context, Key) Mappable
	Update(sdkTypes.Context, Mappable)
	Delete(sdkTypes.Context, Key)
	Iterate(sdkTypes.Context, Key, func(Mappable) bool)
	ReverseIterate(sdkTypes.Context, Key, func(Mappable) bool)

	StoreDecoder(*codec.Codec, kv.Pair, kv.Pair) string

	Initialize(*sdkTypes.KVStoreKey) Mapper
}
