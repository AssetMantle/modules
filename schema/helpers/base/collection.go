/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type collection struct {
	Key  helpers.Key        `json:"key" valid:"required~required field key missing"`
	List []helpers.Mappable `json:"list" valid:"required~required field list missing"`

	mapper  helpers.Mapper
	context sdkTypes.Context
}

var _ helpers.Collection = (*collection)(nil)

func (collection collection) GetKey() helpers.Key { return collection.Key }
func (collection collection) Get(key helpers.Key) helpers.Mappable {
	for _, mappable := range collection.List {
		if mappable.GetKey().Matches(key) {
			return mappable
		}
	}

	return nil
}
func (collection collection) GetList() []helpers.Mappable {
	return collection.List
}
func (collection collection) Iterate(partialKey helpers.Key, accumulator func(helpers.Mappable) bool) {
	collection.mapper.Iterate(collection.context, partialKey, accumulator)
}
func (collection collection) Fetch(key helpers.Key) helpers.Collection {
	var mappableList []helpers.Mappable

	if key.IsPartial() {
		accumulator := func(mappable helpers.Mappable) bool {
			mappableList = append(mappableList, mappable)
			return false
		}
		collection.mapper.Iterate(collection.context, key, accumulator)
	} else {
		mappable := collection.mapper.Read(collection.context, key)
		if mappable != nil {
			mappableList = append(mappableList, mappable)
		}
	}

	collection.Key, collection.List = key, mappableList

	return collection
}
func (collection collection) Add(mappable helpers.Mappable) helpers.Collection {
	collection.Key = nil
	collection.mapper.Create(collection.context, mappable)
	collection.List = append(collection.List, mappable)

	return collection
}
func (collection collection) Remove(mappable helpers.Mappable) helpers.Collection {
	collection.mapper.Delete(collection.context, mappable.GetKey())

	for i, oldMappable := range collection.List {
		if oldMappable.GetKey().Matches(mappable.GetKey()) {
			collection.List = append(collection.List[:i], collection.List[i+1:]...)
			break
		}
	}

	return collection
}
func (collection collection) Mutate(mappable helpers.Mappable) helpers.Collection {
	collection.mapper.Update(collection.context, mappable)

	for i, oldMappable := range collection.List {
		if oldMappable.GetKey().Matches(mappable.GetKey()) {
			collection.List[i] = mappable
			break
		}
	}

	return collection
}

func (collection collection) Initialize(context sdkTypes.Context, mapper helpers.Mapper) helpers.Collection {
	collection.mapper = mapper
	collection.context = context

	return collection
}
