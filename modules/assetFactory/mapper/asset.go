package mapper

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFT = (*asset)(nil)

type asset struct {
	assetID    assetID
	properties types.Properties
	lock       types.Height
	burn       types.Height
}

func (asset asset) String() string {
	bytes, Error := json.Marshal(asset)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (asset asset) ID() types.ID { return asset.assetID }

func (asset asset) ChainID() types.ID          { return asset.assetID.chainID }
func (asset asset) ClassificationID() types.ID { return asset.assetID.classificationID }
func (asset asset) MaintainersID() types.ID    { return asset.assetID.maintainersID }
func (asset asset) HashID() types.ID           { return asset.assetID.hashID }

func (asset asset) Properties() types.Properties { return asset.properties }

func (asset asset) GetLock() types.Height { return asset.lock }
func (asset asset) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGraterThat(asset.lock)
}

func (asset asset) GetBurn() types.Height { return asset.burn }
func (asset asset) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGraterThat(asset.burn)
}
