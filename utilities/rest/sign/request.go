// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package sign

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

type request struct {
	BaseRequest rest.BaseReq    `json:"baseReq"`
	Type        string          `json:"type" valid:"required~required field to missing, matches(^.*$)~invalid field type"`
	StdTx       authTypes.StdTx `json:"value"`
}

var _ helpers.Request = request{}

func (request request) Validate() error {
	_, err := govalidator.ValidateStruct(request)
	return err
}
