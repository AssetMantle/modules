package base

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Property(t *testing.T) {

	id := NewID("ID")
	fact := NewFact(NewStringData("Data"))

	testProperty := NewProperty(id, fact)
	require.Equal(t, property{ID: id, Fact: fact}, testProperty)
	require.Equal(t, id, testProperty.GetID())
	require.Equal(t, fact, testProperty.GetFact())

	readProperty, error := ReadProperty("ID2:S|SomeData")
	require.Equal(t, NewProperty(NewID("ID2"), NewFact(NewStringData("SomeData"))), readProperty)
	require.Nil(t, error)

	readProperty, error = ReadProperty("Random")
	require.Equal(t, nil, readProperty)
	require.Equal(t, errors.IncorrectFormat, error)
}
