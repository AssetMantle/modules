package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_IDData(t *testing.T) {

	idValue := NewID("ID")
	testIDData := NewIDData(idValue)
	testIDData2 := NewIDData(NewID(""))

	require.Equal(t, "ID", testIDData.String())
	require.Equal(t, meta.Hash("ID"), testIDData.GenerateHash())
	require.Equal(t, "", testIDData2.GenerateHash())

	dataAsString, error := testIDData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.EntityNotFound, error)

	dataAsID, error := testIDData.AsID()
	require.Equal(t, idValue, dataAsID)
	require.Equal(t, nil, error)

	dataAsHeight, error := testIDData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.EntityNotFound, error)

	dataAsDec, error := testIDData.AsDec()
	require.Equal(t, sdkTypes.Dec{}, dataAsDec)
	require.Equal(t, errors.EntityNotFound, error)

	require.Equal(t, idValue, testIDData.Get())

	data, error := ReadIDData("testString")
	require.Equal(t, idData{Value: id{IDString: "testString"}}, data)
	require.Nil(t, error)

	//The Equal method is written incorrectly for maybe all data . Its calling itself recursively

	require.Equal(t, false, testIDData.Equal(NewStringData("")))
	require.Equal(t, true, testIDData.Equal(testIDData))

}
