// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	"context"
	"github.com/AssetMantle/modules/x/assets/constants"

	"github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
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
