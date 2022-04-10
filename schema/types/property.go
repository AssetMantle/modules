// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/AssetMantle/modules/schema/traits"
)

type Property interface {
	GetID() ID
	GetDataID() ID
	GetKey() ID
	GetType() ID
	GetHash() ID
	traits.Listable
}
