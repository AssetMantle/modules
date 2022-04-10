// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Scrub_Request(t *testing.T) {

	metaProperty := baseTypes.NewMetaProperty(baseIDs.NewID("id"), baseData.NewStringData("Data"))
	testAuxiliaryRequest := NewAuxiliaryRequest(metaProperty)

	require.Equal(t, auxiliaryRequest{MetaPropertyList: []types.MetaProperty{metaProperty}}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
