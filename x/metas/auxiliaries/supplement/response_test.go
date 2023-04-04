// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/schema/x/data/base"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/AssetMantle/schema/x/lists/base"
	"github.com/AssetMantle/schema/x/properties"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	metaPropertyList := base.NewPropertyList([]properties.Property{metaProperty}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList, nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, PropertyList: metaPropertyList}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList, errorConstants.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: errorConstants.IncorrectFormat, PropertyList: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errorConstants.IncorrectFormat, testAuxiliaryResponse2.GetError())

	Properties, err := GetMetaPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, metaPropertyList, Properties)
	require.Equal(t, nil, err)

	properties2, err := GetMetaPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, properties2)
	require.Equal(t, errorConstants.IncorrectFormat, err)
}
