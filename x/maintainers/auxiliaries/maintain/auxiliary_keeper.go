// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	"context"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/maintainers/key"
	"github.com/AssetMantle/modules/x/maintainers/mappable"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	maintainerID := baseIDs.NewMaintainerID(base.PrototypeMaintainer().GetClassificationID(),
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.IdentityID)),
		)))

	maintainers := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(maintainerID))

	Mappable := maintainers.GetMappable(key.NewKey(maintainerID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("maintainer with ID %s not found", maintainerID.AsString())
	}
	maintainer := mappable.GetMaintainer(Mappable)

	for _, maintainedProperty := range auxiliaryRequest.MaintainedMutables.GetMutablePropertyList().Get() {
		if !maintainer.MaintainsProperty(maintainedProperty.GetID()) {
			return nil, errorConstants.NotAuthorized.Wrapf("maintainer with ID %s does not maintain property with ID %s", maintainerID.AsString(), maintainedProperty.GetID().AsString())
		}
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
