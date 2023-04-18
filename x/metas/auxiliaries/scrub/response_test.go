// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"testing"

	baseData "github.com/AssetMantle/schema/x/data/base"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/AssetMantle/schema/x/lists/base"
	"github.com/AssetMantle/schema/x/properties"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	"github.com/stretchr/testify/require"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	metaPropertyList := base.NewPropertyList([]properties.Property{metaProperty}...)
	property := baseProperties.NewMesaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	propertyList := base.NewPropertyList([]properties.Property{property}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList.ScrubData(), nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, PropertyList: metaPropertyList.ScrubData()}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList.ScrubData(), errorConstants.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Error: errorConstants.IncorrectFormat, PropertyList: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errorConstants.IncorrectFormat, testAuxiliaryResponse2.GetError())

	propertiesFromResponse, err := GetPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, propertyList, propertiesFromResponse)
	require.Equal(t, nil, err)

	propertiesFromResponse2, err := GetPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, propertiesFromResponse2)
	require.Equal(t, errorConstants.IncorrectFormat, err)

	propertiesFromResponse3, err := GetPropertiesFromResponse(nil)
	require.Equal(t, nil, propertiesFromResponse3)
	require.Equal(t, errorConstants.NotAuthorized, err)
}
