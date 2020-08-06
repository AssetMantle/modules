/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package recover

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type response struct {
	KeyOutput keyring.KeyOutput
}

var _ helpers.Response = response{}

func newResponse(keyOutput keyring.KeyOutput) *response {
	return &response{KeyOutput: keyOutput}
}
