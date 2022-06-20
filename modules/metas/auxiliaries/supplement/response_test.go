// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewID("id"), baseData.NewStringData("Data"))
	metaPropertyList := base.NewMetaProperties([]properties.MetaProperty{metaProperty}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList, nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, MetaProperties: metaPropertyList}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList, constants.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: constants.IncorrectFormat, MetaProperties: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, constants.IncorrectFormat, testAuxiliaryResponse2.GetError())

	Properties, err := GetMetaPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, metaPropertyList, Properties)
	require.Equal(t, nil, err)

	properties2, err := GetMetaPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, properties2)
	require.Equal(t, constants.IncorrectFormat, err)
}
