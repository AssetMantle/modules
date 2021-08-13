/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	"github.com/99designs/keyring"
	"github.com/persistenceOne/persistenceSDK/schema"
)

type Fact interface {
	GetHashID() ID
	GetTypeID() ID
	GetSignatures() Signatures

	Sign(keyring.Keyring) Fact

	schema.Proto
}
