// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/go/ids"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryRequest struct {
	OwnerID ids.IdentityID
	ids.AssetID
	Supply sdkTypes.Int
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.OwnerID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("owner id is invalid: %s", err.Error())
	}

	if err := auxiliaryRequest.AssetID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("asset id is invalid: %s", err.Error())
	}

	if auxiliaryRequest.Supply.IsNegative() {
		return constants.InvalidRequest.Wrapf("supply cannot be negative")
	}

	return nil
}

func NewAuxiliaryRequest(ownerID ids.IdentityID, assetID ids.AssetID, supply sdkTypes.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		OwnerID: ownerID,
		AssetID: assetID,
		Supply:  supply,
	}
}
