// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/stretchr/testify/require"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	metaPropertyList := base.NewPropertyList([]properties.Property{metaProperty}...)
	property := baseProperties.NewMesaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	propertyList := base.NewPropertyList([]properties.Property{property}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList.ScrubData())
	require.Equal(t, auxiliaryResponse{PropertyList: metaPropertyList.ScrubData()}, testAuxiliaryResponse)

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList.ScrubData())
	require.Equal(t, auxiliaryResponse{PropertyList: nil}, testAuxiliaryResponse2)

	propertiesFromResponse := GetPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, propertyList, propertiesFromResponse)

	propertiesFromResponse2 := GetPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, propertiesFromResponse2)

	propertiesFromResponse3 := GetPropertiesFromResponse(nil)
	require.Equal(t, nil, propertiesFromResponse3)
}
