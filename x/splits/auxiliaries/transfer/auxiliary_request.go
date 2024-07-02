// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryRequest struct {
	FromID  ids.IdentityID
	ToID    ids.IdentityID
	AssetID ids.AssetID
	Value   sdkTypes.Int
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.FromID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid from id: %s", err)
	}

	if err := auxiliaryRequest.ToID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid to id: %s", err)
	}

	if err := auxiliaryRequest.AssetID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid asset id: %s", err)
	}

	if auxiliaryRequest.Value.LTE(sdkTypes.ZeroInt()) {
		return constants.InvalidRequest.Wrapf("value must be greater than zero")
	}

	return nil
}

func NewAuxiliaryRequest(fromID ids.IdentityID, toID ids.IdentityID, assetID ids.AssetID, value sdkTypes.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:  fromID,
		ToID:    toID,
		AssetID: assetID,
		Value:   value,
	}
}
