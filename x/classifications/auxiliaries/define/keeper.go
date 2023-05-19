// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/key"
	"github.com/AssetMantle/modules/x/classifications/mappable"
	"github.com/AssetMantle/modules/x/classifications/module"
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

	totalWeight := sdkTypes.ZeroInt()
	for _, property := range append(auxiliaryRequest.Immutables.GetImmutablePropertyList().Get(), auxiliaryRequest.Mutables.GetMutablePropertyList().Get()...) {
		totalWeight = totalWeight.Add(property.Get().GetBondWeight())
	}
	bondAmount := baseData.NewNumberData(auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constants.BondRateProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get().Mul(totalWeight))
	immutables := baseQualified.NewImmutables(auxiliaryRequest.Immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), bondAmount)))

	if sdkTypes.NewInt(int64(len(immutables.GetImmutablePropertyList().Get()) + len(auxiliaryRequest.Mutables.GetMutablePropertyList().Get()))).GT(auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constants.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get()) {
		return nil, errorConstants.InvalidRequest.Wrapf("total property count %d exceeds maximum %d", len(immutables.GetImmutablePropertyList().Get())+len(auxiliaryRequest.Mutables.GetMutablePropertyList().Get()), auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constants.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get())
	}

	if err := baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(append(immutables.GetImmutablePropertyList().Get(), auxiliaryRequest.Mutables.GetMutablePropertyList().Get()...)...)...).ValidateBasic(); err != nil {
		return nil, errorConstants.InvalidRequest.Wrapf("invalid properties: %s", err.Error())
	}

	classificationID := baseIDs.NewClassificationID(immutables, auxiliaryRequest.Mutables)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(classificationID))
	if classifications.GetMappable(key.NewKey(classificationID)) != nil {
		return newAuxiliaryResponse(classificationID), errorConstants.EntityAlreadyExists.Wrapf("classification with ID %s already exists", classificationID.AsString())
	}

	if err := auxiliaryKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.AccAddress, module.Name, sdkTypes.NewCoins(sdkTypes.NewCoin(auxiliaryKeeper.stakingKeeper.BondDenom(sdkTypes.UnwrapSDKContext(context)), bondAmount.Get()))); err != nil {
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
