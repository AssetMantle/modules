// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"github.com/AssetMantle/modules/helpers"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (transactionResponse *TransactionResponse) GetResult() *sdkTypes.Result {
	return &sdkTypes.Result{
		Data: []byte(transactionResponse.ClassificationID),
	}
}

func newTransactionResponse(classificationID string) *TransactionResponse {
	return &TransactionResponse{classificationID}
}
