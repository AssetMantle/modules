package mapper

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.Asset = (*baseAsset)(nil)

type baseAsset struct {
	assetID    assetID
	properties types.Properties
	lock       types.Height
	burn       types.Height
}

func (baseAsset baseAsset) String() string {
	bytes, Error := json.Marshal(baseAsset)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (baseAsset baseAsset) ID() types.ID { return baseAsset.assetID }

func (baseAsset baseAsset) ChainID() types.ID          { return baseAsset.assetID.chainID }
func (baseAsset baseAsset) ClassificationID() types.ID { return baseAsset.assetID.classificationID }
func (baseAsset baseAsset) MaintainersID() types.ID    { return baseAsset.assetID.maintainersID }
func (baseAsset baseAsset) HashID() types.ID           { return baseAsset.assetID.hashID }

func (baseAsset *baseAsset) Properties() types.Properties { return baseAsset.properties }

func (baseAsset baseAsset) GetLock() types.Height { return baseAsset.lock }
func (baseAsset baseAsset) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGraterThat(baseAsset.lock)
}

func (baseAsset baseAsset) GetBurn() types.Height { return baseAsset.burn }
func (baseAsset baseAsset) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGraterThat(baseAsset.burn)
}

func (baseAsset *baseAsset) MutateProperties(properties types.Properties) error {
	baseAsset.properties = properties
	return nil
}
func (baseAsset *baseAsset) MutateLock(lock types.Height) error { baseAsset.lock = lock; return nil }
func (baseAsset *baseAsset) MutateBurn(burn types.Height) error { baseAsset.burn = burn; return nil }
