// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
)

type QueryRequest interface {
	GRPCRequest
	FromCLI(CLICommand, client.Context) (QueryRequest, error)
	FromHTTPRequest(*http.Request) (QueryRequest, error)
}
