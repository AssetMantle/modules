package mapper

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types"
)

type baseInterNFTAddress struct {
	Address string `json:"address" yaml:"address" valid:"required~address"`
}

var _ types.InterNFTAddress = (*baseInterNFTAddress)(nil)

func (baseInterNFTAddress baseInterNFTAddress) Bytes() []byte {
	return []byte(baseInterNFTAddress.Address)
}
func (baseInterNFTAddress baseInterNFTAddress) String() string {
	bytes, Error := json.Marshal(baseInterNFTAddress)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}
