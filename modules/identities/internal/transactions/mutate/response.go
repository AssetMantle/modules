// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

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
