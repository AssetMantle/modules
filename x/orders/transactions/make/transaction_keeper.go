// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/record"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	authorizeAuxiliary    helpers.Auxiliary
	bondAuxiliary         helpers.Auxiliary
	conformAuxiliary      helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	if _, err := transactionKeeper.authorizeAuxiliary.GetKeeper().Help(context, authorize.NewAuxiliaryRequest(message.ClassificationID, message.GetFromIdentityID())); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.authorizeAuxiliary.GetKeeper().Help(context, authorize.NewAuxiliaryRequest(message.GetClassificationID(), message.GetFromIdentityID(), constants.CanMakeOrderPermission)); err != nil {
		// TODO enable after permissioning is implemented
		// return nil, err
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}
	makerSplit, ok := sdkTypes.NewIntFromString(message.MakerSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("MakerSplit is not a valid integer")
	}
	takerSplit, ok := sdkTypes.NewIntFromString(message.TakerSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("TakerSplit is not a valid integer")
	}
	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.GetFromIdentityID(), constants.ModuleIdentity.GetModuleIdentityID(), message.MakerAssetID, makerSplit)); err != nil {
		return nil, err
	}

	immutableMetaProperties := message.ImmutableMetaProperties.
		Add(baseProperties.NewMetaProperty(propertyConstants.ExchangeRateProperty.GetKey(), baseData.NewDecData(takerSplit.ToLegacyDec().QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(makerSplit.ToLegacyDec())))).
		Add(baseProperties.NewMetaProperty(propertyConstants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())))).
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(message.MakerAssetID))).
		Add(baseProperties.NewMetaProperty(propertyConstants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(message.TakerAssetID))).
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerIDProperty.GetKey(), baseData.NewIDData(message.GetFromIdentityID()))).
		Add(baseProperties.NewMetaProperty(propertyConstants.TakerIDProperty.GetKey(), baseData.NewIDData(message.TakerID)))

	immutables := baseQualified.NewImmutables(immutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.ImmutableProperties.Get()...)...))

	orderID := baseIDs.NewOrderID(message.ClassificationID, immutables)

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(orderID))
	if orders.GetMappable(key.NewKey(orderID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("order with ID %s already exists", orderID.AsString())
	}

	if message.ExpiresIn.Get() > transactionKeeper.parameterManager.Fetch(context).Get().GetParameter(propertyConstants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get() {
		return nil, errorConstants.InvalidRequest.Wrapf("order expiry exceeds maximum allowed %d", transactionKeeper.parameterManager.Fetch(context).Get().GetParameter(propertyConstants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get())
	}

	mutableMetaProperties := message.MutableMetaProperties.
		Add(baseProperties.NewMetaProperty(propertyConstants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+sdkTypes.UnwrapSDKContext(context).BlockHeight())))).
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(makerSplit)))

	mutables := baseQualified.NewMutables(mutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.MutableProperties.Get()...)...))

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); err != nil {
		return nil, err
	}

	bondAmount := sdkTypes.ZeroInt()
	if bondAmountProperty := mutables.GetProperty(propertyConstants.BondAmountProperty.GetID()); bondAmountProperty == nil || !bondAmountProperty.IsMeta() {
		return nil, errorConstants.MetaDataError.Wrapf("order with ID %s has no revealed bond amount", orderID.AsString())
	} else {
		bondAmount = bondAmountProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	if _, err := transactionKeeper.bondAuxiliary.GetKeeper().Help(context, bond.NewAuxiliaryRequest(message.ClassificationID, message.GetFromAddress(), bondAmount)); err != nil {
		return nil, err
	}

	order := base.NewOrder(message.ClassificationID, immutables, mutables)

	if err := order.ValidateBasic(); err != nil {
		return nil, err

	}
	orders.Add(record.NewRecord(order))

	return newTransactionResponse(orderID), nil
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
			case bond.Auxiliary.GetName():
				transactionKeeper.bondAuxiliary = value
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
