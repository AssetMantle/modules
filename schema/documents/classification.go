// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/cosmos/cosmos-sdk/types"
)

type Classification interface {
	GetBondAmount() types.Dec
	Document
}
