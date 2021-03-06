/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterLegacyCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*HasImmutables)(nil), nil)
	codec.RegisterInterface((*HasMutables)(nil), nil)
}
