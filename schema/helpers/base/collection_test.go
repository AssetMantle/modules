// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/test"
	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestCollection(t *testing.T) {
	context, storeKey, _ := test.SetupTest(t)
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

	// GetProperty
	Collection.Get(base.NewKey("test1"))
	require.Equal(t, nil, Collection.Get(base.NewKey("test1")))
	require.Equal(t, base.NewMappable("test1", "value1"), Collection.Fetch(base.NewKey("test1")).Get(base.NewKey("test1")))

	// GetKey
	require.Equal(t, nil, Collection.GetKey())
	require.Equal(t, base.NewKey("test1"), Collection.Fetch(base.NewKey("test1")).GetKey())
	require.Equal(t, base.NewKey("test4"), Collection.Fetch(base.NewKey("test4")).GetKey())

	// Get
	Collection.GetList()
	require.Equal(t, []helpers.Mappable{base.NewMappable("test1", "value1")}, Collection.Fetch(base.NewKey("test1")).GetList())

	// Remove
	Collection.Remove(base.NewMappable("test1", "value0"))
	require.Equal(t, []helpers.Mappable(nil), Collection.Fetch(base.NewKey("test1")).GetList())

}
