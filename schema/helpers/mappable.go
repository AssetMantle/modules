// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

type Mappable interface {
	GetKey() Key
	RegisterCodec(*codec.Codec)
}
