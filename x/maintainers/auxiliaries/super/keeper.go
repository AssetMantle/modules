// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	"context"

	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/documents/base"
	"github.com/AssetMantle/schema/x/documents/constants"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	"github.com/AssetMantle/schema/x/helpers"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	constantProperties "github.com/AssetMantle/schema/x/properties/constants"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"

	"github.com/AssetMantle/modules/x/maintainers/internal/key"
	"github.com/AssetMantle/modules/x/maintainers/internal/mappable"
	"github.com/AssetMantle/modules/x/maintainers/internal/utilities"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	maintainerID := baseIDs.NewMaintainerID(constants.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.MaintainedClassificationID)),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.ToIdentityID)),
		)))

	maintainers := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(maintainerID))
	if maintainers.Get(key.NewKey(maintainerID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("maintainer with ID %s already exists", maintainerID)
	}

	maintainers.Add(mappable.NewMappable(base.NewMaintainer(auxiliaryRequest.ToIdentityID, auxiliaryRequest.MaintainedClassificationID, auxiliaryRequest.MaintainedMutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetPermissions(true, true, true, true, true, true))))

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
