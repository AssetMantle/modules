// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"context"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/properties"
	"github.com/AssetMantle/schema/go/properties/constants"
	"github.com/AssetMantle/schema/go/properties/utilities"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                     helpers.Mapper
	parameterManager           helpers.ParameterManager
	authenticateAuxiliary      helpers.Auxiliary
	bondAuxiliary              helpers.Auxiliary
	conformAuxiliary           helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
	mintAuxiliary              helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	if !transactionKeeper.parameterManager.Fetch(context).GetParameter(constants.MintEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("minting is not enabled")
	}

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

	immutables := baseQualified.NewImmutables(message.ImmutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.ImmutableProperties.GetList()...)...))

	assetID := baseIDs.NewAssetID(message.ClassificationID, immutables)

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(assetID))
	if assets.Get(key.NewKey(assetID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("asset with ID %s already exists", assetID.AsString())
	}

	mutables := baseQualified.NewMutables(message.MutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.MutableProperties.GetList()...)...))

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); err != nil {
		return nil, err
	}

	split := sdkTypes.OneInt()

	if metaPropertyList := message.ImmutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.MutableMetaProperties.GetList()...)...); metaPropertyList.GetProperty(constants.SupplyProperty.GetID()) != nil {
		split = metaPropertyList.GetProperty(constants.SupplyProperty.GetID()).Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	if _, err := transactionKeeper.mintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(message.ToID, assetID, split)); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.bondAuxiliary.GetKeeper().Help(context, bond.NewAuxiliaryRequest(message.ClassificationID, fromAddress)); err != nil {
		return nil, err
	}

	assets.Add(mappable.NewMappable(base.NewAsset(message.ClassificationID, immutables, mutables)))

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
			case bond.Auxiliary.GetName():
				transactionKeeper.bondAuxiliary = value
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case mint.Auxiliary.GetName():
				transactionKeeper.mintAuxiliary = value
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
