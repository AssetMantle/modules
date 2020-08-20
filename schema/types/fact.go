/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	"github.com/99designs/keyring"
)

type Fact interface {
	GetHash() string
	GetSignatures() Signatures

	Sign(keyring.Keyring) Fact
}
