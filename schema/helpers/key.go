// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import "github.com/cosmos/cosmos-sdk/codec"

// Key SHOULD be derivable from the object it is referencing and SHOULD not be totally arbitrary or sequential
type Key interface {
	String() string
	GenerateStoreKeyBytes() []byte
	// TODO Check is register legacyAmino is still required
	RegisterLegacyAminoCodec(*codec.LegacyAmino)
	IsPartial() bool
	// TODO ** check all key impls
	Equals(Key) bool
}
