// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"context"

	"github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
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

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	fromAddress := message.GetFromAddress()

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	if err := transactionKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), fromAddress, constants.ModuleName, message.Coins); err != nil {
		return nil, err
	}

	for _, coin := range message.Coins {
		if err := sdkTypes.ValidateDenom(coin.Denom); err != nil {
			return nil, errorConstants.InvalidRequest.Wrapf("coin denom %s is invalid", coin.Denom)
		}

		if _, found := transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.WrapAllowedCoinsProperty.GetID()).GetMetaProperty().GetData().Get().(*base.ListData).Search(base.NewStringData(coin.Denom)); !found {
			return nil, errorConstants.NotAuthorized.Wrapf("coin %s is not allowed to be wrapped", coin.Denom)
		}

		coinAsset := baseDocuments.NewCoinAsset(coin.Denom)

		value := coin.Amount
		if value.IsNegative() {
			return nil, errorConstants.InvalidRequest.Wrapf("coin amount %s is negative", coin.Amount)
		}

		if _, err := transactionKeeper.mintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(message.FromID, coinAsset.GetCoinAssetID(), value)); err != nil {
			return nil, err
		}

		assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(coinAsset.GetCoinAssetID()))

		Mappable := assets.GetMappable(key.NewKey(coinAsset.GetCoinAssetID()))
		if Mappable == nil {
			if err := coinAsset.ValidateBasic(); err != nil {
				return nil, err
			}

			assets.Add(record.NewRecord(coinAsset))
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
