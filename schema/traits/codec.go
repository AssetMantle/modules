/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/schema/qualified"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*qualified.HasImmutables)(nil), nil)
	codec.RegisterInterface((*qualified.HasMutables)(nil), nil)
}
