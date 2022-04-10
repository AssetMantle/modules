// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Maintain_Request(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	identityID := baseIDs.NewID("identityID")
	mutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, identityID, mutableProperties)

	require.Equal(t, auxiliaryRequest{ClassificationID: classificationID, IdentityID: identityID, MaintainedProperties: mutableProperties}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
