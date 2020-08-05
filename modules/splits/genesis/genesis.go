package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sm "github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

//TODO define genesis state
type genesisState struct{
	splitList []mappables.InterNFT
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}

func (genesisState genesisState) Validate() error {

	return nil
}

func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {
	for _, split := range genesisState.splitList {
		mapper.Create(ctx, split)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	splitID := base.NewID("")

	var splitList []mappables.Split
	appendableSplitList := func(mappable traits.Mappable) bool {
		splitList = append(splitList, mappable.(sm.Split))
		return false
	}
	mapper.Iterate(context, splitID, appendableSplitList)
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

func newGenesisState(splitList []mappables.InterNFT) helpers.GenesisState {
	return genesisState{
		splitList: splitList,
	}
}

var GenesisState = newGenesisState([]mappables.InterNFT{})