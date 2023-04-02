// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"context"
	"github.com/AssetMantle/modules/modules/maintainers/internal/module"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	internalUtilities "github.com/AssetMantle/modules/modules/maintainers/internal/utilities"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/documents/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	propertiesUtilities "github.com/AssetMantle/modules/schema/properties/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
	memberAuxiliary  helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	if !auxiliaryKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.DeputizeAllowedProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("deputize is not allowed")
	}

	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	fromMaintainerID := baseIDs.NewMaintainerID(constants.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.FromID)),
		)))

	maintainers := auxiliaryKeeper.mapper.NewCollection(context)

	Mappable := maintainers.Fetch(key.NewKey(fromMaintainerID)).Get(key.NewKey(fromMaintainerID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("maintainer with ID %s not found", fromMaintainerID.AsString())
	}
	fromMaintainer := mappable.GetMaintainer(Mappable)

	// TODO test
	if !(fromMaintainer.IsPermitted(module.Add) || !auxiliaryRequest.CanAddMaintainer && fromMaintainer.IsPermitted(module.Mutate) || !auxiliaryRequest.CanMutateMaintainer && fromMaintainer.IsPermitted(module.Remove) || !auxiliaryRequest.CanRemoveMaintainer) {
		return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have the required permissions")
	}

	// Checking if the fromMaintainer has access to maintain the requested properties
	// ADD calculating the properties that the fromMaintainer is trying to REMOVE for existing maintainer
	removeMaintainedPropertyList := fromMaintainer.GetMutables().GetMutablePropertyList()

	// TODO test
	for _, maintainedProperty := range auxiliaryRequest.MaintainedProperties.GetList() {
		if !fromMaintainer.MaintainsProperty(maintainedProperty.GetID()) {
			return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have access to maintain property %s", maintainedProperty.GetID().AsString())
		}
		removeMaintainedPropertyList = removeMaintainedPropertyList.Remove(maintainedProperty)
	}

	if _, err := auxiliaryKeeper.memberAuxiliary.GetKeeper().Help(context, member.NewAuxiliaryRequest(auxiliaryRequest.MaintainedClassificationID, nil, baseQualified.NewMutables(auxiliaryRequest.MaintainedProperties))); err != nil {
		return nil, err
	}

	toMaintainerID := baseIDs.NewMaintainerID(constants.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.ToID)),
		)))

	if Mappable = maintainers.Fetch(key.NewKey(toMaintainerID)).Get(key.NewKey(toMaintainerID)); Mappable == nil {
		if !fromMaintainer.IsPermitted(module.Add) {
			return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have the permission to add maintainers")
		}

		maintainers.Add(mappable.NewMappable(base.NewMaintainer(auxiliaryRequest.ToID, auxiliaryRequest.MaintainedClassificationID, auxiliaryRequest.MaintainedProperties.GetPropertyIDList(), internalUtilities.SetPermissions(auxiliaryRequest.CanMintAsset, auxiliaryRequest.CanBurnAsset, auxiliaryRequest.CanRenumerateAsset, auxiliaryRequest.CanAddMaintainer, auxiliaryRequest.CanRemoveMaintainer, auxiliaryRequest.CanMutateMaintainer))))
	} else {
		if !fromMaintainer.IsPermitted(module.Mutate) {
			return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have the permission to mutate maintainers")
		}

		maintainedProperties := mappable.GetMaintainer(Mappable).GetMutables().GetMutablePropertyList().Add(propertiesUtilities.AnyPropertyListToPropertyList(auxiliaryRequest.MaintainedProperties.GetList()...)...).Remove(propertiesUtilities.AnyPropertyListToPropertyList(removeMaintainedPropertyList.GetList()...)...)
		maintainers.Mutate(mappable.NewMappable(base.NewMaintainer(auxiliaryRequest.ToID, auxiliaryRequest.MaintainedClassificationID, maintainedProperties.GetPropertyIDList(), internalUtilities.SetPermissions(auxiliaryRequest.CanMintAsset, auxiliaryRequest.CanBurnAsset, auxiliaryRequest.CanRenumerateAsset, auxiliaryRequest.CanAddMaintainer, auxiliaryRequest.CanRemoveMaintainer, auxiliaryRequest.CanMutateMaintainer))))
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper
	auxiliaryKeeper.parameterManager = parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case member.Auxiliary.GetName():
				auxiliaryKeeper.memberAuxiliary = value
			}
		}
	}

	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
