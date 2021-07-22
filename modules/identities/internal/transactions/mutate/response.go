/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mutate

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error" swaggertype:"string"`
}

var _ helpers.TransactionResponse = (*transactionResponse)(nil)

// Transaction Request godoc
// @Summary nub identities transaction
// @Descrption nub transaction
// @Accept text/plain
// @Produce json
// @Tags Identities
// @Param body body  transactionRequest true "request body"
// @Success 200 {object} transactionResponse   "A successful response."
// @Failure default  {object}  transactionResponse "An unexpected error response."
// @Router /identities/nub [post]
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
