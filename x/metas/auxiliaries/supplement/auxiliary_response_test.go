// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

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

	testAuxiliaryResponse := NewAuxiliaryResponse(metaPropertyList)
	require.Equal(t, auxiliaryResponse{PropertyList: metaPropertyList}, testAuxiliaryResponse)

	testAuxiliaryResponse2 := NewAuxiliaryResponse(metaPropertyList)
	require.Equal(t, auxiliaryResponse{PropertyList: nil}, testAuxiliaryResponse2)

	Properties := GetMetaPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, metaPropertyList, Properties)

	properties2 := GetMetaPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, properties2)
}
