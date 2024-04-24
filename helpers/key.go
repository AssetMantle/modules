// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

// Key SHOULD be derivable from the object it is referencing and SHOULD not be totally arbitrary or sequential
type Key interface {
	ValidateBasic() error
	GenerateStorePrefixBytes() []byte
	GenerateStoreKeyBytes() []byte
	GeneratePrefixedStoreKeyBytes() []byte
	IsPartial() bool
	// TODO ** check all key impls
	Equals(Key) bool
}
