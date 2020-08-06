/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Mapper interface {
	GetKVStoreKey() *sdkTypes.KVStoreKey

	Create(sdkTypes.Context, traits.Mappable)
	Read(sdkTypes.Context, types.ID) traits.Mappable
	Update(sdkTypes.Context, traits.Mappable)
	Delete(sdkTypes.Context, types.ID)
	Iterate(sdkTypes.Context, types.ID, func(traits.Mappable) bool)

	RegisterCodec(*codec.Codec)
}
