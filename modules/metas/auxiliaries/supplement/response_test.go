// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/errors"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewID("id"), baseData.NewStringData("Data"))
	metaPropertyList := base.NewMetaPropertyList([]properties.MetaProperty{metaProperty}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList, nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, MetaProperties: metaPropertyList}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList, errors.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: errors.IncorrectFormat, MetaProperties: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testAuxiliaryResponse2.GetError())

	Properties, Error := GetMetaPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, metaPropertyList, Properties)
	require.Equal(t, nil, Error)

	properties2, Error := GetMetaPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, properties2)
	require.Equal(t, errors.IncorrectFormat, Error)
}
