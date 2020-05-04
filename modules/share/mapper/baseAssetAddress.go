package mapper

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types"
)

type baseShareAddress struct {
	Address string `json:"address" yaml:"address" valid:"required~address"`
}

func NewShareAddress(address string) types.ShareAddress {
	return baseShareAddress{
		Address: address,
	}
}

var _ types.ShareAddress = (*baseShareAddress)(nil)

func (baseShareAddress baseShareAddress) Bytes() []byte { return []byte(baseShareAddress.Address) }
func (baseShareAddress baseShareAddress) String() string {
	bytes, Error := json.Marshal(baseShareAddress)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}
