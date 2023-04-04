// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"context"

	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/documents/constants"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	"github.com/AssetMantle/schema/x/helpers"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	constantProperties "github.com/AssetMantle/schema/x/properties/constants"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	maintainers := auxiliaryKeeper.mapper.NewCollection(context)

	fromMaintainerID := baseIDs.NewMaintainerID(constants.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.FromID)),
		)))

	if Mappable := maintainers.Fetch(key.NewKey(fromMaintainerID)).Get(key.NewKey(fromMaintainerID)); Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("maintainer with ID %s not found", fromMaintainerID.AsString())
	} else if !mappable.GetMaintainer(Mappable).IsPermitted(module.Remove) {
		return nil, errorConstants.NotAuthorized.Wrapf("maintainer with ID %s is not authorized to remove maintainers", fromMaintainerID.AsString())
	}

	toMaintainerID := baseIDs.NewMaintainerID(constants.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.ToID)),
		)))

	if Mappable := maintainers.Fetch(key.NewKey(toMaintainerID)).Get(key.NewKey(toMaintainerID)); Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("maintainer with ID %s not found", toMaintainerID.AsString())
	} else {
		maintainers.Remove(Mappable)
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
