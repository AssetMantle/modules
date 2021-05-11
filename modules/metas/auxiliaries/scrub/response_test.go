/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package scrub

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := base.NewMetaProperty(base.NewID("id"), base.NewMetaFact(base.NewStringData("Data")))
	metaPropertyList := base.NewMetaProperties([]types.MetaProperty{metaProperty}...)
	property := base.NewProperty(base.NewID("id"), base.NewFact(base.NewStringData("Data")))
	propertyList := base.NewProperties([]types.Property{property}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList.RemoveData(), nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, Properties: metaPropertyList.RemoveData()}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList.RemoveData(), errors.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: errors.IncorrectFormat, Properties: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testAuxiliaryResponse2.GetError())

	propertiesFromResponse, Error := GetPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, propertyList, propertiesFromResponse)
	require.Equal(t, nil, Error)

	propertiesFromResponse2, Error := GetPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, propertiesFromResponse2)
	require.Equal(t, errors.IncorrectFormat, Error)

	propertiesFromResponse3, Error := GetPropertiesFromResponse(nil)
	require.Equal(t, nil, propertiesFromResponse3)
	require.Equal(t, errors.NotAuthorized, Error)
}
