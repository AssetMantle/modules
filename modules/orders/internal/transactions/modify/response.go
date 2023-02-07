// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"errors"
	"github.com/AssetMantle/modules/schema/helpers"
)

//type transactionResponse struct {
//	Success bool  `json:"success"`
//	Error   error `json:"error" swaggertype:"string"`
//}

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (transactionResponse *TransactionResponse) IsSuccessful() bool {
	return transactionResponse.Success
}
func (transactionResponse *TransactionResponse) GetError() error {
	return errors.New(transactionResponse.Error)
}
func newTransactionResponse(error error) helpers.TransactionResponse {
	if error != nil {
		return &TransactionResponse{
			Success: false,
			Error:   error.Error(),
		}
	}

	return &TransactionResponse{
		Success: true,
	}
}
