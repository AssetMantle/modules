// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"context"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
)

type transactionKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context, message.(*Message))
	return newTransactionResponse(err)
}
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {
	if !transactionKeeper.parameterManager.GetParameter(constantProperties.RevealEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized
	}

	dataID := baseIDs.GenerateDataID(message.Data)
	metas := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(dataID))

	Mappable := metas.Get(key.NewKey(dataID))
	if Mappable != nil {
		return nil, errorConstants.EntityAlreadyExists
	}

	if message.Data.GenerateHashID().Compare(baseIDs.GenerateHashID()) != 0 {
		metas.Add(mappable.NewMappable(message.Data))
	}

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameters
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
