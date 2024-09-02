// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"net/http"
)

type TransactionRequest interface {
	GetCommonTransactionRequest() CommonTransactionRequest

	FromCLI(CLICommand, client.Context) (TransactionRequest, error)
	FromHTTPRequest(*http.Request) (TransactionRequest, error)
	MakeMsg() (sdkTypes.Msg, error)
	Request
}
