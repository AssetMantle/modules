package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	am "github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

//TODO define genesis state
type genesisState struct{
	assetList []mappables.InterNFT
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}

func (genesisState genesisState) Validate() error {
	return nil
}

func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, asset := range genesisState.assetList {
		mapper.Create(ctx, asset)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	assetsID := base.NewID("")

	var assetList []mappables.InterNFT
	appendableAssetList := func(mappable traits.Mappable) bool {
		assetList = append(assetList, mappable.(am.Asset))
		return false
	}
	mapper.Iterate(context, assetsID, appendableAssetList)
	return genesisState
}

func (genesisState genesisState) Marshall() []byte {
	return PackageCodec.MustMarshalJSON(genesisState)
}
func (genesisState genesisState) Unmarshall(byte []byte) helpers.GenesisState {
	if Error := PackageCodec.UnmarshalJSON(byte, &genesisState); Error != nil {
		return nil
	}
	return genesisState
}

func newGenesisState(assetList []mappables.InterNFT) helpers.GenesisState {
	return genesisState{
		assetList: assetList,
	}
}

var GenesisState = newGenesisState([]mappables.InterNFT{})
