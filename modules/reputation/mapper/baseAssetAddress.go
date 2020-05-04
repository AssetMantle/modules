package mapper

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types"
)

type baseAssetAddress struct {
	Address string `json:"address" yaml:"address" valid:"required~address"`
}

var _ types.AssetAddress = (*baseAssetAddress)(nil)

func (baseAssetAddress baseAssetAddress) Bytes() []byte { return []byte(baseAssetAddress.Address) }
func (baseAssetAddress baseAssetAddress) String() string {
	bytes, Error := json.Marshal(baseAssetAddress)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}
