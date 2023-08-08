// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	"context"

	"github.com/AssetMantle/schema/go/data/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/utilities"
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

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	if err := transactionKeeper.bankKeeper.SendCoinsFromModuleToAccount(sdkTypes.UnwrapSDKContext(context), constants.ModuleName, fromAddress, message.Coins); err != nil {
		return nil, err
	}

	for _, coin := range message.Coins {
		if _, found := transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.UnwrapAllowedCoinsProperty.GetID()).GetMetaProperty().GetData().Get().(*base.ListData).Search(base.NewStringData(coin.Denom)); !found {
			return nil, errorConstants.NotAuthorized.Wrapf("coin %s is not allowed to be unwrapped", coin.Denom)
		}

		if _, err := utilities.SubtractSplits(transactionKeeper.mapper.NewCollection(context), message.FromID, baseIDs.GenerateCoinAssetID(coin.Denom), coin.Amount); err != nil {
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
			}
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
