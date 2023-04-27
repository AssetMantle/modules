// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"context"

	"github.com/AssetMantle/modules/x/splits/module"
	"github.com/AssetMantle/modules/x/splits/utilities"

	"github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	bankKeeper            bankKeeper.Keeper
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	fromAddress, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	var coinID ids.CoinID

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	if err := transactionKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), fromAddress, module.Name, message.Coins); err != nil {
		return nil, err
	}

	for _, coin := range message.Coins {
		if _, found := transactionKeeper.parameterManager.Fetch(context).GetParameter(constants.WrapAllowedCoinsProperty.GetID()).GetMetaProperty().GetData().Get().(*base.ListData).Search(base.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID(coin.Denom)))); !found {
			return nil, errorConstants.NotAuthorized.Wrapf("coin %s is not allowed to be wrapped", coin.Denom)
		}
		coinID = baseIDs.NewCoinID(baseIDs.NewStringID(coin.Denom))
		if _, err := utilities.AddSplits(transactionKeeper.mapper.NewCollection(context), message.FromID, coinID, coin.Amount); err != nil {
			return nil, err
		}
	}

	return newTransactionResponse(coinID.AsString()), nil
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
