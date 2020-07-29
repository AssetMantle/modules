package add

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
