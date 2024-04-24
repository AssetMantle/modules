// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"

	"github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/burn"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	bankKeeper            bankKeeper.Keeper
	burnAuxiliary         helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

// Handle is a method of the transactionKeeper struct. It processes the transaction message passed to it.
//
// Parameters:
// - context (context.Context): Used for timing, cancellation signals, and carrying deadlines, among other things.
// - message (*Message): The transaction message being processed. It contains information about the coins involved and the source of the transaction.
//
// Return values:
// - *TransactionResponse: A response object detailing the outcome of the transaction.
// - error: In case of an error, this will contain the error message.
//
// The Handle method performs the following steps:
//
// 1. It authenticates the transaction sender using the authenticateAuxiliary getter from the transactionKeeper object. If this fails, it returns the error encountered.
//
// 2. It fetches the list of allowed coins for unwrap operation from the parameter manager attached to the transactionKeeper object.
//
// 3. For each coin mentioned in the transaction message:
// - A new CoinAsset object is created and validated. If this is unsuccessful, it returns an error.
// - It checks whether the coin amount is negative. If it is, it returns an error.
// - Tests whether the coin is in the list of allowed coins for unwrap operation. If not, returns a not authorized error.
// - It fetches the coinAsset from the mapper collection attached to the transactionKeeper. If it's not present, returns an entity not found error.
// - It sends a burn request to the auxiliary to deduct the coin amount from the sender's account. If unsuccessful, the error is returned.
// - It attempts to transfer the coins from the module to the sender's account using the bankKeeper attached to transactionKeeper object. If this fails, the error is returned.
//
// 4. If the process completes successfully for all coins, it returns a new TransactionResponse object.
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message.GetFromAddress(), message.FromID)); err != nil {
		return nil, err
	}

	unwrapAllowedCoins := transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.UnwrapAllowedCoinsProperty.GetID()).GetMetaProperty().GetData().Get().(*base.ListData)

	for _, coin := range message.Coins {
		coinAsset := baseDocuments.NewCoinAsset(coin.Denom)

		if err := coinAsset.ValidateCoinAsset(); err != nil {
			return nil, errorConstants.InvalidParameter.Wrapf("%s", err.Error())
		}

		if coin.IsNegative() {
			return nil, errorConstants.InvalidParameter.Wrapf("coin %s is negative", coin.Denom)
		}

		if _, found := unwrapAllowedCoins.Search(base.NewStringData(coin.Denom)); !found {
			return nil, errorConstants.NotAuthorized.Wrapf("coin %s is not allowed to be unwrapped", coin.Denom)
		}

		if transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(coinAsset.GetCoinAssetID())).GetMappable(key.NewKey(coinAsset.GetCoinAssetID())) == nil {
			return nil, errorConstants.EntityNotFound.Wrapf("asset %s not found", coinAsset.GetDenom())
		}

		if _, err := transactionKeeper.burnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(message.FromID, coinAsset.GetCoinAssetID(), coin.Amount)); err != nil {
			return nil, err
		}

		if err := transactionKeeper.bankKeeper.SendCoinsFromModuleToAccount(sdkTypes.UnwrapSDKContext(context), constants.ModuleName, message.GetFromAddress(), sdkTypes.NewCoins(coin)); err != nil {
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
			case burn.Auxiliary.GetName():
				transactionKeeper.burnAuxiliary = value
			}
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
