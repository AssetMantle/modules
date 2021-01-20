/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Transactional interface {
	GetValue() sdkTypes.Dec
	Send(sdkTypes.Dec) Transactional
	Receive(sdkTypes.Dec) Transactional

	CanSend(sdkTypes.Dec) bool
}
