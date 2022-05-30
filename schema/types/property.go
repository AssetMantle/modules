// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/ids"
)

// TODO add update method
type Property interface {
	GetID() ids.PropertyID
	GetDataID() ids.DataID
	GetKey() ID
	GetType() ID
	GetHash() ID

	capabilities.Listable
}
