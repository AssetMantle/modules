// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

var ID = base.NewID("dummy")
var DefaultData = base.NewDecData(sdkTypes.SmallestDec())
