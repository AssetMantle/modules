// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
	"github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	bankKeeper            bankKeeper.Keeper
	authenticateAuxiliary helpers.Auxiliary
	mintAuxiliary         helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

// Handle is a method of the transactionKeeper struct. It processes the transaction message passed to it.
//
// Parameters:
// - context (context.Context): Used for timing, cancelation signals, and carrying deadlines, among other things.
// - message (*Message): The transaction message being processed. It should contain details such as the source, destination, and coins involved.
//
// Return values:
// - *TransactionResponse: A response object detailing the result of the transaction.
// - error: In case of an error, this will contain the error message.
//
// The Handle method performs the following steps:
//
// 1. It authenticates the transaction request using the authenticateAuxiliary getter from the transactionKeeper object. If this fails, it returns the error encountered.
//
// 2. It fetches allowed coins from the parameter manager attached to the transactionKeeper object.
//
// 3. For each coin in the transaction message:
// - It first validates the coin asset. If this fails, it returns an error.
// - It then checks whether the coin value is negative. If it is, it returns an error.
// - It checks if the coin is on the list of allowed coins. If it isn't, it returns a not authorized error.
// - It tries to mint the coin using the mintAuxiliary getter of the transactionKeeper object. If this fails, the error is returned.
// - It fetches the coinAsset from the mapper attached to the transactionKeeper.
// - If the coinAsset isn't already in the collection, it is added using a new record.
// - It tries to send the coins from the sender's account to the module account using the bankKeeper attached to transactionKeeper object. If this fails, the error is returned.
//
// 4. If the process completes successfully for all coins, it returns a new transaction response object.
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}

	wrapAllowedCoins := transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.WrapAllowedCoinsProperty.GetID()).GetMetaProperty().GetData().Get().(*base.ListData)

	for _, coin := range message.Coins {
		coinAsset := baseDocuments.NewCoinAsset(coin.Denom)

		if err := coinAsset.ValidateCoinAsset(); err != nil {
			return nil, errorConstants.InvalidParameter.Wrapf("%s", err.Error())
		}

		if coin.IsNegative() {
			return nil, errorConstants.InvalidParameter.Wrapf("coin %s is negative", coin.Denom)
		}

		if _, found := wrapAllowedCoins.Search(base.NewStringData(coin.Denom)); !found {
			return nil, errorConstants.NotAuthorized.Wrapf("coin %s is not allowed to be wrapped", coin.Denom)
		}

		if _, err := transactionKeeper.mintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(message.GetFromIdentityID(), coinAsset.GetCoinAssetID(), coin.Amount)); err != nil {
			return nil, err
		}

		if assets := transactionKeeper.mapper.NewCollection(context); assets.Fetch(key.NewKey(coinAsset.GetCoinAssetID())).GetMappable(key.NewKey(coinAsset.GetCoinAssetID())) == nil {
			assets.Add(record.NewRecord(coinAsset))
		}

		if err := transactionKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), message.GetFromAddress(), constants.ModuleName, sdkTypes.NewCoins(coin)); err != nil {
			return nil, err
		}
	}

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bankKeeper.Keeper:
			transactionKeeper.bankKeeper = value
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case mint.Auxiliary.GetName():
				transactionKeeper.mintAuxiliary = value
			default:
				break
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
