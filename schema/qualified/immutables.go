// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

/*
 *  Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
 *  SPDX-License-Identifier: Apache-2.0
 */

package qualified

import (
	"github.com/AssetMantle/modules/schema/types"
)

type Immutables interface {
	// GetImmutableProperties return the immutable properties object
	// does not return nil
	GetImmutableProperties() types.Properties

	GenerateHashID() types.ID
}
