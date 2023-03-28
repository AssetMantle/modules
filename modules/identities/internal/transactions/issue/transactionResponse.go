// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (*TransactionResponse) GetResult() *sdkTypes.Result {
	return &sdkTypes.Result{}
}

func newTransactionResponse(identityID string) *TransactionResponse {
	return &TransactionResponse{identityID}
}
