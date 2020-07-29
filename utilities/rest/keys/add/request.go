package add

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type request struct {
	Name string `json:"name"`
}

var _ helpers.Request = request{}
