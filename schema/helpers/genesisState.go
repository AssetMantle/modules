/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState interface {
	Default() GenesisState
	Validate(sdkTypes.Context) error
	Initialize(sdkTypes.Context, Mapper)
	Export(sdkTypes.Context, Mapper) GenesisState
	RegisterCodec(*codec.Codec)
	Marshall() []byte
	Unmarshall([]byte) GenesisState
}
