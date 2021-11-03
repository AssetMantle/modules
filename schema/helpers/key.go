/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Key interface {
	GenerateStoreKeyBytes() []byte
	IsPartial() bool
	Equals(Key) bool
	types.Proto
}
