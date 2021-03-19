/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	conformAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	maintainers := auxiliaryKeeper.mapper.NewCollection(context)

	fromMaintainerID := key.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.FromID)

	var fromMaintainer mappables.Maintainer

	if maintainer := maintainers.Fetch(key.FromID(fromMaintainerID)).Get(key.FromID(fromMaintainerID)); maintainer != nil {
		fromMaintainer = maintainer.(mappables.Maintainer)
	} else {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}

	if !(fromMaintainer.CanAddMaintainer() || !auxiliaryRequest.AddMaintainer && fromMaintainer.CanMutateMaintainer() || !auxiliaryRequest.MutateMaintainer && fromMaintainer.CanRemoveMaintainer() || !auxiliaryRequest.RemoveMaintainer) {
		return newAuxiliaryResponse(errors.NotAuthorized)
	}

	if auxiliaryResponse := auxiliaryKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(auxiliaryRequest.ClassificationID, nil, auxiliaryRequest.MaintainedProperties)); !auxiliaryResponse.IsSuccessful() {
		return newAuxiliaryResponse(auxiliaryResponse.GetError())
	}

	removeMaintainedProperties := fromMaintainer.(mappables.Maintainer).GetMaintainedProperties()

	for _, maintainedProperty := range auxiliaryRequest.MaintainedProperties.GetList() {
		if !fromMaintainer.(mappables.Maintainer).MaintainsProperty(maintainedProperty.GetID()) {
			return newAuxiliaryResponse(errors.NotAuthorized)
		}

		removeMaintainedProperties.Remove(maintainedProperty)
	}

	toMaintainerID := key.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.ToID)

	toMaintainer := maintainers.Fetch(key.FromID(toMaintainerID)).Get(key.FromID(toMaintainerID))
	if toMaintainer == nil {
		if !fromMaintainer.(mappables.Maintainer).CanAddMaintainer() {
			return newAuxiliaryResponse(errors.NotAuthorized)
		}

		maintainers.Add(mappable.NewMaintainer(toMaintainerID, auxiliaryRequest.MaintainedProperties, auxiliaryRequest.AddMaintainer, auxiliaryRequest.RemoveMaintainer, auxiliaryRequest.MutateMaintainer))
	} else {
		if !fromMaintainer.(mappables.Maintainer).CanMutateMaintainer() {
			return newAuxiliaryResponse(errors.NotAuthorized)
		}

		maintainedProperties := toMaintainer.(mappables.Maintainer).GetMaintainedProperties().Add(auxiliaryRequest.MaintainedProperties.GetList()...).Remove(removeMaintainedProperties.GetList()...)
		maintainers.Mutate(mappable.NewMaintainer(toMaintainerID, maintainedProperties, auxiliaryRequest.AddMaintainer, auxiliaryRequest.RemoveMaintainer, auxiliaryRequest.MutateMaintainer))
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
			panic(errors.UninitializedUsage)
		}
	}

	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
