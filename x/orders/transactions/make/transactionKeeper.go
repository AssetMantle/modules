// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"context"

	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/module"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/properties/utilities"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

type transactionKeeper struct {
	mapper                     helpers.Mapper
	parameterManager           helpers.ParameterManager
	bondAuxiliary              helpers.Auxiliary
	conformAuxiliary           helpers.Auxiliary
	supplementAuxiliary        helpers.Auxiliary
	transferAuxiliary          helpers.Auxiliary
	authenticateAuxiliary      helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	if _, err := transactionKeeper.maintainersVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.ClassificationID, message.FromID)); err != nil {
		return nil, err
	}

	fromAddress, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}
	makerOwnableSplit, ok := sdkTypes.NewIntFromString(message.MakerOwnableSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("MakerOwnableSplit is not a valid integer")
	}
	takerOwnableSplit, ok := sdkTypes.NewIntFromString(message.TakerOwnableSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("TakerOwnableSplit is not a valid integer")
	}
	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, module.ModuleIdentityID, message.MakerOwnableID, makerOwnableSplit)); err != nil {
		return nil, err
	}

	immutableMetaProperties := message.ImmutableMetaProperties.
		Add(baseProperties.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(takerOwnableSplit.ToDec().QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(makerOwnableSplit.ToDec())))).
		Add(baseProperties.NewMetaProperty(constants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())))).
		Add(baseProperties.NewMetaProperty(constants.MakerOwnableIDProperty.GetKey(), baseData.NewIDData(message.MakerOwnableID))).
		Add(baseProperties.NewMetaProperty(constants.TakerOwnableIDProperty.GetKey(), baseData.NewIDData(message.TakerOwnableID))).
		Add(baseProperties.NewMetaProperty(constants.MakerIDProperty.GetKey(), baseData.NewIDData(message.FromID))).
		Add(baseProperties.NewMetaProperty(constants.TakerIDProperty.GetKey(), baseData.NewIDData(message.TakerID)))

	immutables := baseQualified.NewImmutables(immutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.ImmutableProperties.GetList()...)...))

	orderID := baseIDs.NewOrderID(message.ClassificationID, immutables)
	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(orderID))
	if orders.Get(key.NewKey(orderID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("order with ID %s already exists", orderID.AsString())
	}

	if message.ExpiresIn.Get() > transactionKeeper.parameterManager.Fetch(context).GetParameter(constants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get() {
		return nil, errorConstants.InvalidRequest.Wrapf("order expiry exceeds maximum allowed %d", transactionKeeper.parameterManager.Fetch(context).GetParameter(constants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get())
	}

	mutableMetaProperties := message.MutableMetaProperties.
		Add(baseProperties.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+sdkTypes.UnwrapSDKContext(context).BlockHeight())))).
		Add(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewNumberData(makerOwnableSplit)))

	mutables := baseQualified.NewMutables(mutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.MutableProperties.GetList()...)...))

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.bondAuxiliary.GetKeeper().Help(context, bond.NewAuxiliaryRequest(message.ClassificationID, fromAddress)); err != nil {
		return nil, err
	}

	orders.Add(mappable.NewMappable(base.NewOrder(message.ClassificationID, immutables, mutables)))

	return newTransactionResponse(orderID.AsString()), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case bond.Auxiliary.GetName():
				transactionKeeper.bondAuxiliary = value
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.maintainersVerifyAuxiliary = value
			}
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
