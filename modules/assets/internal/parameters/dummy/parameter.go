// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
)

var Parameter = baseTypes.NewParameter(ID, DefaultData, validator)
