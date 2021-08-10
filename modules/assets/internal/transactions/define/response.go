/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error" swaggertype:"string"`
}


var _ helpers.TransactionResponse = (*transactionResponse)(nil)

func (transactionResponse transactionResponse) IsSuccessful() bool {
	return transactionResponse.Success
}
func (transactionResponse transactionResponse) GetError() error {
	return transactionResponse.Error
}
func newTransactionResponse(error error) helpers.TransactionResponse {
	success := true
	if error != nil {
		success = false
	}

	return transactionResponse{
		Success: success,
		Error:   error,
	}
}
