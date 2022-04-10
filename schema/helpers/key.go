// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import "github.com/cosmos/cosmos-sdk/codec"

type Key interface {
	GenerateStoreKeyBytes() []byte
	RegisterCodec(*codec.Codec)
	IsPartial() bool
	Equals(Key) bool
}
