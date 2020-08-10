/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type metas struct {
	ID   types.ID         `json:"id" valid:"required~required field id missing"`
	List []mappables.Meta `json:"list" valid:"required~required field list missing"`

	mapper  helpers.Mapper
	context sdkTypes.Context
}

var _ mappers.Metas = (*metas)(nil)

func (metas metas) GetID() types.ID { return metas.ID }
func (metas metas) Get(id types.ID) mappables.Meta {
	metaID := metaIDFromInterface(id)
	for _, oldMeta := range metas.List {
		if oldMeta.GetID().Compare(metaID) == 0 {
			return oldMeta
		}
	}
	return nil
}
func (metas metas) GetList() []mappables.Meta {
	return metas.List
}

func (metas metas) Fetch(id types.ID) mappers.Metas {
	var metaList []mappables.Meta
	metasID := metaIDFromInterface(id)
	mappable := metas.mapper.Read(metas.context, metasID)
	if mappable != nil {
		metaList = append(metaList, mappable.(text))
	}
	metas.ID, metas.List = id, metaList
	return metas
}
func (metas metas) Add(meta mappables.Meta) mappers.Metas {
	metas.ID = readMetaID("")
	metas.mapper.Create(metas.context, meta)
	metas.List = append(metas.List, meta)
	return metas
}
func (metas metas) Remove(meta mappables.Meta) mappers.Metas {
	metas.mapper.Delete(metas.context, meta.GetID())
	for i, oldMeta := range metas.List {
		if oldMeta.GetID().Compare(meta.GetID()) == 0 {
			metas.List = append(metas.List[:i], metas.List[i+1:]...)
			break
		}
	}
	return metas
}
func (metas metas) Mutate(meta mappables.Meta) mappers.Metas {
	metas.mapper.Update(metas.context, meta)
	for i, oldMeta := range metas.List {
		if oldMeta.GetID().Compare(meta.GetID()) == 0 {
			metas.List[i] = meta
			break
		}
	}
	return metas
}

func NewMetas(mapper helpers.Mapper, context sdkTypes.Context) mappers.Metas {
	return metas{
		ID:      readMetaID(""),
		List:    []mappables.Meta{},
		mapper:  mapper,
		context: context,
	}
}
