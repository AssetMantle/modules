// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	authorizeAuxiliary    helpers.Auxiliary
	bondAuxiliary         helpers.Auxiliary
	conformAuxiliary      helpers.Auxiliary
	mintAuxiliary         helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	if !transactionKeeper.parameterManager.Fetch(context).Get().GetParameter(propertyConstants.MintEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("minting is not enabled")
	}

	if _, err := transactionKeeper.authorizeAuxiliary.GetKeeper().Help(context, authorize.NewAuxiliaryRequest(message.ClassificationID, message.GetFromIdentityID(), constants.CanMintAssetPermission)); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}

	immutables := baseQualified.NewImmutables(message.ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.ImmutableProperties.Get()...)...))

	assetID := baseIDs.NewAssetID(message.ClassificationID, immutables)

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(assetID))
	if assets.GetMappable(key.NewKey(assetID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("asset with ID %s already exists", assetID.AsString())
	}

	mutables := baseQualified.NewMutables(message.MutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.MutableProperties.Get()...)...))

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); err != nil {
		return nil, err
	}

	asset := base.NewAsset(message.ClassificationID, immutables, mutables)

	if err := asset.ValidateBasic(); err != nil {
		return nil, err
	}

	supply := asset.GetSupply()

	if supply.IsNegative() {
		return nil, errorConstants.IncorrectFormat.Wrapf("asset supply is negative")
	}

	if _, err := transactionKeeper.mintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(message.ToID, assetID, supply)); err != nil {
		return nil, err
	}

	bondAmount := sdkTypes.ZeroInt()
	if bondAmountProperty := mutables.GetProperty(propertyConstants.BondAmountProperty.GetID()); bondAmountProperty == nil || !bondAmountProperty.IsMeta() {
		return nil, errorConstants.MetaDataError.Wrapf("asset with ID %s has no revealed bond amount", assetID.AsString())
	} else {
		bondAmount = bondAmountProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	if _, err := transactionKeeper.bondAuxiliary.GetKeeper().Help(context, bond.NewAuxiliaryRequest(message.ClassificationID, message.GetFromAddress(), bondAmount)); err != nil {
		return nil, err
	}

	assets.Add(record.NewRecord(asset))

	return newTransactionResponse(assetID), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper
	transactionKeeper.parameterManager = parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
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
			case mint.Auxiliary.GetName():
				transactionKeeper.mintAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
