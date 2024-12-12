// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/key"
	"github.com/AssetMantle/modules/x/metas/record"
)

type transactionKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if !transactionKeeper.parameterManager.Fetch(context).Get().GetParameter(constantProperties.RevealEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("revealing is not enabled")
	}

	dataID := baseIDs.GenerateDataID(message.Data)
	metas := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(dataID))

	Mappable := metas.GetMappable(key.NewKey(dataID))
	if Mappable != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("data with ID %s already exists", dataID)
	}

	if message.Data.GenerateHashID().Compare(baseIDs.GenerateHashID()) != 0 {

		if err := message.Data.ValidateBasic(); err != nil {
			return nil, err
		}

		metas.Add(record.NewRecord(message.Data))
	}

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager
	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
