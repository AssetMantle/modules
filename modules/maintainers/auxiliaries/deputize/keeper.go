// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	"github.com/AssetMantle/modules/modules/maintainers/internal/utilities"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type auxiliaryKeeper struct {
	mapper          helpers.Mapper
	memberAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	maintainers := auxiliaryKeeper.mapper.NewCollection(context)

	fromMaintainerID := baseIDs.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.FromID)

	Mappable := maintainers.Fetch(key.NewKey(fromMaintainerID)).Get(key.NewKey(fromMaintainerID))
	if Mappable == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}
	fromMaintainer := Mappable.(documents.Maintainer)

	if !(fromMaintainer.CanAddMaintainer() || !auxiliaryRequest.CanAddMaintainer && fromMaintainer.CanMutateMaintainer() || !auxiliaryRequest.CanMutateMaintainer && fromMaintainer.CanRemoveMaintainer() || !auxiliaryRequest.CanRemoveMaintainer) {
		return newAuxiliaryResponse(constants.NotAuthorized)
	}

	if auxiliaryResponse := auxiliaryKeeper.memberAuxiliary.GetKeeper().Help(context, member.NewAuxiliaryRequest(auxiliaryRequest.ClassificationID, nil, baseQualified.NewMutables(auxiliaryRequest.MaintainedProperties))); !auxiliaryResponse.IsSuccessful() {
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

		maintainers.Add(mappable.NewMappable(base.NewMaintainer(auxiliaryRequest.ToID, auxiliaryRequest.ClassificationID, auxiliaryRequest.MaintainedProperties, utilities.SetPermissions(auxiliaryRequest.CanMintAsset, auxiliaryRequest.CanBurnAsset, auxiliaryRequest.CanRenumerateAsset, auxiliaryRequest.CanAddMaintainer, auxiliaryRequest.CanRemoveMaintainer, auxiliaryRequest.CanMutateMaintainer))))
	} else {
		if !fromMaintainer.CanMutateMaintainer() {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}
		maintainedProperties := toMaintainer.(documents.Maintainer).GetMutables().GetMutablePropertyList().Add(auxiliaryRequest.MaintainedProperties.GetList()...).Remove(removeMaintainedPropertyList.GetList()...)
		maintainers.Mutate(mappable.NewMappable(base.NewMaintainer(auxiliaryRequest.ToID, auxiliaryRequest.ClassificationID, maintainedProperties, utilities.SetPermissions(auxiliaryRequest.CanMintAsset, auxiliaryRequest.CanBurnAsset, auxiliaryRequest.CanRenumerateAsset, auxiliaryRequest.CanAddMaintainer, auxiliaryRequest.CanRemoveMaintainer, auxiliaryRequest.CanMutateMaintainer))))
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case member.Auxiliary.GetName():
				auxiliaryKeeper.memberAuxiliary = value
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
