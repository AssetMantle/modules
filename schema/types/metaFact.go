/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/99designs/keyring"

type MetaFact interface {
	GetHashID() ID
	GetTypeID() ID
	GetData() Data
	GetSignatures() Signatures

	Sign(keyring.Keyring) MetaFact
	RemoveData() Fact
}
