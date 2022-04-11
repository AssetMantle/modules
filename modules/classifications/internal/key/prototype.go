// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Prototype() helpers.Key {
	if classificationID, err := classificationIDFromInterface(baseIDs.NewID("")); err != nil {
		panic(classificationID)
	} else {
		return classificationID
	}
}
