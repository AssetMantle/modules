package scrub

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := base.NewMetaProperty(base.NewID("id"), base.NewMetaFact(base.NewStringData("Data")))
	metaPropertyList := base.NewMetaProperties([]types.MetaProperty{metaProperty})

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList, nil)
	require.Equal(t, AuxiliaryResponse{Success: true, Error: nil, Properties: metaPropertyList}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList, errors.IncorrectFormat)
	require.Equal(t, AuxiliaryResponse{Success: false, Error: errors.IncorrectFormat, Properties: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testAuxiliaryResponse2.GetError())

	properties, Error := GetPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, metaPropertyList, properties)
	require.Equal(t, nil, Error)

	properties2, Error := GetPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, properties2)
	require.Equal(t, errors.IncorrectFormat, Error)

	properties3, Error := GetPropertiesFromResponse(nil)
	require.Equal(t, nil, properties3)
	require.Equal(t, errors.NotAuthorized, Error)
}
