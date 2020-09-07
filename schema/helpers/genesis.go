/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Genesis interface {
	Default() Genesis
	Validate() error
	Initialize(sdkTypes.Context, Mapper)
	Export(sdkTypes.Context, Mapper) Genesis
	RegisterCodec(*codec.Codec)
	Marshall() []byte
	Unmarshall([]byte) Genesis
}
