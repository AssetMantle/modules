// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"github.com/AssetMantle/modules/schema/helpers"
)

type auxiliaryResponse struct {
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func newAuxiliaryResponse() helpers.AuxiliaryResponse {
	return auxiliaryResponse{}
}
