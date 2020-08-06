/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Burnable interface {
	CanBurn(types.Height) bool
	GetBurn() types.Height
}
