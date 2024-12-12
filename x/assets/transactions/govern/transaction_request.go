// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package govern

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"net/http"
)

type transactionRequest struct {
}

func (transactionRequest transactionRequest) GetCommonTransactionRequest() helpers.CommonTransactionRequest {
	return helpers.CommonTransactionRequest{}
}

func (transactionRequest transactionRequest) FromCLI(_ helpers.CLICommand, _ client.Context) (helpers.TransactionRequest, error) {
	return requestPrototype(), constants.InvalidRequest.Wrapf("internal transaction,  please use the gov module to  create proposal to update params")
}

func (transactionRequest transactionRequest) FromHTTPRequest(_ *http.Request) (helpers.TransactionRequest, error) {
	return requestPrototype(), constants.InvalidRequest.Wrapf("internal transaction,  please use the gov module to  create proposal to update params")
}

func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	return messagePrototype(), constants.InvalidRequest.Wrapf("internal transaction,  please use the gov module to  create proposal to update params")
}

func (transactionRequest transactionRequest) Validate() error {
	return constants.InvalidRequest.Wrapf("internal transaction,  please use the gov module to  create proposal to update params")
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
