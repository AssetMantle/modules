// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"github.com/AssetMantle/modules/schema/helpers"
)

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (transactionResponse *TransactionResponse) IsSuccessful() bool {
	return transactionResponse.Success
}
func newTransactionResponse(error error) helpers.TransactionResponse {
	success := true
	if error != nil {
		success = false
	}

	return &TransactionResponse{
		Success: success,
		Error:   error.Error(),
	}
}
