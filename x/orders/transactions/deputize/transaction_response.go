// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/AssetMantle/modules/helpers"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (*TransactionResponse) GetResult() *sdkTypes.Result {
	return &sdkTypes.Result{}
}

func newTransactionResponse() *TransactionResponse {
	return &TransactionResponse{}
}
