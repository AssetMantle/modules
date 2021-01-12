/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Splittable interface {
	GetSplit() sdkTypes.Dec
}
