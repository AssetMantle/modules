package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	om "github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

//TODO define genesis state
type genesisState struct{
	orderList []mappables.InterNFT
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}

func (genesisState genesisState) Validate() error {

	return nil
}

func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, order := range genesisState.orderList {
		mapper.Create(ctx, order)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	orderID := base.NewID("")

	var orderList []mappables.Order
	appendableOrderList := func(mappable traits.Mappable) bool {
		orderList = append(orderList, mappable.(om.Order))
		return false
	}
	mapper.Iterate(context, orderID, appendableOrderList)
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

func newGenesisState(orderList []mappables.InterNFT) helpers.GenesisState {
	return genesisState{
		orderList: orderList,
	}
}

var GenesisState = newGenesisState([]mappables.InterNFT{})