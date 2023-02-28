// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	"fmt"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
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
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	totalWeight := int64(0)
	for _, property := range append(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList(), auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()...) {
		totalWeight += property.Get().GetBondWeight()
	}
	immutables := baseQualified.NewImmutables(auxiliaryRequest.Immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), baseData.NewNumberData(auxiliaryKeeper.parameterManager.GetParameter(constants.BondRateProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get()*totalWeight))))

	x := auxiliaryKeeper.parameterManager.GetParameter(constants.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get()
	fmt.Println(x)
	if int64(len(immutables.GetImmutablePropertyList().GetList())+len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList())) > auxiliaryKeeper.parameterManager.GetParameter(constants.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get() {
		return nil, errorConstants.InvalidRequest.Wrapf("total property count %d exceeds maximum %d", len(immutables.GetImmutablePropertyList().GetList())+len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()), auxiliaryKeeper.parameterManager.GetParameter(constants.MaxPropertyCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get())
	}

	if utilities.IsDuplicate(append(immutables.GetImmutablePropertyList().GetList(), auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()...)) {
		return nil, errorConstants.InvalidRequest.Wrapf("duplicate properties")
	}

	classificationID := baseIDs.NewClassificationID(immutables, auxiliaryRequest.Mutables)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(classificationID))
	if classifications.Get(key.NewKey(classificationID)) != nil {
		return newAuxiliaryResponse(classificationID), errorConstants.EntityAlreadyExists.Wrapf("classification with ID %s already exists", classificationID.AsString())
	}

	classifications.Add(mappable.NewMappable(base.NewClassification(immutables, auxiliaryRequest.Mutables)))

	return newAuxiliaryResponse(classificationID), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper, parameterManager: parameterManager}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
