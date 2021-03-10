/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package sign

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type request struct {
	BaseRequest rest.BaseReq   `json:"baseReq"`
	Type        string         `json:"type" valid:"required~required field to missing, matches(^.*$)~invalid field type"`
	StdTx       legacytx.StdTx `json:"value"`
}

var _ helpers.Request = request{}

func (request request) Validate() error {
	_, Error := govalidator.ValidateStruct(request)
	return Error
}
