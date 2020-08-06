/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) error {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	splitID := mapper.NewSplitID(auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID)
	splits := mapper.NewSplits(auxiliaryKeeper.mapper, context).Fetch(splitID)
	split := splits.Get(splitID)
	if split == nil {
		splits.Add(mapper.NewSplit(splitID, auxiliaryRequest.Split))
	} else {
		splits.Mutate(split.Receive(auxiliaryRequest.Split).(mappables.Split))
	}
	return nil
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
