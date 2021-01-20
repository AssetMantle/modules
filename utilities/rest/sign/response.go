/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package sign

import (
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type response struct {
	Success bool            `json:"success"`
	Error   error           `json:"error"`
	StdTx   authTypes.StdTx `json:"tx"`
}

var _ helpers.Response = response{}

func (response response) IsSuccessful() bool {
	return response.Success
}
func (response response) GetError() error {
	return response.Error
}

func newResponse(stdTx authTypes.StdTx, error error) helpers.Response {
	success := true
	if error != nil {
		success = false
	}

	return response{
		Success: success,
		Error:   error,
		StdTx:   stdTx,
	}
}
