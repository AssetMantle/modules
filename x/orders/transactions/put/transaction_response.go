// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (transactionResponse *TransactionResponse) GetResult() *sdkTypes.Result {
	return &sdkTypes.Result{
		Data: []byte(transactionResponse.OrderID.AsString()),
	}
}
func newTransactionResponse(orderID ids.OrderID) *TransactionResponse {
	return &TransactionResponse{orderID.(*baseIDs.OrderID)}
}
