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

	Upsert(context.Context, Record)
	Read(context.Context, Key) Record
	Delete(context.Context, Key)
	FetchAll(context.Context) []Record
	Iterate(context.Context, Key, func(Record) bool)
	IterateAll(context.Context, func(Record) bool)
	IteratePaginated(context.Context, Key, int32, func(Record) bool)

	StoreDecoder(kv.Pair, kv.Pair) string
	Initialize(*sdkTypes.KVStoreKey) Mapper
}
