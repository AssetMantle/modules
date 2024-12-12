// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import "github.com/AssetMantle/modules/helpers"

type auxiliaryResponse struct{}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func newAuxiliaryResponse() helpers.AuxiliaryResponse {
	return auxiliaryResponse{}
}
