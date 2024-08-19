// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package cancel

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/unbond"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/record"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	authorizeAuxiliary    helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
	unbondAuxiliary       helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	fromAddress := message.GetSigners()[0]

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.OrderID))

	Mappable := orders.GetMappable(key.NewKey(message.OrderID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("order with ID %s not found", message.OrderID.AsString())
	}
	order := mappable.GetOrder(Mappable)

	//if _, err := transactionKeeper.authorizeAuxiliary.GetKeeper().Help(context, authorize.NewAuxiliaryRequest(order.GetClassificationID(), message.FromID, constants.CanCancelOrderPermission)); err != nil {
	//	 TODO enable after permissioning is implemented
	//	 return nil, err
	//}

	//if _, err := transactionKeeper.unbondAuxiliary.GetKeeper().Help(context, unbond.NewAuxiliaryRequest(order.GetClassificationID(), fromAddress)); err != nil {
	//	 TODO enable after order bonding is implemented
	//	return nil, err
	//}

	if message.FromID.Compare(order.GetMakerID()) != 0 {
		return nil, errorConstants.NotAuthorized.Wrapf("order with ID %s not owned by %s", message.OrderID.AsString(), message.FromID.AsString())
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), message.FromID, order.GetMakerAssetID(), order.GetMakerSplit())); err != nil {
		return nil, err
	}

	orders.Remove(record.NewRecord(order))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case authorize.Auxiliary.GetName():
				transactionKeeper.authorizeAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			case unbond.Auxiliary.GetName():
				transactionKeeper.unbondAuxiliary = value
			}
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
