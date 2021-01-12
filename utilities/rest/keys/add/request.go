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
	Name     string `json:"name"`
	Mnemonic string `json:"mnemonic"`
}

var _ helpers.Request = request{}

func (request request) Validate() error {
	_, Error := govalidator.ValidateStruct(request)
	return Error
}
