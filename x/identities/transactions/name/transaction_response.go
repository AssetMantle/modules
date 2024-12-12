// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package name

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (transactionResponse *TransactionResponse) GetResult() *sdkTypes.Result {
	return &sdkTypes.Result{
		Data: []byte(transactionResponse.NameIdentityID.AsString()),
	}
}
func newTransactionResponse(nameIdentityID ids.IdentityID) *TransactionResponse {
	return &TransactionResponse{nameIdentityID.(*baseIDs.IdentityID)}
}
