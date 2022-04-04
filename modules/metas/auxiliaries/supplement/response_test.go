// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

func Test_Super_Response(t *testing.T) {

	metaProperty := base.NewMetaProperty(base.NewID("id"), base.NewStringData("Data"))
	metaPropertyList := base.NewMetaProperties([]types.MetaProperty{metaProperty}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList, nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, MetaProperties: metaPropertyList}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList, errors.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: errors.IncorrectFormat, MetaProperties: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testAuxiliaryResponse2.GetError())

	properties, Error := GetMetaPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, metaPropertyList, properties)
	require.Equal(t, nil, Error)

	properties2, Error := GetMetaPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, properties2)
	require.Equal(t, errors.IncorrectFormat, Error)
}
