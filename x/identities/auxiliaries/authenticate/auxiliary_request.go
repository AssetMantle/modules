// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryRequest struct {
	from   sdkTypes.AccAddress
	fromID ids.IdentityID
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if auxiliaryRequest.from.Empty() {
		return constants.InvalidRequest.Wrapf("address cannot be empty")
	}

	if err := auxiliaryRequest.fromID.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid identity id: %s", err.Error())
	}

	return nil
}

func NewAuxiliaryRequest(message helpers.Message) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		from:   message.GetFromAddress(),
		fromID: message.GetFromIdentityID(),
	}
}
