// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

/*
 *  Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
 *  SPDX-License-Identifier: Apache-2.0
 */

package qualified

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

// TODO rename to Immutables
type HasImmutables interface {
	// GetImmutableProperties return the immutable properties object
	// does not return nil
	GetImmutableProperties() types.Properties

	GenerateHashID() types.ID
}
