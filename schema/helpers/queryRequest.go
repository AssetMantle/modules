/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
)

type QueryRequest interface {
	Request
	FromCLI(CLICommand, client.Context) QueryRequest
	FromMap(map[string]string) QueryRequest
	LegacyAminoEncode() ([]byte, error)
	LegacyAminoDecode([]byte) (QueryRequest, error)
}
