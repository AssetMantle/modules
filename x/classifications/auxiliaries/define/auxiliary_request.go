// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/qualified"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryRequest struct {
	sdkTypes.AccAddress
	qualified.Immutables
	qualified.Mutables
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if auxiliaryRequest.AccAddress.Empty() {
		return constants.InvalidRequest.Wrapf("address cannot be empty")
	}

	if err := auxiliaryRequest.Immutables.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid immutables: %s", err.Error())
	}

	if err := auxiliaryRequest.Mutables.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid mutables: %s", err.Error())
	}

	return nil
}

func NewAuxiliaryRequest(accAddress sdkTypes.AccAddress, immutables qualified.Immutables, mutables qualified.Mutables) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		AccAddress: accAddress,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
