// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Define_Request(t *testing.T) {

	immutableProperties := base.NewPropertyList(baseTypes.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("Data2")))
	mutableProperties := base.NewPropertyList(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")))

	testAuxiliaryRequest := NewAuxiliaryRequest(immutableProperties, mutableProperties)

	require.Equal(t, auxiliaryRequest{ImmutableProperties: immutableProperties, MutableProperties: mutableProperties}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
