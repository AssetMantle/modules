package recover

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type request struct {
	Name     string `json:"name"`
	Mnemonic string `json:"mnemonic"`
}

var _ helpers.Request = request{}
