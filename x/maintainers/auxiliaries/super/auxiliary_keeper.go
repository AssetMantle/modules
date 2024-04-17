// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/maintainers/key"
	"github.com/AssetMantle/modules/x/maintainers/record"
	"github.com/AssetMantle/modules/x/maintainers/utilities"
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
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(auxiliaryRequest.ToIdentityID)),
		)))

	maintainers := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(maintainerID))
	if maintainers.GetMappable(key.NewKey(maintainerID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("maintainer with ID %s already exists", maintainerID)
	}

	// remove system reserved properties
	maintainedPropertyIDLIst := auxiliaryRequest.MaintainedMutables.GetMutablePropertyList().GetPropertyIDList().Remove(constantProperties.BondAmountProperty.GetID(), constantProperties.AuthenticationProperty.GetID())

	maintainer := base.NewMaintainer(auxiliaryRequest.ToIdentityID, auxiliaryRequest.MaintainedClassificationID, maintainedPropertyIDLIst, utilities.SetModulePermissions(true, true, true).Add(baseIDs.StringIDsToIDs(auxiliaryRequest.PermissionIDs)...))

	if err := maintainer.ValidateBasic(); err != nil {
		return nil, err
	}

	maintainers.Add(record.NewRecord(maintainer))

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
