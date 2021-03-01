/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/stretchr/testify/require"
)

func Test_MetaProperties(t *testing.T) {
	testMetaProperty := NewMetaProperty(NewID("ID"), NewMetaFact(NewHeightData(NewHeight(123))))
	testMetaProperty2 := NewMetaProperty(NewID("ID2"), NewMetaFact(NewStringData("Data")))
	testMetaPropertyList := []types.MetaProperty{testMetaProperty}
	testMetaProperties := NewMetaProperties(testMetaPropertyList...)

	require.Equal(t, metaProperties{MetaPropertyList: testMetaPropertyList}, testMetaProperties)
	require.Equal(t, testMetaProperty, testMetaProperties.Get(NewID("ID")))
	require.Equal(t, nil, testMetaProperties.Get(NewID("randomID")))
	require.Equal(t, testMetaPropertyList, testMetaProperties.GetList())

	newTestMetaProperties := testMetaProperties.Add(testMetaProperty2)
	require.Equal(t, metaProperties{MetaPropertyList: append(testMetaPropertyList, testMetaProperty2)}, newTestMetaProperties)
	require.Equal(t, metaProperties{MetaPropertyList: []types.MetaProperty{testMetaProperty}}, newTestMetaProperties.Remove(testMetaProperty2))

	newMetaProperty := NewMetaProperty(NewID("ID"), NewMetaFact(NewDecData(sdkTypes.NewDec(12))))
	require.Equal(t, metaProperties{MetaPropertyList: []types.MetaProperty{newMetaProperty}}, testMetaProperties.Mutate(newMetaProperty))

	require.Equal(t, testMetaProperty.RemoveData(), testMetaProperties.Get(NewID("ID")))
	require.Equal(t, []types.Property{testMetaProperty.RemoveData()}, testMetaProperties.GetList())

	newProperty := NewMetaProperty(NewID("ID3"), NewMetaFact(NewStringData("Data3")))
	newTestMetaProperties2 := testMetaProperties.Add(newProperty)
	propertyMutated := NewMetaProperty(NewID("ID"), NewMetaFact(NewDecData(sdkTypes.NewDec(34))))
	require.Equal(t, properties{PropertyList: []types.Property{testMetaProperty.RemoveData(), newProperty.RemoveData()}}, newTestMetaProperties2)
	require.Equal(t, properties{PropertyList: []types.Property{testMetaProperty2.RemoveData()}}, newTestMetaProperties.Remove(testMetaProperty))
	require.Equal(t, properties{PropertyList: []types.Property{propertyMutated.RemoveData(), testMetaProperty2.RemoveData()}}, newTestMetaProperties.Mutate(propertyMutated))
	require.Equal(t, properties{PropertyList: []types.Property{testMetaProperty.RemoveData(), testMetaProperty2.RemoveData()}}, newTestMetaProperties.RemoveData())

}
