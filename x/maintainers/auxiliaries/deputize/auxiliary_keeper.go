// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/x/maintainers/constants"
	"github.com/AssetMantle/modules/x/maintainers/key"
	"github.com/AssetMantle/modules/x/maintainers/mappable"
	"github.com/AssetMantle/modules/x/maintainers/record"
	internalUtilities "github.com/AssetMantle/modules/x/maintainers/utilities"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
	memberAuxiliary  helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, AuxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest, ok := AuxiliaryRequest.(auxiliaryRequest)
	if !ok {
		return nil, errorConstants.InvalidRequest.Wrapf("invalid request type %T", AuxiliaryRequest)
	}

	if err := auxiliaryRequest.Validate(); err != nil {
		return nil, err
	}

	if !auxiliaryKeeper.parameterManager.Fetch(context).Get().GetParameter(constantProperties.DeputizeAllowedProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("deputize is not allowed")
	}

	fromMaintainerID := baseIDs.NewMaintainerID(base.PrototypeMaintainer().GetClassificationID(),
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.FromIdentityID)),
		)))

	maintainers := auxiliaryKeeper.mapper.NewCollection(context)

	Mappable := maintainers.Fetch(key.NewKey(fromMaintainerID)).GetMappable(key.NewKey(fromMaintainerID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("maintainer with ID %s not found", fromMaintainerID.AsString())
	}
	fromMaintainer := mappable.GetMaintainer(Mappable)

	// TODO test
	if !(fromMaintainer.IsPermitted(constants.CanAddMaintainerPermission) || !auxiliaryRequest.CanAddMaintainer && fromMaintainer.IsPermitted(constants.CanMutateMaintainerPermission) || !auxiliaryRequest.CanMutateMaintainer && fromMaintainer.IsPermitted(constants.CanRemoveMaintainerPermission) || !auxiliaryRequest.CanRemoveMaintainer) {
		return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have the required permissions")
	}

	for _, permissionID := range auxiliaryRequest.PermissionIDs {
		if !fromMaintainer.IsPermitted(permissionID) {
			return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have the permission to grant permission %s", permissionID.AsString())
		}
	}

	// Checking if the fromMaintainer has access to maintain the requested properties
	// ADD calculating the properties that the fromMaintainer is trying to REMOVE for existing maintainer
	removeMaintainedPropertyList := fromMaintainer.GetMutables().GetMutablePropertyList()

	// TODO test
	for _, maintainedProperty := range auxiliaryRequest.MaintainedProperties.Get() {
		if !fromMaintainer.MaintainsProperty(maintainedProperty.GetID()) {
			return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have access to maintain property %s", maintainedProperty.GetID().AsString())
		}
		removeMaintainedPropertyList = removeMaintainedPropertyList.Remove(maintainedProperty)
	}

	if _, err := auxiliaryKeeper.memberAuxiliary.GetKeeper().Help(context, member.NewAuxiliaryRequest(auxiliaryRequest.MaintainedClassificationID, nil, baseQualified.NewMutables(auxiliaryRequest.MaintainedProperties))); err != nil {
		return nil, err
	}

	toMaintainerID := baseIDs.NewMaintainerID(base.PrototypeMaintainer().GetClassificationID(),
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.ToIdentityID)),
		)))

	if Mappable = maintainers.Fetch(key.NewKey(toMaintainerID)).GetMappable(key.NewKey(toMaintainerID)); Mappable == nil {
		if !fromMaintainer.IsPermitted(constants.CanAddMaintainerPermission) {
			return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have the permission to add maintainers")
		}
		maintainer := base.NewMaintainer(auxiliaryRequest.ToIdentityID, auxiliaryRequest.MaintainedClassificationID, auxiliaryRequest.MaintainedProperties.GetPropertyIDList(), internalUtilities.SetModulePermissions(auxiliaryRequest.CanAddMaintainer, auxiliaryRequest.CanMutateMaintainer, auxiliaryRequest.CanRemoveMaintainer).Add(baseIDs.StringIDsToIDs(auxiliaryRequest.PermissionIDs)...))

		if err := maintainer.ValidateBasic(); err != nil {
			return nil, err
		}

		maintainers.Add(record.NewRecord(maintainer))
	} else {
		if !fromMaintainer.IsPermitted(constants.CanMutateMaintainerPermission) {
			return nil, errorConstants.NotAuthorized.Wrapf("maintainer does not have the permission to mutate maintainers")
		}

		maintainedProperties := mappable.GetMaintainer(Mappable).GetMaintainedProperties().
			Add(idListToListData(auxiliaryRequest.MaintainedProperties.GetPropertyIDList())...).
			Remove(idListToListData(removeMaintainedPropertyList.GetPropertyIDList())...)

		maintainer := base.NewMaintainer(auxiliaryRequest.ToIdentityID, auxiliaryRequest.MaintainedClassificationID, listDataToIDList(maintainedProperties), internalUtilities.SetModulePermissions(auxiliaryRequest.CanAddMaintainer, auxiliaryRequest.CanMutateMaintainer, auxiliaryRequest.CanRemoveMaintainer).Add(baseIDs.StringIDsToIDs(auxiliaryRequest.PermissionIDs)...))

		if err := maintainer.ValidateBasic(); err != nil {
			return nil, err
		}

		maintainers.Mutate(record.NewRecord(maintainer))
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

	helpers.PanicOnUninitializedKeeperFields(auxiliaryKeeper)
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}

func idListToListData(idList lists.IDList) []data.ListableData {
	listableData := make([]data.ListableData, len(idList.GetList()))

	for i, id := range idList.GetList() {
		listableData[i] = baseData.NewIDData(id)
	}

	return listableData
}

func listDataToIDList(listData data.ListData) lists.IDList {
	idList := make([]ids.ID, len(listData.Get()))

	for i, datum := range listData.Get() {
		idList[i] = datum.Get().(data.IDData).Get()
	}

	return baseLists.NewIDList(idList...)
}
