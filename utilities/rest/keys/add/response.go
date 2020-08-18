/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package add

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type response struct {
	KeyOutput keys.KeyOutput
}

var _ helpers.Response = response{}

func newResponse(keyOutput keys.KeyOutput) *response {
	return &response{KeyOutput: keyOutput}
}
