// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package traits

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Transactional interface {
	// TODO change to getSupply
	GetValue() sdkTypes.Dec
	Send(sdkTypes.Dec) Transactional
	Receive(sdkTypes.Dec) Transactional

	CanSend(sdkTypes.Dec) bool
}
