/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package add

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type response struct {
	Success   bool           `json:"success"`
	Error     error          `json:"error"`
	KeyOutput keys.KeyOutput `json:"keyOutput"`
}

var _ helpers.Response = response{}

func (response response) IsSuccessful() bool {
	return response.Success
}
func (response response) GetError() error {
	return response.Error
}
func newResponse(keyOutput keys.KeyOutput, error error) helpers.Response {
	success := true
	if error != nil {
		success = false
	}

	return response{
		Success:   success,
		Error:     error,
		KeyOutput: keyOutput,
	}
}
