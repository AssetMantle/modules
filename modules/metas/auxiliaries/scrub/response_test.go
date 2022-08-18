// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := base2.NewMetaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	metaPropertyList := base.NewMetaProperties([]properties.MetaProperty{metaProperty}...)
	property := base2.NewMesaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	propertyList := base.NewPropertyList([]properties.Property{property}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList.ToPropertyList(), nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, Properties: metaPropertyList.ToPropertyList()}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList.ToPropertyList(), constants.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: constants.IncorrectFormat, Properties: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, constants.IncorrectFormat, testAuxiliaryResponse2.GetError())

	propertiesFromResponse, err := GetPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, propertyList, propertiesFromResponse)
	require.Equal(t, nil, err)

	propertiesFromResponse2, err := GetPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, propertiesFromResponse2)
	require.Equal(t, constants.IncorrectFormat, err)

	propertiesFromResponse3, err := GetPropertiesFromResponse(nil)
	require.Equal(t, nil, propertiesFromResponse3)
	require.Equal(t, constants.NotAuthorized, err)
}
