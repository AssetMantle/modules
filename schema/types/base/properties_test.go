/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Properties(t *testing.T) {

	testProperty := NewProperty(NewID("ID"), NewFact(NewStringData("Data")))
	testProperty2 := NewProperty(NewID("ID2"), NewFact(NewHeightData(NewHeight(12))))
	testProperties := NewProperties(testProperty, testProperty2)

	require.Equal(t, properties{PropertyList: []types.Property{testProperty, testProperty2}}, testProperties)
	require.Equal(t, testProperty, testProperties.Get(NewID("ID")))
	require.Equal(t, nil, testProperties.Get(NewID("RandomID")))
	require.Equal(t, []types.Property{testProperty, testProperty2}, testProperties.GetList())

	newProperty := NewProperty(NewID("ID3"), NewFact(NewDecData(sdkTypes.NewDec(12))))
	require.Equal(t, properties{PropertyList: []types.Property{testProperty, testProperty2, newProperty}}, testProperties.Add(newProperty))
	require.Equal(t, properties{PropertyList: []types.Property{testProperty2}}, testProperties.Remove(testProperty))

	mutatedProperty := NewProperty(NewID("ID2"), NewFact(NewIDData(NewID("IDString"))))
	require.Equal(t, properties{PropertyList: []types.Property{mutatedProperty, testProperty2}}, testProperties.Mutate(mutatedProperty))
	readProperties, Error := ReadProperties("ID:S|Data,ID2:H|12")
	require.Equal(t, properties{PropertyList: []types.Property{testProperty, testProperty2}}, readProperties)
	require.Nil(t, Error)
}
