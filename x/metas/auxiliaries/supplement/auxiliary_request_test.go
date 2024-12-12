// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/stretchr/testify/require"
)

func Test_Supplement_Request(t *testing.T) {

	property := baseProperties.NewMesaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	testAuxiliaryRequest := NewAuxiliaryRequest(property)

	require.Equal(t, auxiliaryRequest{PropertyList: []properties.Property{property}}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())

}
