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

func Test_Property(t *testing.T) {

	id := NewID("ID")
	fact := NewFact(NewStringData("Data"))

	testProperty := NewProperty(id, fact)
	require.Equal(t, property{ID: id, Fact: fact}, testProperty)
	require.Equal(t, id, testProperty.GetID())
	require.Equal(t, fact, testProperty.GetFact())

	readProperty, err := ReadProperty("ID2:S|SomeData")
	require.Equal(t, NewProperty(NewID("ID2"), NewFact(NewStringData("SomeData"))), readProperty)
	require.Nil(t, err)

	readProperty, err = ReadProperty("Random")
	require.Equal(t, nil, readProperty)
	require.Equal(t, errors.IncorrectFormat, err)
}
