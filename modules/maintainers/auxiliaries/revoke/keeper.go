// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	maintainerID := baseIDs.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.FromID)
	maintainers := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(maintainerID))
	fromMaintainerID := baseIDs.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.FromID)
	maintainers := auxiliaryKeeper.mapper.NewCollection(context)

	fromMaintainer := maintainers.Fetch(key.FromID(fromMaintainerID)).Get(key.NewKey(fromMaintainerID))
	if fromMaintainer == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}

	if !fromMaintainer.(mappables.Maintainer).CanRemoveMaintainer() {
		return newAuxiliaryResponse(constants.NotAuthorized)
	}

	toMaintainerID := key.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.ToID)
	toMaintainer := maintainers.Fetch(key.FromID(toMaintainerID)).Get(key.FromID(toMaintainerID))
	if toMaintainer == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}

	maintainers.Remove(toMaintainer)

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
