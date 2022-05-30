// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package capabilities

import "github.com/AssetMantle/modules/schema/types"

type Lockable interface {
	GetLock() types.Property
}
