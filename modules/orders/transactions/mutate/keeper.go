package mutate

import (
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
	"reflect"
)

type transactionKeeper struct {
	mapper types.Mapper
	bk     bank.Keeper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	assetID := message.AssetID
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	asset := assets.Get(assetID)
	if asset == nil {
		return constants.EntityNotFound
	}
	mutableProperties := asset.GetMutables().Get()
	//for _, property := range message.Properties.GetList() {
	//	if mutableProperties.Get(property.GetID()) == nil {
	//		return constants.MutableNotFound
	//	}
	//	mutableProperties = mutableProperties.Mutate(property)
	//}
	assetStruct, _ := asset.(*mapper.Order)
	//transactionKeeper.bk.SubtractCoins(context, assetStruct.From, sdkTypes.Coins{assetStruct.SellOrder})
	//transactionKeeper.bk.AddCoins(context, assetStruct.From, sdkTypes.Coins{assetStruct.BuyOrder})
	//transactionKeeper.bk.SubtractCoins(context, message.From, sdkTypes.Coins{assetStruct.BuyOrder})
	//transactionKeeper.bk.AddCoins(context, message.From, sdkTypes.Coins{assetStruct.SellOrder})
	asset = mapper.NewOrder(asset.GetID(), message.From, assetStruct.BuyOrder, assetStruct.SellOrder, types.NewMutables(mutableProperties, asset.GetClassificationID()), asset.GetImmutables(), types.NewHeight(5), asset.GetBurn())
	assets = assets.Mutate(asset)
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper, externalKeepers ...interface{}) types.TransactionKeeper {
	bk, ok := externalKeepers[0].(*bank.BaseKeeper)
	fmt.Println(ok, bk, "0")

	for _, n := range externalKeepers {
		bk7, ok7 := n.(*bank.Keeper)
		fmt.Printf("Hello %s\n", ok7, bk7)
	}

	bk1, ok1 := externalKeepers[0].(*bank.Keeper)
	fmt.Println(ok1, bk1, "1")

	bk5, ok5 := externalKeepers[0].(interface{})
	fmt.Println(ok5, "bk5", "5")

	_, ok6 := bk5.(bank.Keeper)
	fmt.Println(ok6, "bk6", "6")
	b := externalKeepers[0]
	fmt.Println("type b = ", reflect.TypeOf(b), reflect.TypeOf(bk5))

	bk4, ok4 := externalKeepers[0].(bank.Keeper)
	fmt.Println(ok4, bk4, "4")

	bk2, ok2 := externalKeepers[0].(*bank.SendKeeper)
	fmt.Println(ok2, bk2, "2")

	bk3, ok3 := externalKeepers[0].(*bank.BaseSendKeeper)
	fmt.Println(ok3, bk3, "3")

	switch externalKeepers[0].(type) {
	case bank.BaseKeeper:
		fmt.Println("here2")
	case bank.Keeper:
		fmt.Println("here")
	case bank.BaseSendKeeper:
		fmt.Println("here1")
	case bank.BaseViewKeeper:
		fmt.Println("here0")
		fmt.Println("here0")
	default:
		fmt.Println("no type")
	}
	return transactionKeeper{mapper: mapper, bk: bk}
}
