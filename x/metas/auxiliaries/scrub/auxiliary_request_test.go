// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/stretchr/testify/require"
)

func Test_Scrub_Request(t *testing.T) {

	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	metaPropertyList := base.NewPropertyList([]properties.Property{metaProperty}...)

	testAuxiliaryRequest := NewAuxiliaryRequest(metaPropertyList)

	require.Equal(t, auxiliaryRequest{PropertyList: metaPropertyList}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
