// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import "github.com/AssetMantle/schema/x/helpers"

type auxiliaryResponse struct {
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func newAuxiliaryResponse() helpers.AuxiliaryResponse {
	return auxiliaryResponse{}
}
