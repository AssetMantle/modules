// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/properties/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
	bankKeeper       bankKeeper.Keeper
	stakingKeeper    stakingKeeper.Keeper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	totalWeight := int64(0)
	for _, property := range append(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList(), auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()...) {
		totalWeight += property.Get().GetBondWeight()
	}
	bondAmount := baseData.NewNumberData(auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constants.BondRateProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get() * totalWeight)
	immutables := baseQualified.NewImmutables(auxiliaryRequest.Immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), bondAmount)))

	if int64(len(immutables.GetImmutablePropertyList().GetList())+len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList())) > auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constants.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get() {
		return nil, errorConstants.InvalidRequest.Wrapf("total property count %d exceeds maximum %d", len(immutables.GetImmutablePropertyList().GetList())+len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()), auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constants.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get())
	}

	if utilities.IsDuplicate(append(immutables.GetImmutablePropertyList().GetList(), auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()...)) {
		return nil, errorConstants.InvalidRequest.Wrapf("duplicate properties")
	}

	classificationID := baseIDs.NewClassificationID(immutables, auxiliaryRequest.Mutables)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(classificationID))
	if classifications.Get(key.NewKey(classificationID)) != nil {
		return newAuxiliaryResponse(classificationID), errorConstants.EntityAlreadyExists.Wrapf("classification with ID %s already exists", classificationID.AsString())
	}

	if err := auxiliaryKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.AccAddress, module.Name, sdkTypes.NewCoins(sdkTypes.NewCoin(auxiliaryKeeper.stakingKeeper.BondDenom(sdkTypes.UnwrapSDKContext(context)), sdkTypes.NewInt(bondAmount.Get())))); err != nil {
		return nil, err
	}

	classifications.Add(mappable.NewMappable(base.NewClassification(immutables, auxiliaryRequest.Mutables)))

	return newAuxiliaryResponse(classificationID), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper
	auxiliaryKeeper.parameterManager = parameterManager
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bankKeeper.Keeper:
			auxiliaryKeeper.bankKeeper = value
		case stakingKeeper.Keeper:
			auxiliaryKeeper.stakingKeeper = value
		}
	}
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
