package make

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper     helpers.Mapper
	BankKeeper bank.Keeper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	salt := base.NewHeight(context.BlockHeight())
	orderHash := message.GenerateHash(salt)
	orderHashProperty := base.NewProperty(base.NewID(OrderHash), base.NewFact(orderHash.String(), nil))
	properties := message.Properties.Add(orderHashProperty)
	immutables := base.NewImmutables(properties)
	orderID := mapper.NewOrderID(base.NewID(context.ChainID()), immutables.GetHashID())
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	if orders.Get(orderID) != nil {
		return constants.EntityAlreadyExists
	}
	makerAsset, Error := configureAssetData(message.MakerAssetType, message.MakerAssetData, message.MakerAssetAmount)
	if Error != nil {
		return Error
	}
	Error = checkExists(context, transactionKeeper.BankKeeper, message.From, makerAsset)
	if Error != nil {
		return Error
	}
	takerAsset, Error := configureAssetData(message.TakerAssetType, message.TakerAssetData, message.TakerAssetAmount)
	if Error != nil {
		return Error
	}

	orders.Add(mapper.NewOrder(orderID, message.Burn, message.Lock, immutables, message.From, message.TakerAddress,
		message.MakerAssetAmount, makerAsset, message.TakerAssetAmount, takerAsset, salt))
	return nil
}

func configureAssetData(assetType types.ID, assetData types.ID, assetAmount sdkTypes.Dec) (traits.Exchangeable, error) {

	switch assetType.String() {
	case "coin":
		return sdkTypes.NewCoin(assetData.String(), assetAmount.TruncateInt()), nil
	default:
		return nil, errors.New(fmt.Sprintf("type %v not supported", assetType.String()))
	}
}

func checkExists(context sdkTypes.Context, bankKeeper bank.Keeper, makerAddress sdkTypes.AccAddress, asset traits.Exchangeable) error {
	switch value := asset.(type) {
	case sdkTypes.Coin:
		{
			if !bankKeeper.HasBalance(context, makerAddress, value) {
				return errors.New(fmt.Sprintf("insufficient amount to place order"))
			}
		}
	}
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case bank.Keeper:
			transactionKeeper.BankKeeper = value
		}
	}
	return transactionKeeper
}
