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
	require.Equal(t, NewID(meta.Hash("ID")), testIDData.GenerateHashID())
	require.Equal(t, NewID(""), testIDData2.GenerateHashID())

	dataAsString, Error := testIDData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.EntityNotFound, Error)

	dataAsID, Error := testIDData.AsID()
	require.Equal(t, idValue, dataAsID)
	require.Equal(t, nil, Error)

	dataAsHeight, Error := testIDData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.EntityNotFound, Error)

	dataAsDec, Error := testIDData.AsDec()
	require.Equal(t, sdkTypes.Dec{}, dataAsDec)
	require.Equal(t, errors.EntityNotFound, Error)

	require.Equal(t, idValue, testIDData.Get())

	data, Error := ReadIDData("testString")
	require.Equal(t, idData{Value: id{IDString: "testString"}}, data)
	require.Nil(t, Error)

	require.Equal(t, false, testIDData.Equal(NewStringData("")))
	require.Equal(t, true, testIDData.Equal(testIDData))

	require.Equal(t, "", testIDData.ZeroValue().String())

}
