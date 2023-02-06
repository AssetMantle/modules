// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/properties/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

type auxiliaryKeeper struct {
	mapper        helpers.Mapper
	parameterList helpers.ParameterList
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	if len(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList())+len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()) > module.MaxPropertyCount {
		return newAuxiliaryResponse(nil, errorConstants.InvalidRequest)
	}

	if utilities.IsDuplicate(append(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList(), auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()...)) {
		return newAuxiliaryResponse(nil, errorConstants.InvalidRequest)
	}

	totalSize := 0
	for _, prop := range append(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList(), auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()...) {
		if prop.IsMeta() {
			totalSize += prop.Get().(properties.MetaProperty).GetData().GetWidth()
		}
	}

	bondedImmutables := baseQualified.NewImmutables(auxiliaryRequest.Immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(baseIDs.NewStringID("BondingAmount"),
		baseData.NewDecData(
			func() sdkTypes.Dec {
				val1, _ := sdkTypes.NewDecFromStr(strconv.Itoa(totalSize))
				result := val1.Mul(func() sdkTypes.Dec {
					for _, param := range auxiliaryKeeper.parameterList.Get() {
						if param.GetMetaProperty().GetID().AsString() == constants.BondingWeightageString {
							res, _ := sdkTypes.NewDecFromStr(param.GetMetaProperty().GetData().AsString())
							return res
						}
					}
					return sdkTypes.ZeroDec()
				}())
				return result
			}()))))

	classificationID := baseIDs.NewClassificationID(bondedImmutables, auxiliaryRequest.Mutables)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(classificationID))
	if classifications.Get(key.NewKey(classificationID)) != nil {
		return newAuxiliaryResponse(classificationID, errorConstants.EntityAlreadyExists)
	}

	classifications.Add(mappable.NewMappable(base.NewClassification(bondedImmutables, auxiliaryRequest.Mutables)))

	return newAuxiliaryResponse(classificationID, nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterList helpers.ParameterList, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper, parameterList: parameterList}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
