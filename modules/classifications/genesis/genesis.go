package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	cm "github.com/persistenceOne/persistenceSDK/modules/classifications/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

//TODO define genesis state
type genesisState struct{
	classificationList []mappables.InterNFT
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}

func (genesisState genesisState) Validate() error {
	return nil
}

func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, cls := range genesisState.classificationList {
		mapper.Create(ctx, cls)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	clsID := base.NewID("")

	var clsList []mappables.Classification
	appendableClsList := func(mappable traits.Mappable) bool {
		clsList = append(clsList, mappable.(cm.Classification))
		return false
	}
	mapper.Iterate(context, clsID, appendableClsList)
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

func newGenesisState(classificationList []mappables.InterNFT) helpers.GenesisState {
	return genesisState{
		classificationList: classificationList,
	}
}

var GenesisState = newGenesisState([]mappables.InterNFT{})