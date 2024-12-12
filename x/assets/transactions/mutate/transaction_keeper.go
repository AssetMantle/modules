// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/schema/documents/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	authenticateAuxiliary helpers.Auxiliary
	maintainAuxiliary     helpers.Auxiliary
	conformAuxiliary      helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.AssetID))

	Mappable := assets.GetMappable(key.NewKey(message.AssetID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("asset with ID %s not found", message.AssetID.AsString())
	}
	asset := mappable.GetAsset(Mappable)

	mutables := baseQualified.NewMutables(message.MutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.MutableProperties.Get()...)...))

	if _, err := transactionKeeper.maintainAuxiliary.GetKeeper().Help(context, maintain.NewAuxiliaryRequest(asset.GetClassificationID(), message.GetFromIdentityID(), mutables)); err != nil {
		return nil, err
	}

	mutatedAsset := base.NewAsset(asset.GetClassificationID(), asset.GetImmutables(), asset.GetMutables().Mutate(baseLists.AnyPropertiesToProperties(mutables.GetMutablePropertyList().Get()...)...))

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(asset.GetClassificationID(), mutatedAsset.GetImmutables(), mutatedAsset.GetMutables())); err != nil {
		return nil, err
	}

	if err := asset.ValidateBasic(); err != nil {
		return nil, err
	}

	assets.Mutate(record.NewRecord(mutatedAsset))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case maintain.Auxiliary.GetName():
				transactionKeeper.maintainAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
