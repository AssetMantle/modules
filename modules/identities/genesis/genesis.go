package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	im "github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

//TODO define genesis state
type genesisState struct{
	identityList []mappables.InterNFT
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}

func (genesisState genesisState) Validate() error {

	return nil
}

func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, identity := range genesisState.identityList {
		mapper.Create(ctx, identity)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	identityID := base.NewID("")

	var identityList []mappables.InterIdentity
	appendableIdentityList := func(mappable traits.Mappable) bool {
		identityList = append(identityList, mappable.(im.Identity))
		return false
	}
	mapper.Iterate(context, identityID, appendableIdentityList)
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

func newGenesisState(identityList []mappables.InterNFT) helpers.GenesisState {
	return genesisState{
		identityList: identityList,
	}
}

var GenesisState = newGenesisState([]mappables.InterNFT{})