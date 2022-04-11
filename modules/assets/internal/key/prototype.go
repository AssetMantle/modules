// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Prototype() helpers.Key {
	if assetID, err := assetIDFromInterface(baseIDs.NewID("")); err != nil {
		panic(err)
	} else {
		return assetID
	}
}
