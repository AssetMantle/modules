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
	Name     string `json:"name" valid:"required~required field to missing, matches(.+?)~invalid field name"`
	Mnemonic string `json:"mnemonic" valid:"optional"`
}

var _ helpers.Request = request{}

func (request request) Validate() error {
	_, err := govalidator.ValidateStruct(request)
	return err
}
