// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"

	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/properties/utilities"
	"github.com/AssetMantle/schema/x/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	defineAuxiliary       helpers.Auxiliary
	superAuxiliary        helpers.Auxiliary
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

	immutables := base.NewImmutables(message.ImmutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.ImmutableProperties.GetList()...)...))

	mutables := base.NewMutables(message.MutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.MutableProperties.GetList()...)...))

	auxiliaryResponse, err := transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(fromAddress, immutables, mutables))
	if err != nil {
		return nil, err
	}
	classificationID := define.GetClassificationIDFromResponse(auxiliaryResponse)

	if _, err := transactionKeeper.superAuxiliary.GetKeeper().Help(context, super.NewAuxiliaryRequest(classificationID, message.FromID, mutables)); err != nil {
		return nil, err
	}

	return newTransactionResponse(classificationID.AsString()), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case super.Auxiliary.GetName():
				transactionKeeper.superAuxiliary = value
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
