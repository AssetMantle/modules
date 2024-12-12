// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/lists"
)

type auxiliaryRequest struct {
	lists.PropertyList
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	if err := auxiliaryRequest.PropertyList.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf("invalid property list: %s", err.Error())
	}

	return nil
}

func NewAuxiliaryRequest(propertyList lists.PropertyList) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		PropertyList: propertyList,
	}
}
