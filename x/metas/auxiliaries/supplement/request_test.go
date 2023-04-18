// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/stretchr/testify/require"
)

func Test_Supplement_Request(t *testing.T) {

	property := baseProperties.NewMesaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	testAuxiliaryRequest := NewAuxiliaryRequest(property)

	require.Equal(t, auxiliaryRequest{PropertyList: []properties.Property{property}}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
