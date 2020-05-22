package mapper

import (
	"bytes"
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.ID = (*baseAssetID)(nil)

type baseAssetID struct {
	chainID          types.ID
	maintainersID    types.ID
	classificationID types.ID
	hashID           types.ID
}

func (baseAssetID baseAssetID) Bytes() []byte {
	return append(append(append(
		baseAssetID.chainID.Bytes(),
		baseAssetID.maintainersID.Bytes()...),
		baseAssetID.classificationID.Bytes()...),
		baseAssetID.hashID.Bytes()...)
}

func (baseAssetID baseAssetID) String() string {
	bytes, Error := json.Marshal(baseAssetID)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (baseAssetID baseAssetID) IsEqualTo(id types.ID) bool {
	if bytes.Compare(baseAssetID.Bytes(), id.Bytes()) == 0 {
		return true
	}
	return false
}
