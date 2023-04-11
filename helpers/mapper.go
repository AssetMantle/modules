// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

type Mapper interface {
	NewCollection(context.Context) Collection
	Create(context.Context, Mappable)
	Read(context.Context, Key) Mappable
	Update(context.Context, Mappable)
	Delete(context.Context, Key)
	IterateAll(context.Context, func(Mappable) bool)
	Iterate(context.Context, Key, func(Mappable) bool)
	ReverseIterate(context.Context, Key, func(Mappable) bool)

	StoreDecoder(kv.Pair, kv.Pair) string

	Initialize(*sdkTypes.KVStoreKey) Mapper
}
