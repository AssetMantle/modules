// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import (
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	Address sdkTypes.AccAddress
	ids.IdentityID
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if auxiliaryRequest.Address.Empty() {
		return constants.InvalidRequest.Wrapf("address cannot be empty")
	}

	if err := auxiliaryRequest.IdentityID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid identity id: %s", err.Error())
	}

	return nil
}

func NewAuxiliaryRequest(address sdkTypes.AccAddress, identityID ids.IdentityID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		Address:    address,
		IdentityID: identityID,
	}
}
