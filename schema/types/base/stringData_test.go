package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_StringData(t *testing.T) {

	value := "data"
	testStringData := NewStringData(value)
	testStringData2 := NewStringData("")

	require.Equal(t, value, testStringData.String())
	require.Equal(t, meta.Hash(value), testStringData.GenerateHash())
	require.Equal(t, "", testStringData2.GenerateHash())

	dataAsString, error := testStringData.AsString()
	require.Equal(t, value, dataAsString)
	require.Equal(t, nil, error)

	dataAsID, error := testStringData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.EntityNotFound, error)

	dataAsHeight, error := testStringData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.EntityNotFound, error)

	dataAsDec, error := testStringData.AsDec()
	require.Equal(t, sdkTypes.Dec{}, dataAsDec)
	require.Equal(t, errors.EntityNotFound, error)

	require.Equal(t, value, testStringData.Get())

	data, error := ReadStringData("testString")
	require.Equal(t, stringData{Value: "testString"}, data)
	require.Nil(t, error)

	require.Equal(t, false, testStringData.Equal(testStringData2))
	require.Equal(t, true, testStringData.Equal(testStringData))
}
