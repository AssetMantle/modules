// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	"net/http"
)

type QueryRequest interface {
	Request
	FromCLI(CLICommand, client.Context) (QueryRequest, error)
	FromHTTPRequest(*http.Request) (QueryRequest, error)
	Encode() ([]byte, error)
	Decode([]byte) (QueryRequest, error)
}
