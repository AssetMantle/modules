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

type HasMutables interface {
	// GetMutableProperties return the mutable properties object
	// does not return nil
	GetMutableProperties() types.Properties

	Mutate(propertyList ...types.Property) HasMutables
}
