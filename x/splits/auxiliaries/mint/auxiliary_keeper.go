// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"context"

	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/types/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/record"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	splitID := baseIDs.NewSplitID(auxiliaryRequest.OwnableID, auxiliaryRequest.OwnerID)
	splits := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(splitID))

	Mappable := splits.GetMappable(key.NewKey(splitID))
	if Mappable == nil {
		splits.Add(record.NewRecord(baseIDs.NewSplitID(auxiliaryRequest.OwnableID, auxiliaryRequest.OwnerID), base.NewSplit(auxiliaryRequest.Value)))
	} else {
		splits.Mutate(record.NewRecord(splitID, mappable.GetSplit(Mappable).Add(auxiliaryRequest.Value)))
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
