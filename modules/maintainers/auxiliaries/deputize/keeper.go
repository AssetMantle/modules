// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	"github.com/AssetMantle/modules/modules/maintainers/internal/utilities"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type auxiliaryKeeper struct {
	mapper          helpers.Mapper
	memberAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	fromMaintainerID := baseIDs.NewMaintainerID(constansts.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.FromID)),
		)))

	maintainers := auxiliaryKeeper.mapper.NewCollection(context)

	Mappable := maintainers.Fetch(key.NewKey(fromMaintainerID)).Get(key.NewKey(fromMaintainerID))
	if Mappable == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}
	fromMaintainer := Mappable.(documents.Maintainer)

	// TODO test
	if !(fromMaintainer.CanAddMaintainer() || !auxiliaryRequest.CanAddMaintainer && fromMaintainer.CanMutateMaintainer() || !auxiliaryRequest.CanMutateMaintainer && fromMaintainer.CanRemoveMaintainer() || !auxiliaryRequest.CanRemoveMaintainer) {
		return newAuxiliaryResponse(constants.NotAuthorized)
	}

	// Checking if the fromMaintainer has access to maintain the requested properties
	// ADD calculating the properties that the fromMaintainer is trying to REMOVE for existing maintainer
	removeMaintainedPropertyList := fromMaintainer.GetMutables().GetMutablePropertyList()

	// TODO test
	for _, maintainedProperty := range auxiliaryRequest.MaintainedProperties.GetList() {
		if !fromMaintainer.MaintainsProperty(maintainedProperty.GetID()) {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}
		removeMaintainedPropertyList = removeMaintainedPropertyList.Remove(maintainedProperty)
	}

	if auxiliaryResponse := auxiliaryKeeper.memberAuxiliary.GetKeeper().Help(context, member.NewAuxiliaryRequest(auxiliaryRequest.MaintainedClassificationID, nil, baseQualified.NewMutables(auxiliaryRequest.MaintainedProperties))); !auxiliaryResponse.IsSuccessful() {
		return newAuxiliaryResponse(auxiliaryResponse.GetError())
	}

	toMaintainerID := baseIDs.NewMaintainerID(constansts.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.ToID)),
		)))

	if Mappable = maintainers.Fetch(key.NewKey(toMaintainerID)).Get(key.NewKey(toMaintainerID)); Mappable == nil {
		if !fromMaintainer.CanAddMaintainer() {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}

		maintainers.Add(mappable.NewMappable(base.NewMaintainer(auxiliaryRequest.ToID, auxiliaryRequest.MaintainedClassificationID, auxiliaryRequest.MaintainedProperties.GetPropertyIDList(), utilities.SetPermissions(auxiliaryRequest.CanMintAsset, auxiliaryRequest.CanBurnAsset, auxiliaryRequest.CanRenumerateAsset, auxiliaryRequest.CanAddMaintainer, auxiliaryRequest.CanRemoveMaintainer, auxiliaryRequest.CanMutateMaintainer))))
	} else {
		if !fromMaintainer.CanMutateMaintainer() {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}

		maintainedProperties := Mappable.(documents.Maintainer).GetMutables().GetMutablePropertyList().Add(auxiliaryRequest.MaintainedProperties.GetList()...).Remove(removeMaintainedPropertyList.GetList()...)
		maintainers.Mutate(mappable.NewMappable(base.NewMaintainer(auxiliaryRequest.ToID, auxiliaryRequest.MaintainedClassificationID, maintainedProperties.GetPropertyIDList(), utilities.SetPermissions(auxiliaryRequest.CanMintAsset, auxiliaryRequest.CanBurnAsset, auxiliaryRequest.CanRenumerateAsset, auxiliaryRequest.CanAddMaintainer, auxiliaryRequest.CanRemoveMaintainer, auxiliaryRequest.CanMutateMaintainer))))
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
