// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/lists"
)

type signatureList struct {
	lists.List
}

var _ lists.SignatureList = (*signatureList)(nil)
