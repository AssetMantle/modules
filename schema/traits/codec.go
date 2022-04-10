// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package traits

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/qualified"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*qualified.HasImmutables)(nil), nil)
	codec.RegisterInterface((*qualified.HasMutables)(nil), nil)
}
