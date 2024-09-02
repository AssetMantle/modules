// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryRequest struct {
	OwnerID ids.IdentityID
	ids.AssetID
	Value sdkTypes.Int
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.OwnerID.ValidateBasic(); err != nil {
		return errorConstants.InvalidRequest.Wrapf("invalid owner id: %s", err)
	}

	if err := auxiliaryRequest.AssetID.ValidateBasic(); err != nil {
		return errorConstants.InvalidRequest.Wrapf("invalid asset id: %s", err)
	}

	if auxiliaryRequest.Value.LTE(sdkTypes.ZeroInt()) {
		return errorConstants.InvalidRequest.Wrapf("value must be greater than zero")
	}

	return nil
}

func NewAuxiliaryRequest(ownerID ids.IdentityID, assetID ids.AssetID, value sdkTypes.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		OwnerID: ownerID,
		AssetID: assetID,
		Value:   value,
	}
}
