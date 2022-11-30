// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import "github.com/AssetMantle/modules/schema/ids/base"

type PropertyID interface {
	GetKey() *base.StringID
	GetType() *base.StringID
	ID
	IsPropertyID()
}
