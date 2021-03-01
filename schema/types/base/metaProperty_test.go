/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/stretchr/testify/require"
)

func Test_MetaProperty(t *testing.T) {

	metaFact1 := NewMetaFact(NewHeightData(NewHeight(123)))
	testMetaProperty := NewMetaProperty(NewID("ID"), metaFact1)
	require.Equal(t, metaFact1, testMetaProperty.GetMetaFact())
	require.Equal(t, NewProperty(NewID("ID"), NewFact(NewHeightData(NewHeight(123)))), testMetaProperty.RemoveData())
	require.Equal(t, NewID("ID"), testMetaProperty.GetID())
	require.Equal(t, metaFact1, testMetaProperty.GetMetaFact())
	readMetaProperty, Error := ReadMetaProperty("ID2:S|SomeData")
	require.Equal(t, NewMetaProperty(NewID("ID2"), NewMetaFact(NewStringData("SomeData"))), readMetaProperty)
	require.Nil(t, Error)

	readMetaProperty, Error = ReadMetaProperty("RandomValue")
	require.Equal(t, nil, readMetaProperty)
	require.Equal(t, errors.IncorrectFormat, Error)

	readMetaProperty, Error = ReadMetaProperty("RandomID:RandomValue")
	require.Equal(t, nil, readMetaProperty)
	require.Equal(t, errors.IncorrectFormat, Error)
}
