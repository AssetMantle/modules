/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestCollection(t *testing.T) {

	context, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)

	// Initialize
	Collection := collection{}.Initialize(context, Mapper).(collection)
	require.Equal(t, reflect.TypeOf(Mapper), reflect.TypeOf(Collection.mapper))

	// Add
	collection1 := Collection.Add(base.NewMappable("test1", "value1"))
	require.Equal(t, []helpers.Mappable{base.NewMappable("test1", "value1")}, collection1.GetList())
	require.Nil(t, collection1.GetKey())

	_ = Collection.Add(base.NewMappable("test2", "value2"))
	_ = Collection.Add(base.NewMappable("test3", "value3"))

	// Mutate
	Collection.Mutate(base.NewMappable("test2", "value3"))
	require.Equal(t, base.NewMappable("test2", "value3"), Collection.Fetch(base.NewKey("test2")).Get(base.NewKey("test2")))
	require.NotEqual(t, base.NewMappable("test2", "value2"), Collection.Fetch(base.NewKey("test2")).Get(base.NewKey("test2")))

	// Fetch
	require.Equal(t, []helpers.Mappable{base.NewMappable("test1", "value1")}, Collection.Fetch(base.NewKey("test1")).GetList())

	// Get
	Collection.Get(base.NewKey("test1"))
	require.Equal(t, nil, Collection.Get(base.NewKey("test1")))
	require.Equal(t, base.NewMappable("test1", "value1"), Collection.Fetch(base.NewKey("test1")).Get(base.NewKey("test1")))

	// GetKey
	require.Equal(t, nil, Collection.GetKey())
	require.Equal(t, base.NewKey("test1"), Collection.Fetch(base.NewKey("test1")).GetKey())
	require.Equal(t, base.NewKey("test4"), Collection.Fetch(base.NewKey("test4")).GetKey())

	// GetList
	Collection.GetList()
	require.Equal(t, []helpers.Mappable{base.NewMappable("test1", "value1")}, Collection.Fetch(base.NewKey("test1")).GetList())

	// Remove
	Collection.Remove(base.NewMappable("test1", "value0"))
	require.Equal(t, []helpers.Mappable(nil), Collection.Fetch(base.NewKey("test1")).GetList())

}
