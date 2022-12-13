// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

/*
 *  Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
 *  SPDX-License-Identifier: Apache-2.0
 */

package qualified

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
)

type Mutables interface {
	// GetMutablePropertyList return the mutable properties object
	// does not return nil
	GetMutablePropertyList() lists.PropertyList

	Mutate(propertyList ...properties.Property) Mutables
}
