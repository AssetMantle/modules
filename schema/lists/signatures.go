// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/types"
)

// TODO move to list
type Signatures interface {
	Get(ids.ID) types.Signature

	GetList() []types.Signature

	Add(types.Signature) Signatures
	Remove(types.Signature) Signatures
	Mutate(types.Signature) Signatures
}
