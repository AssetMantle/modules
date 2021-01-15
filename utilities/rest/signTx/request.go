/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package signTx

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/types/rest"
	//sdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type request struct {
	BaseRequest rest.BaseReq    `json:"baseReq"`
	Type        string          `json:"type" valid:"required~required field to missing, matches(^.*$)~invalid field type"`
	StdTx       authTypes.StdTx `json:"value"`
}

var _ helpers.Request = request{}

func (request request) Validate() error {
	_, Error := govalidator.ValidateStruct(request)
	return Error
}
