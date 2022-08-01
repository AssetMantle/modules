// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	conformAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	maintainers := auxiliaryKeeper.mapper.NewCollection(context)

	fromMaintainerID := baseIDs.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.FromID)

	var fromMaintainer mappables.Maintainer

	if maintainer := maintainers.Fetch(key.NewKey(fromMaintainerID)).Get(key.NewKey(fromMaintainerID)); maintainer != nil {
		fromMaintainer = maintainer.(mappables.Maintainer)
	} else {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}

	if !(fromMaintainer.CanAddMaintainer() || !auxiliaryRequest.AddMaintainer && fromMaintainer.CanMutateMaintainer() || !auxiliaryRequest.MutateMaintainer && fromMaintainer.CanRemoveMaintainer() || !auxiliaryRequest.RemoveMaintainer) {
		return newAuxiliaryResponse(constants.NotAuthorized)
	}

	if auxiliaryResponse := auxiliaryKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(auxiliaryRequest.ClassificationID, nil, auxiliaryRequest.MaintainedProperties)); !auxiliaryResponse.IsSuccessful() {
		return newAuxiliaryResponse(auxiliaryResponse.GetError())
	}

	removeMaintainedPropertyList := fromMaintainer.GetMutables().GetMutablePropertyList()

	// TODO test
	for _, maintainedProperty := range auxiliaryRequest.MaintainedProperties.GetList() {
		if !fromMaintainer.MaintainsProperty(maintainedProperty.GetID()) {

			return newAuxiliaryResponse(constants.NotAuthorized)
		}
		removeMaintainedPropertyList = removeMaintainedPropertyList.Remove(maintainedProperty)
	}

	toMaintainerID := baseIDs.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.ToID)

	toMaintainer := maintainers.Fetch(key.NewKey(toMaintainerID)).Get(key.NewKey(toMaintainerID))
	if toMaintainer == nil {
		if !fromMaintainer.CanAddMaintainer() {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}

		maintainers.Add(mappable.NewMaintainer(toMaintainerID, baseLists.NewPropertyList(), auxiliaryRequest.MaintainedProperties))
	} else {
		if !fromMaintainer.CanMutateMaintainer() {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}
		maintainedProperties := toMaintainer.(mappables.Maintainer).GetMutables().GetMutablePropertyList().Add(auxiliaryRequest.MaintainedProperties.GetList()...).Remove(removeMaintainedPropertyList.GetList()...)
		maintainers.Mutate(mappable.NewMaintainer(toMaintainerID, baseLists.NewPropertyList(), maintainedProperties))
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				auxiliaryKeeper.conformAuxiliary = value
			default:
				break
			}
		default:
			panic(constants.UninitializedUsage)
		}
	}

	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
