/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package initialize

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/metas/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	metaID := mapper.NewMetaID(base.NewID(metaUtilities.Hash(auxiliaryRequest.Data)))
	metas := mapper.NewMetas(auxiliaryKeeper.mapper, context).Fetch(metaID)
	meta := metas.Get(metaID)
	if meta != nil {
		return newAuxiliaryResponse(nil)
	}
	metas.Add(mapper.NewMeta(auxiliaryRequest.Data))
	return newAuxiliaryResponse(nil)
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
