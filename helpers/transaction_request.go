// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"io"
	"net/http"
)

type TransactionRequest interface {
	GetCommonTransactionRequest() CommonTransactionRequest

	FromCLI(CLICommand, client.Context) (TransactionRequest, error)
	FromHTTPRequest(*http.Request) (TransactionRequest, error)
	MakeMsg() (sdkTypes.Msg, error)
	Request
}

func TransactionRequestFromHTTPRequest[T TransactionRequest](httpRequest *http.Request, TransactionRequest *T) (TransactionRequest, error) {
	body, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, TransactionRequest); err != nil {
		return nil, err
	}

	return *TransactionRequest, nil
}
func Validate[T TransactionRequest](transactionRequest T) error {
	if msg, err := transactionRequest.MakeMsg(); err != nil {
		return err
	} else if err := msg.(Message).ValidateBasic(); err != nil {
		return err
	}

	return nil
}
