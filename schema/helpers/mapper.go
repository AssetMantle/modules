/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Mapper interface {
	NewCollection(sdkTypes.Context) Collection
	GetCodec() codec.BinaryMarshaler
	MarshalMappable(Mappable) ([]byte, error)
	UnmarshalMappable(bz []byte) (Mappable, error)
	MarshalParameter(types.Parameter) ([]byte, error)
	UnmarshalParameter(bz []byte) (types.Parameter, error)

	Create(sdkTypes.Context, Mappable)
	Read(sdkTypes.Context, Key) Mappable
	Update(sdkTypes.Context, Mappable)
	Delete(sdkTypes.Context, Key)
	Iterate(sdkTypes.Context, Key, func(Mappable) bool)
	ReverseIterate(sdkTypes.Context, Key, func(Mappable) bool)

	GetParameter(sdkTypes.Context, types.ID) types.Parameter
	GetAllParameter(sdkTypes.Context) []types.Parameter
	SetParameter(sdkTypes.Context, types.Parameter) error

	StoreDecoder(kv.Pair, kv.Pair) string

	Initialize(storeKey sdkTypes.StoreKey, codec codec.BinaryMarshaler, paramSubspace paramTypes.Subspace, parameters []types.Parameter) Mapper
}
