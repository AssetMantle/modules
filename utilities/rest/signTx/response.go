/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package signTx

import (
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type response struct {
	Tx authTypes.StdTx `json:"tx"`
}

var _ helpers.Response = response{}

func newResponse(signedStdTx authTypes.StdTx) *response {
	return &response{Tx: signedStdTx}
}
