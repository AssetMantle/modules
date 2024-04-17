// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/constants"
	"github.com/AssetMantle/modules/x/classifications/key"
	"github.com/AssetMantle/modules/x/classifications/record"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
	bankKeeper       bankKeeper.Keeper
	stakingKeeper    stakingKeeper.Keeper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	if !auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.DefineEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("classification defining is not enabled")
	}

	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	// calculating minimum bound amount
	totalWeight := sdkTypes.ZeroInt()
	for _, property := range append(auxiliaryRequest.Immutables.GetImmutablePropertyList().Get(), auxiliaryRequest.Mutables.GetMutablePropertyList().Get()...) {
		totalWeight = totalWeight.Add(property.Get().GetBondWeight())
	}

	// adding weight of bond amount itself when property is not given
	if boundAmountProperty := auxiliaryRequest.Mutables.GetProperty(constantProperties.BondAmountProperty.GetID()); boundAmountProperty == nil {
		totalWeight = totalWeight.Add(constantProperties.BondAmountProperty.GetBondWeight())
	}

	minBondAmount := baseData.NewNumberData(auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.BondRateProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get().Mul(totalWeight))

	bondAmount := minBondAmount

	mutables := baseQualified.NewMutables(auxiliaryRequest.Mutables.GetMutablePropertyList())

	if boundAmountProperty := auxiliaryRequest.Mutables.GetProperty(constantProperties.BondAmountProperty.GetID()); boundAmountProperty == nil {
		// adding min bond amount as bond amount if not supplied
		mutables = baseQualified.NewMutables(mutables.GetMutablePropertyList().Add(baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetKey(), minBondAmount)))
	} else if !boundAmountProperty.Get().IsMeta() {
		return nil, errorConstants.InvalidRequest.Wrapf("bound amount is not revealed")
	} else if bondAmount = boundAmountProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData); bondAmount.Compare(minBondAmount) < 0 {
		return nil, errorConstants.InvalidRequest.Wrapf("bound amount is less than min allowed %s", minBondAmount.Get().String())
	}

	if totalPropertyCount := sdkTypes.NewInt(int64(len(auxiliaryRequest.Immutables.GetImmutablePropertyList().Get()) + len(mutables.GetMutablePropertyList().Get()))); totalPropertyCount.GT(auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get()) {
		return nil, errorConstants.InvalidRequest.Wrapf("total property count %s exceeds maximum %s", totalPropertyCount.String(), auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get().String())
	}

	classificationID := baseIDs.NewClassificationID(auxiliaryRequest.Immutables, mutables)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(classificationID))
	if classifications.GetMappable(key.NewKey(classificationID)) != nil {
		return newAuxiliaryResponse(classificationID), errorConstants.EntityAlreadyExists.Wrapf("classification with ID %s already exists", classificationID.AsString())
	}

	if err := auxiliaryKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.AccAddress, constants.ModuleName, sdkTypes.NewCoins(sdkTypes.NewCoin(auxiliaryKeeper.stakingKeeper.BondDenom(sdkTypes.UnwrapSDKContext(context)), bondAmount.Get()))); err != nil {
		return nil, err
	}

	classification := base.NewClassification(auxiliaryRequest.Immutables, mutables)

	if err := classification.ValidateBasic(); err != nil {
		return nil, err
	}

	classifications.Add(record.NewRecord(classification))

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
