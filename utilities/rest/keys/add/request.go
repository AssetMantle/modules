/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package add

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type request struct {
	Name     string `json:"name" valid:"required~required field to missing, matches(^[A-Za-z0-9]+$)~invalid field name"`
	Mnemonic string `json:"mnemonic" valid:"optional"`
}

var _ helpers.Request = request{}

func (request request) Validate() error {
	_, Error := govalidator.ValidateStruct(request)
	return Error
}
