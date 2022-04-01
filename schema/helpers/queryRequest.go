// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client/context"
)

type QueryRequest interface {
	Request
	FromCLI(CLICommand, context.CLIContext) QueryRequest
	FromMap(map[string]string) QueryRequest
	Encode() ([]byte, error)
	Decode([]byte) (QueryRequest, error)
}
