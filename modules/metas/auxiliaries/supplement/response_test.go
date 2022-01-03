/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package supplement

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Super_Response(t *testing.T) {
	metaProperty := base.NewMetaProperty(base.NewID("id"), base.NewMetaFact(base.NewStringData("Data")))
	metaPropertyList := base.NewMetaProperties([]types.MetaProperty{metaProperty}...)

	testAuxiliaryResponse := newAuxiliaryResponse(metaPropertyList, nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, MetaProperties: metaPropertyList}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(metaPropertyList, errors.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: errors.IncorrectFormat, MetaProperties: nil}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testAuxiliaryResponse2.GetError())

	properties, err := GetMetaPropertiesFromResponse(testAuxiliaryResponse)
	require.Equal(t, metaPropertyList, properties)
	require.Equal(t, nil, err)

	properties2, err := GetMetaPropertiesFromResponse(testAuxiliaryResponse2)
	require.Equal(t, nil, properties2)
	require.Equal(t, errors.IncorrectFormat, err)
}
