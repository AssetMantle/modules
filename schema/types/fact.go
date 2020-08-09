/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/cosmos/cosmos-sdk/crypto/keyring"

type Fact interface {
	GetHash() string
	GetSignatures() Signatures
	IsMeta() bool

	Sign(keyring.Keyring) Fact
}
