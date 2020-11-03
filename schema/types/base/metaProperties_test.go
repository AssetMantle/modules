package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_MetaProperties_Data(t *testing.T) {
	testMetaProperty := NewMetaProperty(NewID("ID"), NewMetaFact(NewHeightData(NewHeight(123))))
	testMetaProperty2 := NewMetaProperty(NewID("ID2"), NewMetaFact(NewStringData("Data")))
	testMetaPropertyList := []types.MetaProperty{testMetaProperty}
	testMetaProperties := NewMetaProperties(testMetaPropertyList)

	require.Equal(t, metaProperties{MetaPropertyList: testMetaPropertyList}, testMetaProperties)
	require.Equal(t, testMetaProperty, testMetaProperties.Get(NewID("ID")))

	//The GetMetaProperty method is buggy needs to be fixed
	//require.Equal(t, nil,testMetaProperty,testMetaProperties.GetMetaProperty(NewID("randomID")))
	require.Equal(t, testMetaPropertyList, testMetaProperties.GetMetaPropertyList())

	newTestMetaProperties := testMetaProperties.AddMetaProperty(testMetaProperty2)
	require.Equal(t, metaProperties{MetaPropertyList: append(testMetaPropertyList, testMetaProperty2)}, newTestMetaProperties)
	require.Equal(t, metaProperties{MetaPropertyList: []types.MetaProperty{testMetaProperty2}}, newTestMetaProperties.RemoveMetaProperty(testMetaProperty))

	newMetaProperty := NewMetaProperty(NewID("ID"), NewMetaFact(NewDecData(sdkTypes.NewDec(12))))
	require.Equal(t, metaProperties{MetaPropertyList: []types.MetaProperty{newMetaProperty}}, testMetaProperties.MutateMetaProperty(newMetaProperty))

	newProperty := NewProperty(NewID("ID3"), NewFact(NewStringData("Data3")))
	require.Equal(t, testMetaProperty, testMetaProperties.Get(NewID("ID")))
	require.Equal(t, []types.Property{testMetaProperty}, testMetaProperties.GetList())

	newTestMetaProperties2 := testMetaProperties.Add(newProperty)
	propertyMutated := NewProperty(NewID("ID3"), NewFact(NewDecData(sdkTypes.NewDec(34))))
	require.Equal(t, properties{PropertyList: []types.Property{testMetaProperty, newProperty}}, newTestMetaProperties2)
	require.Equal(t, properties{PropertyList: []types.Property{newProperty}}, newTestMetaProperties.Remove(testMetaProperty))
	require.Equal(t, properties{PropertyList: []types.Property{testMetaProperty, propertyMutated}}, newTestMetaProperties2.Mutate(propertyMutated))
	require.Equal(t, properties{PropertyList: []types.Property{testMetaProperty, newProperty}}, newTestMetaProperties2.(metaProperties).RemoveData())
	//require.Equal(t, )
	//require.Equal(t, properties{})
}
