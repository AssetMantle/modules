/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	"github.com/tendermint/tendermint/crypto"
)

type Signature interface {
	String() string
	Bytes() []byte

	GetID() ID

	Verify(crypto.PubKey, []byte) bool
	GetValidityHeight() Height
	HasExpired(Height) bool
}
